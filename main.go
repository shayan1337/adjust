package main

import (
	. "adjust/app"
	"adjust/resolver"
	"flag"
	"os"
	"runtime"
)

func main() {

	throttleRequest()
	commandLineArgs := getCommandLineArgs()

	NewApp(resolver.ResolveAppDependency()).Fetch(commandLineArgs)
}

func throttleRequest() {
	maxParallelRequests := flag.Int("parallel", 10, "a flag to limit number of parallel requests")
	flag.Parse()
	runtime.GOMAXPROCS(*maxParallelRequests)
}

func getCommandLineArgs() []string{
	return os.Args[1:]
}