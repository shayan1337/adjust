package provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldReturnResponseForAValidUrl(t *testing.T) {
	url := "http://adjust.com"
	httpProvider := NewHttpProvider()
	response, err := httpProvider.Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
}

func Test_ShouldReturnErrForAnInValidUrl(t *testing.T) {
	url := "abcd.com"
	httpProvider := NewHttpProvider()
	_, err := httpProvider.Get(url)

	assert.NotNil(t, err)
}
