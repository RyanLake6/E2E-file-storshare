package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	BaseURL string `json:"baseURL"`
	Token   string `json:"token"`
}

var (
	AppConfig  = Config{}
	configDir  = filepath.Join(os.Getenv("USERPROFILE"), ".e2e-file-storshare-cli")
	configFile = filepath.Join(configDir, "config.json")
	mutex      sync.RWMutex
)

func init() {
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			fmt.Println("Error creating config directory:", err)
		}
	}
}

// SaveConfig saves the configuration to a file.
func SaveConfig() error {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(AppConfig)
}

// LoadConfig loads the configuration from a file.
func LoadConfig() error {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := os.Open(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No config file is okay
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&AppConfig)
}

func SetBaseURL(baseURL string) {
	mutex.Lock()
	defer mutex.Unlock()
	AppConfig.BaseURL = baseURL
}

func GetBaseURL() string {
	mutex.RLock()
	defer mutex.RUnlock()
	return AppConfig.BaseURL
}

func SetToken(token string) {
	mutex.Lock()
	defer mutex.Unlock()
	AppConfig.Token = token
}

func GetToken() string {
	mutex.RLock()
	defer mutex.RUnlock()
	return AppConfig.Token
}
