package consolelogger

import (
	"log"

	interface_log "github.com/obynonwane/di_with_packages/interface"
)

type ConsoleLogger struct {
}

func New(msg string) interface_log.Logger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(message string) {
	log.Println(message)
}
