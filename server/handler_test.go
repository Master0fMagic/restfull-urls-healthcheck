package server

import (
	"fmt"
	"github.com/Master0fMagic/rest-health-check-service/health"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	hMock := health.NewMockURLHealthCheck(ctrl)

	t.Run("HandlerPositive", func(t *testing.T) {
		req := toJSON(health.TestUrls)
		ctx, rr := prepareGinContext(req)

		hMock.EXPECT().PingUrls(ctx, health.TestUrls).Return(nil, nil)

		getUrlsHealthRequestHandlerFunc(hMock)(ctx)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("HandlerNegative_InvalidUrls", func(t *testing.T) {
		for _, url := range InvalidUrls {
			url := url
			t.Run(fmt.Sprintf("HandlerNegative_InvalidUrl:<%s>", url), func(t *testing.T) {
				t.Parallel()
				req := toJSON([]string{url})
				ctx, rr := prepareGinContext(req)

				getUrlsHealthRequestHandlerFunc(hMock)(ctx)

				assert.Equal(t, http.StatusBadRequest, rr.Code)
			})
		}
	})

	t.Run("HandlerNegative_InvalidRequestBody", func(t *testing.T) {
		for _, url := range InvalidRequestBodies {
			url := url
			t.Run(fmt.Sprintf("HandlerNegative_InvalidRequestBody:<%s>", url), func(t *testing.T) {
				t.Parallel()
				req := toJSON(url)
				ctx, rr := prepareGinContext(req)

				getUrlsHealthRequestHandlerFunc(hMock)(ctx)

				assert.Equal(t, http.StatusBadRequest, rr.Code)
			})
		}
	})

	t.Run("HandlerNegative_EmptyUrls", func(t *testing.T) {
		t.Parallel()
		req := toJSON([]string{})
		ctx, rr := prepareGinContext(req)

		getUrlsHealthRequestHandlerFunc(hMock)(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

	})
}
