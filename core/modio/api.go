package modio

import (
	"anno-modmanager/core/config"
	"anno-modmanager/core/events"

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

func (api *AMMModioApi) InitModioApi() {
	api.app.Event.On(string(events.REFRESH_CONFIG), func(event *application.CustomEvent) {
		api.app.Logger.Debug("config:refresh event handler in modio api called")
		var newConfig = event.Data.(*config.AMMConfigData)
		api.app.Logger.Debug("CONFIG UserApiKey: %v, UserApiEndpoint: %v", newConfig.UserApiKey, newConfig.UserApiEndpoint)
		api.apikey = newConfig.UserApiKey
		api.baseurl = newConfig.UserApiEndpoint
	})
}

// TODO implement modio api
