package app

import (
	"adjust/hasher"
	"adjust/logger"
	"adjust/provider"
	"sync"
	"testing"
)

var variable = 0

func Test_ShouldAppendProtocolToUrlIfNotPresent(t *testing.T) {
	url := "adjust.com"
	expected := "http://adjust.com"
	result, err := parseUrl(url)

	if result != expected {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

func Test_ShouldFetchUrlResponseAndPrintResponseHashToTheStandardOutput(t *testing.T) {
	variable = 0
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

func Test_ShouldFetchUrlResponseAndPrintResponseHashToTheStandardOutputForAllURLs(t *testing.T) {
	variable = 0
	urls := []string{ "http://adjust.com", "facebook.com", "http://instagram.com"}
	writer := TestStdWriter{}
	app := NewApp(hasher.NewHash(), provider.NewHttpProvider(), logger.NewLogger(), writer)

	app.Fetch(urls)

	if variable != len(urls) {
		t.Fail()
	}
}

type TestStdWriter struct {}

func (this TestStdWriter) Write(url string, value string) {
	variable += 1
}