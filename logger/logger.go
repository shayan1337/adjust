package logger

import (
	"log"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (this *Logger) Error(err error) {
	log.Println(err)
}
