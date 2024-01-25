package server

import (
	"errors"
	"github.com/Master0fMagic/rest-health-check-service/health"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

var (
	ErrInvalidURLMethod = errors.New("invalid url method")
)

func getUrlsHealthRequestHandlerFunc(h health.URLHealthCheck) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var urls []string

		if err := ctx.BindJSON(&urls); err != nil {
			log.WithError(err).Error("error parsing request")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		for _, rawURL := range urls {
			if err := validateURL(rawURL); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url format"})
				return
			}
		}

		resp, err := h.PingUrls(ctx, urls)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func validateURL(rawURL string) error {
	logger := log.WithField("url", rawURL)
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		logger.WithError(err).Error("error parsing url")
		return err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		logger.WithError(err).Error("invalid url method")
		return ErrInvalidURLMethod
	}
	return nil
}
