package utils

import (
	"errors"
	"net/http"
)

type HttpClient interface {
	Get(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct{}

func NewClient() HttpClient {
	return &httpClient{}
}

func (cl *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return cl.do(http.MethodGet, url, headers, nil)
}

func (cl *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("could not create request")
	}

	for name, values := range headers {
		if len(values) > 0 {
			request.Header.Set(name, values[0])
		}
	}

	return client.Do(request)
}
