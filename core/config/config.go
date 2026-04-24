package config

import (
	"anno-modmanager/core/events"
	"anno-modmanager/core/helpers"
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
		c.app.Logger.Error("User config dir not found", "error", err)
		os.Exit(1)
	}
	configFolder := filepath.Join(configRoot, "anno-modmanager")
	err = os.MkdirAll(configFolder, 0750)
	if err != nil {
		c.app.Logger.Error("Could not create anno modmanager config dir", "error", err)
		os.Exit(1)
	}
	configFile := filepath.Join(configFolder, "config.json")
	c.app.Logger.Debug("config loaded from", "path", configFile)
	cf := helpers.OpenOrCreateFile(configFile)
	return cf
}

func (c *AMMConfig) loadOrCreateConfig() {
	cf := c.openOrCreateConfigfile()
	defer helpers.CloseFile(cf)
	configData, err := helpers.LoadOrInitializeFromJsonFile[AMMConfigData](cf)
	if err != nil {
		c.app.Logger.Error("loading or creating config.json failed", "error", err)
		os.Exit(1)
	}
	c.config = configData
}

func (c *AMMConfig) InitAMMConfig() {
	c.loadOrCreateConfig()
	c.app.Event.Emit(string(events.REFRESH_CONFIG), c.config)
}

func (c *AMMConfig) GetConfigData() AMMConfigData {
	return *c.config
}

func (c *AMMConfig) SaveConfigData(cd AMMConfigData) {
	cf := c.openOrCreateConfigfile()
	defer helpers.CloseFile(cf)
	err := helpers.SaveToJsonFile(cf, &cd)
	if err != nil {
		c.app.Logger.Error("Saving config.json failed", "error", err)
		os.Exit(1)
	}
	c.config = &cd
	c.app.Event.Emit(string(events.REFRESH_CONFIG), c.config)
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
		c.app.Logger.Error("Anno mods folder selection failed", "error", err)
		return ""
	}
	return folder
}
