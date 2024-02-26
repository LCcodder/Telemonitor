package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AuthParams struct {
		Whitelist []string `json:"whitelist"`
	} `json:"__auth"`
}

var Cfg Config

// Reads the json file, mutates global variable and returns its pointer
func LoadConfig(currentPath string) *Config {
	file, err := os.Open(currentPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parser := json.NewDecoder(file)
	parser.Decode(&Cfg)
	return &Cfg
}
