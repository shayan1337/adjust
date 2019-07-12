package app

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type Provider interface {
	Get(url string) (*http.Response, error)
}

type Hasher interface {
	Hash(response []byte) string
}

type Logger interface {
	Error(err error)
}

type Writer interface {
	Write(url string, value string)
}

type App struct {
	hasher   Hasher
	provider Provider
	logger   Logger
	writer   Writer
}

func NewApp(hasher Hasher, provider Provider, logger Logger, writer Writer) *App {
	return &App{
		hasher:   hasher,
		provider: provider,
		logger:   logger,
		writer:   writer,
	}
}

func (this *App) fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := this.provider.Get(url)
	if err != nil {
		this.logger.Error(err)
		return
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		this.logger.Error(err)
		return
	}
	this.writer.Write(url, this.hasher.Hash(bodyBytes))
}

func (this *App) Fetch(urls []string) {
	var wg sync.WaitGroup
	for _, rawUrl := range urls {
		parsedUrl, err := parseUrl(rawUrl)
		if err != nil {
			this.logger.Error(err)
			continue
		}
		wg.Add(1)
		go this.fetch(parsedUrl, &wg)
	}
	wg.Wait()
}

func parseUrl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl, err
	}
	u.Scheme = "http"
	return u.String(), nil
}
