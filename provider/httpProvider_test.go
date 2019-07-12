package provider

import (
	"testing"
)

func Test_ShouldReturnResponseForAValidUrl(t *testing.T) {
	url := "http://adjust.com"
	httpProvider := NewHttpProvider()
	response, err := httpProvider.Get(url)

	if err != nil {
		t.FailNow()
	}

	if response.StatusCode != 200 {
		t.Fail()
	}
}

func Test_ShouldReturnErrForAnInValidUrl(t *testing.T) {
	url := "abcd.com"
	httpProvider := NewHttpProvider()
	_, err := httpProvider.Get(url)

	if err == nil {
		t.Fail()
	}
}
