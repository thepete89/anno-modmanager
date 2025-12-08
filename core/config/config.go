package config

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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

// helpers to load or create config
func (c *AMMConfig) closeConfigFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *AMMConfig) loadOrCreateConfig() {
	// load app config
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
	cf, err := os.OpenFile(configFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer c.closeConfigFile(cf)
	// read config
	cfr := bufio.NewScanner(cf)
	var cfc strings.Builder
	for cfr.Scan() {
		cfc.WriteString(cfr.Text())
	}
	// parse config
	configData := AMMConfigData{}
	dec := json.NewDecoder(strings.NewReader(cfc.String()))
	err = dec.Decode(&configData)
	if err == io.EOF {
		log.Println("JSON Config was empty - recreating...")
		configData = AMMConfigData{"", "", ""}
		enc := json.NewEncoder(cf)
		enc.Encode(configData)
	} else if err != nil {
		log.Fatal("JSON Decode error:", err)
	}
	c.config = &configData
}

func (c *AMMConfig) InitAMMConfig(ctx context.Context) {
	c.ctx = ctx
	c.loadOrCreateConfig()
}

func (c *AMMConfig) GetConfigData() AMMConfigData {
	return *c.config
}
