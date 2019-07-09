package provider

import (
	"log"
	"net/http"
)

type httpProvider struct {
	logger log.Logger
}

func NewHttpProvider() *httpProvider {
	return &httpProvider{}
}

func (this *httpProvider) Get(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
