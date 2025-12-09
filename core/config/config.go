package config

import (
	"anno-modmanager/core/helpers"
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type AMMConfigData struct {
	ModFolder       string `json:"modfolder"`
	UserApiKey      string `json:"apikey"`
	UserApiEndpoint string `json:"apiendpoint"`
}

type AMMConfig struct {
	ctx    context.Context
	config *AMMConfigData
}

func NewAMMConfig() *AMMConfig {
	return &AMMConfig{}
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

func (c *AMMConfig) InitAMMConfig(ctx context.Context) {
	c.ctx = ctx
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
		log.Fatal("Saving config.json failed", err)
	}
	c.config = &cd
	// TODO refresh config in other app parts
}

func (c *AMMConfig) SelectAnnoModsFolder() string {
	// TODO
	userhome, _ := os.UserHomeDir()
	dialogOptions := runtime.OpenDialogOptions{
		DefaultDirectory: userhome,
		DefaultFilename:  "",
		Title:            "Select Anno 1800 Mods Folder",
	}
	folder, err := runtime.OpenDirectoryDialog(c.ctx, dialogOptions)
	if err != nil {
		log.Println("Anno mods folder selection failed", err)
		return ""
	}
	return folder
}
