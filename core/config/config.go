package config

import (
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
	config *AMMConfigData
}

func NewAMMConfig() *AMMConfig {
	return &AMMConfig{}
}

func (c *AMMConfig) openOrCreateConfigfile() *os.File {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		application.Get().Logger.Error("User config dir not found", "error", err)
		os.Exit(1)
	}
	configFolder := filepath.Join(configRoot, "anno-modmanager")
	err = os.MkdirAll(configFolder, 0750)
	if err != nil {
		application.Get().Logger.Error("Could not create anno modmanager config dir", "error", err)
		os.Exit(1)
	}
	configFile := filepath.Join(configFolder, "config.json")
	application.Get().Logger.Debug("config loaded from", "path", configFile)
	cf := helpers.OpenOrCreateFile(configFile)
	return cf
}

func (c *AMMConfig) loadOrCreateConfig() {
	cf := c.openOrCreateConfigfile()
	defer helpers.CloseFile(cf)
	configData, err := helpers.LoadOrInitializeFromJsonFile[AMMConfigData](cf)
	if err != nil {
		application.Get().Logger.Error("loading or creating config.json failed", "error", err)
		os.Exit(1)
	}
	c.config = configData
}

func (c *AMMConfig) InitAMMConfig() {
	c.loadOrCreateConfig()
}

func (c *AMMConfig) GetConfigData() AMMConfigData {
	return *c.config
}

func (c *AMMConfig) SaveConfigData(cd AMMConfigData) {
	cf := c.openOrCreateConfigfile()
	defer helpers.CloseFile(cf)
	err := helpers.SaveToJsonFile(cf, &cd)
	if err != nil {
		application.Get().Logger.Error("Saving config.json failed", "error", err)
		os.Exit(1)
	}
	c.config = &cd
}

func (c *AMMConfig) SelectAnnoModsFolder() string {
	// TODO
	userhome, _ := os.UserHomeDir()
	folder, err := application.Get().Dialog.OpenFile().
		SetTitle("Select Mods Folder").
		SetDirectory(userhome).
		CanChooseDirectories(true).
		CanChooseFiles(false).
		PromptForSingleSelection()
	if err != nil {
		application.Get().Logger.Error("Anno mods folder selection failed", "error", err)
		return ""
	}
	return folder
}
