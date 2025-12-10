package main

import (
	"context"
	"embed"

	"anno-modmanager/core/config"
	"anno-modmanager/core/events"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	config := config.NewAMMConfig()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Anno Modmanager",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.InitApp(ctx)
			config.InitAMMConfig(ctx)
		},
		Bind: []any{
			app,
			config,
		},
		EnumBind: []any{
			events.AMMEvents,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
