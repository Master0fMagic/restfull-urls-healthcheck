package server

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	InvalidUrls = []string{
		"http//test.com", "http", "/", "test.com",
	}
	InvalidRequestBodies = []string{
		"http://test.com",           // just string
		"{'url':'http://test.com'}", // json object
	}
)

func toJSON(urls any) []byte {
	res, err := json.Marshal(urls)
	if err != nil {
		log.WithError(err).Fatal("error marshaling to json")
	}
	return res
}

func prepareGinContext(request []byte) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	rr := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rr)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
		Method: http.MethodPost,
	}
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Body = io.NopCloser(bytes.NewReader(request))

	return ctx, rr
}
