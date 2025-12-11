package modio

import (
	"anno-modmanager/core/config"
	"anno-modmanager/core/events"
	"context"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type AMMModioApi struct {
	ctx     context.Context
	baseurl string
	apikey  string
}

func NewModioApi() *AMMModioApi {
	return &AMMModioApi{}
}

func (a *AMMModioApi) InitModioApi(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(a.ctx, string(events.REFRESH_CONFIG), func(data ...any) {
		log.Println("REFRESH_CONFIG event handler in modio api called")
		var config = data[0].(*config.AMMConfigData)
		log.Printf("CONFIG UserApiKey: %v, UserApiEndpoint: %v\n", config.UserApiKey, config.UserApiEndpoint)
		a.apikey = config.UserApiKey
		a.baseurl = config.UserApiEndpoint
	})
}

// TODO implement modio api
