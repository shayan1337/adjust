package logger

import (
	"log"
	"os"
)

func GetLogger() log.Logger {
	logger := log.Logger{}
	logger.SetOutput(os.Stdout)
	return logger
}
