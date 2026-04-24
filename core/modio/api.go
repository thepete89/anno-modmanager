package modio

import (
	//"anno-modmanager/core/config"
	//"anno-modmanager/core/events"
	//"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type AMMModioApi struct {
	app     *application.App
	baseurl string
	apikey  string
}

func NewModioApi(app *application.App) *AMMModioApi {
	return &AMMModioApi{app: app}
}

func (a *AMMModioApi) InitModioApi() {
	// TODO changeme
	// runtime.EventsOn(a.ctx, string(events.REFRESH_CONFIG), func(data ...any) {
	// 	log.Println("REFRESH_CONFIG event handler in modio api called")
	// 	var config = data[0].(*config.AMMConfigData)
	// 	log.Printf("CONFIG UserApiKey: %v, UserApiEndpoint: %v\n", config.UserApiKey, config.UserApiEndpoint)
	// 	a.apikey = config.UserApiKey
	// 	a.baseurl = config.UserApiEndpoint
	// })
}

// TODO implement modio api
