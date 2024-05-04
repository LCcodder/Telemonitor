package config

import (
	"encoding/json"
	"os"

	"slices"
)

type Config struct {
	AuthParams struct {
		Whitelist []string `json:"whitelist"`
		Admin     string   `json:"admin"`
	} `json:"__auth"`
	Messages struct {
		Welcome string `json:"welcome"`
	} `json:"messages"`
}

var (
	Cfg                Config
	PathFromEntrypoint string
)

// Reads the json file, mutates global variable and returns its pointer
func LoadConfig() *Config {
	file, err := os.Open(PathFromEntrypoint)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parser := json.NewDecoder(file)
	parser.Decode(&Cfg)

	// kinda slow, but emmits only once, at start, so i don't actualy care
	slices.Sort(Cfg.AuthParams.Whitelist)

	return &Cfg
}
