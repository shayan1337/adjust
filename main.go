package main

import (
	"fmt"
	"os"
)

func main() {
	commandLineArgs := os.Args

	for _, commandLineArg := range commandLineArgs[1:] {
		fmt.Println(commandLineArg)
	}
}
