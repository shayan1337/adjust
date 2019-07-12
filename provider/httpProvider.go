package provider

import (
	"net/http"
)

type HttpProvider struct {
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
