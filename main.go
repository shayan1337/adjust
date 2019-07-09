package main

import (
	. "adjust/app"
	. "adjust/hasher"
	. "adjust/provider"
	"log"
	"os"
)

func main() {

	hasher := NewHash()
	provider := NewHttpProvider()
	logger := log.Logger{}
	logger.SetOutput(os.Stdout)
	app := NewApp(hasher, provider, logger)
	commandLineArgs := os.Args

	app.Fetch(commandLineArgs[1:])
}