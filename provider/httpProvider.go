package provider

import (
	"log"
	"net/http"
)

type HttpProvider struct {
	logger log.Logger
}

func NewHttpProvider() *HttpProvider {
	return &HttpProvider{}
}

func (this *HttpProvider) Get(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
