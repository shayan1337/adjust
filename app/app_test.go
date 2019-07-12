package app

import (
	"adjust/hasher"
	"adjust/logger"
	"adjust/provider"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var variable = 0

func Test_ShouldAppendProtocolToUrlIfNotPresent(t *testing.T) {
	url := "adjust.com"
	expected := "http://adjust.com"
	result, err := parseUrl(url)

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_ShouldFetchUrlResponseAndPrintResponseHashToTheStandardOutput(t *testing.T) {
	url := "http://adjust.com"
	writer := TestStdWriter{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	app := NewApp(hasher.NewHash(), provider.NewHttpProvider(), logger.NewLogger(), writer)

	app.fetch(url, &wg)

	if variable != 1 {
		t.Fail()
	}
}

type TestStdWriter struct {}

func (this TestStdWriter) Write(url string, value string) {
	variable += 1
}