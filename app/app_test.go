package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldAppendProtocolToUrlIfNotPresent(t *testing.T) {
	url := "adjust.com"
	expected := "http://adjust.com"
	result, err := parseUrl(url)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}
