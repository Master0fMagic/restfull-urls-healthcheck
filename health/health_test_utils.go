package health

import (
	"bytes"
	"io"
	"net/http"
)

var (
	TestUrls         = []string{"http://test.com", "http://test.test.com"}
	MixedStatusesMap = map[string]Status{
		"http://test.com":      InactiveStatus,
		"http://test.test.com": ActiveStatus,
	}

	SuccessHTTPResponse = &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(nil)),
	}

	NonSuccessHTTPResponse = &http.Response{
		Status:     "500 Internal Server Error",
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(bytes.NewBuffer(nil)),
	}
)

func BuildHealthResponse(urls []string, status Status) map[string]Status {
	m := make(map[string]Status)
	for _, url := range urls {
		m[url] = status
	}
	return m
}
