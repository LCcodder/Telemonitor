package config

import (
	"bytes"
	"encoding/json"
	"os"

	"slices"
)

type Config struct {
	AuthParams struct {
		Whitelist []string `json:"whitelist"`
		Admin     string   `json:"admin"`
	} `json:"__auth"`
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

func WriteConfig(newCfg Config) (bool, error) {
	file, openErr := os.OpenFile(PathFromEntrypoint, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if openErr != nil {
		return false, openErr
	}
	defer file.Close()

	writeData := new(bytes.Buffer)
	encodeErr := json.NewEncoder(writeData).Encode(newCfg)
	if encodeErr != nil {
		return false, encodeErr
	}

	writeErr := os.WriteFile(PathFromEntrypoint, writeData.Bytes(), 0777)
	if writeErr != nil {
		return false, writeErr
	}

	return true, nil
}
