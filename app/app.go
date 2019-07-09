package app

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
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
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			this.logger.Println(err)
			continue
		}
		fmt.Println(url, getHash(bodyBytes))
	}
}

func getHash(response []byte) string {
	hasher := md5.New()
	hasher.Write(response)
	return hex.EncodeToString(hasher.Sum(nil))
}

