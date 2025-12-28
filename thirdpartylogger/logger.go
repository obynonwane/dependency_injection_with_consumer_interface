package thirdpartylogger

import (
	"log"
)

type ThirdPartyLogger struct{}

func NewThirdPartyLogger() *ThirdPartyLogger {
	return &ThirdPartyLogger{}
}

func (l *ThirdPartyLogger) Log(message string) {
	log.Println(message)
}
