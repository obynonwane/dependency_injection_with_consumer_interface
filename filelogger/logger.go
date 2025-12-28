package filelogger

import (
	"log"

	interface_log "github.com/obynonwane/di_with_packages/interface"
)

type FileLogger struct {
}

func NewFileLogger() interface_log.Logger {
	return &FileLogger{}
}

func (l *FileLogger) Log(message string) {
	log.Println(message)
}
