package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

// helper to load or create config
func (a *App) loadOrCreateConfig() {
	// load app config
	configRoot, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	configFolder := filepath.Join(configRoot, "anno-modmanager")
	err = os.MkdirAll(configFolder, 0644)
	if err != nil {
		log.Fatal(err)
	}
	configFile := filepath.Join(configFolder, "config.json")
	fmt.Println("CONFIG PATH: ", configFile)
	// TODO load file, initialize config
	cf, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// TODO
	if err := cf.Close(); err != nil {
		log.Fatal(err)
	}
	//a.configLoaded = true
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.configLoaded = false
	a.loadOrCreateConfig()

}

// method to check if the config has been loaded
func (a *App) IsConfigLoaded() bool {
	return a.configLoaded
}
