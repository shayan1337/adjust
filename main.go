package main

import (
	. "adjust/app"
	"adjust/constants"
	"adjust/resolver"
	"flag"
	"runtime"
)

func main() {

	throttleRequest()
	commandLineArgs := getCommandLineArgs()

	NewApp(resolver.ResolveAppDependency()).Fetch(commandLineArgs)
}

func throttleRequest() {
	maxParallelRequests := flag.Int("parallel", constants.DEFAULT_GOMAXPROCS, "a flag to limit number of parallel requests")
	flag.Parse()
	runtime.GOMAXPROCS(*maxParallelRequests)
}

func getCommandLineArgs() []string{
	return flag.Args()
}
