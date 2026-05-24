package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds application configuration
type Config struct {
	Database DatabaseConfig `json:"database"`
	Server   ServerConfig   `json:"server"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string `json:"port"`
}

// Load loads configuration from config.json
func Load() *Config {
	// Try different paths for config.json
	possiblePaths := []string{
		"resources/config.json",
		"./resources/config.json",
		"../resources/config.json",
		"config.json",
	}

	var configPath string
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			configPath = path
			break
		}
	}

	if configPath == "" {
		panic("config.json not found in resources/ directory")
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic("Failed to read config.json: " + err.Error())
	}

	// Parse JSON
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic("Failed to parse config.json: " + err.Error())
	}

	return &cfg
}

// LoadFromFile loads configuration from specified file path
func LoadFromFile(filePath string) *Config {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic("Failed to parse config file: " + err.Error())
	}

	return &cfg
}

// GetDSN returns database connection string
func (c *DatabaseConfig) GetDSN() string {
	if c.Password != "" {
		return "host=" + c.Host + " port=" + c.Port + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " sslmode=" + c.SSLMode
	}
	return "host=" + c.Host + " port=" + c.Port + " user=" + c.User + " dbname=" + c.DBName + " sslmode=" + c.SSLMode
}

// Save saves configuration to config.json file
func (c *Config) Save(filePath string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

