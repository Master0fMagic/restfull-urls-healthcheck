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
	ErrInvalidUrlMethod = errors.New("invalid url method")
)

func getUrlsHealthRequestHandlerFunc(h health.UrlHealthCheck) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var urls []string

		if err := ctx.BindJSON(&urls); err != nil {
			log.WithError(err).Error("error parsing request")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		for _, rawUrl := range urls {
			if err := validateUrl(rawUrl); err != nil {
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

func validateUrl(rawUrl string) error {
	logger := log.WithField("url", rawUrl)
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		logger.WithError(err).Error("error parsing url")
		return err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		logger.WithError(err).Error("invalid url method")
		return ErrInvalidUrlMethod
	}
	return nil
}
