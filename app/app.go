package app

import (
	"fmt"
	"log"
	"net/http"
)

type Provider interface {
	Get(url string) (*http.Response, error)
}

type App struct {
	provider Provider
	logger log.Logger
}

func NewApp(provider Provider, logger log.Logger) *App{
	return &App{
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
		fmt.Println(response)
	}
}

