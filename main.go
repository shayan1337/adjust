package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	commandLineArgs := os.Args

	for _, commandLineArg := range commandLineArgs[1:] {
		resp, err := Get(commandLineArg)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(resp)
		resp.Body.Close()
	}
}

func Get(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}