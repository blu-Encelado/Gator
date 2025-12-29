package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("fail to GetHomeDir: %w", err)
	}

	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, fmt.Errorf("fail to ReadFile: %w", err)
	}
	newCfg := Config{}
	err = json.Unmarshal(bytes, &newCfg)
	if err != nil {
		return Config{}, fmt.Errorf("fail to unMarshal: %w", err)
	}

	return newCfg, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	err := write(*cfg)
	if err != nil {
		return fmt.Errorf("fail to write: %w", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := homeDir + "/" + configFileName

	return fullPath, nil
}

func write(cfg Config) error {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	err = os.WriteFile(cfgPath, bytes, 0600)
	if err != nil {
		return err
	}
	return nil
}
