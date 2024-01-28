package health

import (
	"context"
	"errors"
	"github.com/Master0fMagic/rest-health-check-service/config"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHealthService(t *testing.T) {
	ctx := context.TODO()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	httpMock := NewMockHTTPCaller(ctrl)
	hSrvs := &HTTPHealthCheck{
		cfg: config.HealthCheckConfig{
			Timeout:                 1 * time.Second,
			StopOnFailure:           true,
			MaxProcessingGoroutines: 1,
		},
		http: httpMock,
	}

	t.Run("HealthServicePositive_ActiveUrls", func(t *testing.T) {
		for range TestUrls {
			httpMock.EXPECT().Do(gomock.Any()).Return(SuccessHTTPResponse, nil)
		}

		healthRes, err := hSrvs.PingUrls(ctx, TestUrls)

		assert.NoError(t, err, "error while checking urls health")
		assert.Equal(t, BuildHealthResponse(TestUrls, ActiveStatus), healthRes, "actual health response not equal to expected")
	})

	t.Run("HealthServicePositive_NonActiveUrls", func(t *testing.T) {
		for range TestUrls {
			httpMock.EXPECT().Do(gomock.Any()).Return(NonSuccessHTTPResponse, nil)
		}

		healthRes, err := hSrvs.PingUrls(ctx, TestUrls)

		assert.NoError(t, err, "error while checking urls health")
		assert.Equal(t, BuildHealthResponse(TestUrls, InactiveStatus), healthRes, "actual health response not equal to expected")
	})

	t.Run("HealthServicePositive_MixedUrlsStatuses", func(t *testing.T) {
		httpMock.EXPECT().Do(gomock.Any()).Return(NonSuccessHTTPResponse, nil)
		httpMock.EXPECT().Do(gomock.Any()).Return(SuccessHTTPResponse, nil)

		healthRes, err := hSrvs.PingUrls(ctx, TestUrls)

		assert.NoError(t, err, "error while checking urls health")
		assert.Equal(t, MixedStatusesMap, healthRes, "actual health response not equal to expected")
	})

	t.Run("HealthServicePositive_ErrorCallingHttp", func(t *testing.T) {
		httpMock.EXPECT().Do(gomock.Any()).Return(NonSuccessHTTPResponse, nil)
		httpMock.EXPECT().Do(gomock.Any()).Return(nil, errors.New("test error"))

		healthRes, err := hSrvs.PingUrls(ctx, TestUrls)

		assert.NoError(t, err, "error while checking urls health")
		assert.Equal(t, BuildHealthResponse(TestUrls[0:1], InactiveStatus), healthRes, "actual health response not equal to expected")
	})
}
