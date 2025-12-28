package thirdpartylogger

import (
	"log"

	interface_log "github.com/obynonwane/di_with_packages/interface"
)

type ThirdPartyLogger struct{}

func NewThirdPartyLogger() interface_log.Logger {
	return &ThirdPartyLogger{}
}

func (l *ThirdPartyLogger) Log(message string) {
	log.Println(message)
}
