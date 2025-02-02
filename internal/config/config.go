package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	cfg := Config{}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	rawConfig, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	if err = json.Unmarshal(rawConfig, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil

}

func (cfg Config) SetUser(username string) error {

	if len(username) == 0 {
		return fmt.Errorf("non-empty username must be supplied")
	}

	cfg.CurrentUserName = username

	err := cfg.write()
	if err != nil {
		return err
	}

	return nil

}

func (cfg Config) write() error {

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	configJSON, err := json.Marshal(cfg)

	if err != nil {
		return err
	}

	os.WriteFile(configFilePath, configJSON, 0644)

	return nil

}

func getConfigFilePath() (string, error) {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil

}
