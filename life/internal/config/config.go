package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Username string
	Password string
	ApiUrl   string
}

type Auth struct {
	Jwt     string    `json:"jwt"`
	Expires time.Time `json:"expires"`
}

func (a Auth) isExpired() bool {
	return a.Expires.Before(time.Now())
}

func ReadLifeConfig() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home directory: %v", err)
		return Config{}, err
	}
	configPath := filepath.Join(homeDir, ".life.yaml")

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	// Read in the .env file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading .env file: %v", err)
		return Config{}, err
	}

	apiUrl := viper.GetString("API_URL")
	if apiUrl == "" {
		fmt.Printf("No API_URL found, defaulting to localhost")
		apiUrl = "http://localhost:8080"
	}

	username := viper.GetString("username")
	password := viper.GetString("password")
	if username == "" || password == "" {
		fmt.Println("Config missing: username and/or password")
		return Config{}, err
	}

	return Config{username, password, apiUrl}, nil
}

func ReadAuth(fp string) (Auth, error) {
	file, err := os.Open(fp)
	if err != nil {
		return Auth{}, err
	}

	contents, err := io.ReadAll(file)
	if err != nil {
		return Auth{}, err
	}

	auth := Auth{}
	err = json.Unmarshal(contents, &auth)
	if err != nil {
		return Auth{}, err
	}

	return auth, nil
}
