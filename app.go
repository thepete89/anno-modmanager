package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

// App struct
type App struct {
	ctx          context.Context
	configLoaded bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.configLoaded = false
	// load app config
	configRoot, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := fmt.Sprintf("%s%c%s%c%s", configRoot, os.PathSeparator, "anno-modmanager", os.PathSeparator, "config.json")
	fmt.Println("CONFIG PATH: ", configPath)
	// TODO load file, initialize config
}

// method to check if the config has been loaded
func (a *App) IsConfigLoaded() bool {
	return a.configLoaded
}
