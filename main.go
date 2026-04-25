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

	// TODO switch to service structure
	config := config.NewAMMConfig(app)
	modioapi := modio.NewModioApi(app)
	app.RegisterService(application.NewService(config))
	app.RegisterService(application.NewService(modioapi))
	config.InitAMMConfig()

	err := app.Run()
	if err != nil {
		println("Error:", err.Error())
	}
}
