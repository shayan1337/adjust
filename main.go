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
	app := NewApp(hasher, provider, log.Logger{})
	commandLineArgs := os.Args

	app.Fetch(commandLineArgs[1:])
}