package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
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

func (this *App) fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := this.provider.Get(url)
	if err != nil {
		this.logger.Println(err)
		return
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		this.logger.Println(err)
		return
	}
	fmt.Println(url, this.hasher.Hash(bodyBytes))
}

func (this *App) Fetch(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go this.fetch(url, &wg)
	}
	wg.Wait()
}
