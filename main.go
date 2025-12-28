package main

import (
	"github.com/obynonwane/di_with_packages/api"
	"github.com/obynonwane/di_with_packages/consolelogger"
)

func main() {
	// instantiate the logger
	logger := consolelogger.NewConsoleLogger()
	service := api.NewAPIService(logger)
	service.LogMessage()

}
