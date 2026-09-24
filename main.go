package main

import (
	"embed"
	"log/slog"

	"anno-modmanager/core/config"
	"anno-modmanager/core/modio"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name: "Anno Mod Manager",
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		LogLevel: slog.LevelDebug,
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Anno Mod Manager",
		Width:  1024,
		Height: 800,
	})

	// setup services
	conf := config.NewAMMConfig()
	api := modio.NewModioApi(conf)
	app.RegisterService(application.NewService(conf))
	app.RegisterService(application.NewService(api))

	// init config
	conf.InitAMMConfig()

	err := app.Run()
	if err != nil {
		println("Error:", err.Error())
	}
}
