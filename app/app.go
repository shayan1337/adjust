package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Provider interface {
	Get(url string) (*http.Response, error)
}

type Hasher interface {
	Hash(response []byte) string
}

type App struct {
	hasher   Hasher
	provider Provider
	logger   log.Logger
}

func NewApp(hasher Hasher, provider Provider, logger log.Logger) *App {
	return &App{
		hasher:   hasher,
		provider: provider,
		logger:   logger,
	}
}

func (this *App) Fetch(urls []string) {
	for _, url := range urls {
		response, err := this.provider.Get(url)
		if err != nil {
			this.logger.Println(err)
			continue
		}
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			this.logger.Println(err)
			continue
		}
		fmt.Println(url, this.hasher.Hash(bodyBytes))
	}
}
