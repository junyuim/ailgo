package core_http

import (
	"bytes"
	"crypto/tls"
	"github.com/bytedance/sonic"
	"github.com/junyuim/ailgo/core/core_utils"
	"io"
	"net/http"
	"time"
)

type HttpRequestOption struct {
	Header map[string]string
	Query  map[string]string
}

func NewHttpClient(timeout time.Duration) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	return httpClient
}

func NewHttpRequest(method string, url string, body any, option *HttpRequestOption) *http.Request {
	var bodyReader io.Reader

	if body != nil {
		data, err := sonic.Marshal(body)

		if err != nil {
			core_utils.LogError("NewHttpRequest error:%s", err.Error())
			return nil
		}

		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, url, bodyReader)

	if err != nil {
		core_utils.LogError("NewHttpRequest error:%s", err.Error())
		return nil
	}

	if bodyReader != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if option != nil {
		for k, v := range option.Header {
			req.Header.Set(k, v)
		}
		for k, v := range option.Query {
			q := req.URL.Query()
			q.Set(k, v)
			req.URL.RawQuery = q.Encode()
		}
	}

	return req
}

func NewHttpServer(addr string, handler http.Handler) *http.Server {
	httpServer := &http.Server{
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		Addr:      addr,
		Handler:   handler,
	}

	return httpServer
}
