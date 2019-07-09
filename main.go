package main

import (
	. "adjust/app"
	. "adjust/provider"
	"log"
	"os"
)

func main() {

	provider := NewHttpProvider()
	app := NewApp(provider, log.Logger{})
	commandLineArgs := os.Args

	app.Fetch(commandLineArgs[1:])
}