package api

import interface_log "github.com/obynonwane/di_with_packages/interface"

type APIService struct {
	Logger interface_log.Logger
}

func NewAPIService(logger interface_log.Logger) *APIService {
	return &APIService{
		Logger: logger,
	}
}

func (api *APIService) LogMessage() {
	api.Logger.Log("Hello welcome")
}
