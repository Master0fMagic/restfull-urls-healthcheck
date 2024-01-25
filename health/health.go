package health

import (
	"context"
	"errors"
	"github.com/Master0fMagic/rest-health-check-service/config"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sync"
)

type Status string

var (
	ActiveStatus   Status = "active"
	InactiveStatus Status = "inactive"
)

var (
	ErrNonOkResponse = errors.New("error response code not OK")
)

type URLHealthCheck interface {
	PingUrls(ctx context.Context, urls []string) (map[string]Status, error)
}

type HTTPHealthCheck struct {
	cfg  config.HealthCheckConfig
	http *http.Client
}

func New(cfg config.HealthCheckConfig) *HTTPHealthCheck {
	return &HTTPHealthCheck{
		cfg:  cfg,
		http: &http.Client{Timeout: cfg.Timeout},
	}
}

func (h *HTTPHealthCheck) PingURL(ctx context.Context, url string) error {
	logger := log.WithField("url", url)

	logger.Debug("trying to health check...")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.WithError(err).Error("error creating http request")
		return err
	}

	response, err := h.http.Do(req)
	if err != nil {
		logger.WithError(err).Error("error executing http request")
		return err
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return nil
	}

	logger.Errorf("error wrong response status. expected 200 OK, actual: %s", response.Status)
	return ErrNonOkResponse
}

func (h *HTTPHealthCheck) PingUrls(ctx context.Context, urls []string) (map[string]Status, error) {
	mtx := sync.Mutex{}
	healthMap := make(map[string]Status)
	errGroup, ctx := errgroup.WithContext(ctx)

	if h.cfg.MaxProcessingGoroutines > 0 {
		errGroup.SetLimit(h.cfg.MaxProcessingGoroutines)
	}

	for _, url := range urls {
		url := url
		errGroup.Go(func() error {
			err := h.PingURL(ctx, url)

			mtx.Lock()
			defer mtx.Unlock()

			switch {
			case errors.Is(err, ErrNonOkResponse):
				healthMap[url] = InactiveStatus
				if !h.cfg.StopOnFailure {
					err = nil
				}
			case err == nil:
				healthMap[url] = ActiveStatus
			}
			return err
		})
	}

	err := errGroup.Wait()
	if err != nil && !errors.Is(err, ErrNonOkResponse) {
		log.WithError(err).Error("error executing health checks for urls")
		return nil, err
	}
	return healthMap, nil
}
