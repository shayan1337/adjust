package main

import (
	. "adjust/app"
	. "adjust/hasher"
	. "adjust/provider"
	"flag"
	"log"
	"os"
	"runtime"
)

func main() {

	throttleRequest()
	commandLineArgs := getCommandLineArgs()

	hasher := NewHash()
	logger := getLogger()
	provider := NewHttpProvider()

	NewApp(hasher, provider, logger).Fetch(commandLineArgs)
}

func throttleRequest() {
	maxParallelRequests := flag.Int("parallel", 10, "a flag to limit number of parallel requests")
	flag.Parse()
	runtime.GOMAXPROCS(*maxParallelRequests)
}

func getCommandLineArgs() []string{
	return os.Args[1:]
}

func getLogger() log.Logger {
	logger := log.Logger{}
	logger.SetOutput(os.Stdout)
	return logger
}