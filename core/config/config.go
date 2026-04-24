package config

import (
	//"anno-modmanager/core/events"
	"anno-modmanager/core/helpers"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type AMMConfigData struct {
	ModFolder       string `json:"modfolder"`
	UserApiKey      string `json:"apikey"`
	UserApiEndpoint string `json:"apiendpoint"`
}

type AMMConfig struct {
	app    *application.App
	config *AMMConfigData
}

func NewAMMConfig(app *application.App) *AMMConfig {
	return &AMMConfig{app: app}
}

func (c *AMMConfig) openOrCreateConfigfile() *os.File {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	configFolder := filepath.Join(configRoot, "anno-modmanager")
	err = os.MkdirAll(configFolder, 0750)
	if err != nil {
		log.Fatal(err)
	}
	configFile := filepath.Join(configFolder, "config.json")
	log.Println("CONFIG PATH: ", configFile)
	cf := helpers.OpenOrCreateFile(configFile)
	return cf
}

func (c *AMMConfig) loadOrCreateConfig() {
	cf := c.openOrCreateConfigfile()
	defer helpers.CloseFile(cf)
	configData, err := helpers.LoadOrInitializeFromJsonFile[AMMConfigData](cf)
	if err != nil {
		log.Fatal("loading or creating config.json failed", err)
	}
	c.config = configData
}

func (c *AMMConfig) InitAMMConfig() {
	c.loadOrCreateConfig()
	// TODO changeme
	//runtime.EventsEmit(c.ctx, string(events.REFRESH_CONFIG), c.config)
}

func (c *AMMConfig) GetConfigData() AMMConfigData {
	return *c.config
}

func (c *AMMConfig) SaveConfigData(cd AMMConfigData) {
	cf := c.openOrCreateConfigfile()
	defer helpers.CloseFile(cf)
	err := helpers.SaveToJsonFile(cf, &cd)
	if err != nil {
		log.Fatal("Saving config.json failed", err)
	}
	c.config = &cd
	// TODO changeme
	// runtime.EventsEmit(c.ctx, string(events.REFRESH_CONFIG), c.config)
}

func (c *AMMConfig) SelectAnnoModsFolder() string {
	// TODO
	userhome, _ := os.UserHomeDir()
	folder, err := c.app.Dialog.OpenFile().
		SetTitle("Select Mods Folder").
		SetDirectory(userhome).
		CanChooseDirectories(true).
		CanChooseFiles(false).
		PromptForSingleSelection()
	if err != nil {
		log.Println("Anno mods folder selection failed", err)
		return ""
	}
	return folder
}
