package filelogger

import (
	"log"
)

type FileLogger struct {
}

func NewFileLogger() *FileLogger {
	return &FileLogger{}
}

func (l *FileLogger) Log(message string) {
	log.Println(message)
}
