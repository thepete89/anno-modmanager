package modio

import (
	"anno-modmanager/core/config"
)

type AMMModioApi struct {
	conf *config.AMMConfig
}

func NewModioApi(conf *config.AMMConfig) *AMMModioApi {
	return &AMMModioApi{conf: conf}
}

func (api *AMMModioApi) InitModioApi() {

}

// TODO implement modio api
