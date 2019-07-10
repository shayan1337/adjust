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
	logger := getLogger()
	provider := NewHttpProvider()
	commandLineArgs := getCommandLineArgs()

	app := NewApp(hasher, provider, logger)

	app.Fetch(commandLineArgs[1:])
}

func getCommandLineArgs() []string{
	return os.Args[1:]
}

func getLogger() log.Logger {
	logger := log.Logger{}
	logger.SetOutput(os.Stdout)
	return logger
}