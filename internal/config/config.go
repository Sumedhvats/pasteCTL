package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const configFileName = "config"
const configFolderName = "pastectl"


const defaultBackendURL = "https://api.paste.sumedh.app"
const defaultFrontendURL = "https://paste.sumedh.app/paste"

func InitConfig() {
    home, err := os.UserHomeDir()
    if err != nil {
        fmt.Printf("Could not find home directory: %v\n", err)
        os.Exit(1)
    }

    configPath := filepath.Join(home, ".config", configFolderName)
    if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
        fmt.Printf(" Could not create config directory: %v\n", err)
        os.Exit(1)
    }

    configFilePath := filepath.Join(configPath, configFileName+".yaml")

    // Tell Viper where the config file is
    viper.SetConfigFile(configFilePath)
    viper.SetConfigType("yaml")

    // Set default values
    viper.SetDefault("backend_url", defaultBackendURL)
    viper.SetDefault("frontend_url", defaultFrontendURL)


}
func Get(key string) string {
	return viper.GetString(key)
}

func Set(key, value string) error {
    home, err := os.UserHomeDir()
    if err != nil {
        return fmt.Errorf("could not find home directory: %w", err)
    }

    configPath := filepath.Join(home, ".config", configFolderName)
    if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
        return fmt.Errorf("could not create config directory: %w", err)
    }

    configFilePath := filepath.Join(configPath, configFileName+".yaml")

    // Read existing config if it exists
    viper.SetConfigFile(configFilePath)
    _ = viper.ReadInConfig() // ignore error if file doesn't exist

    // Set new key in memory
    viper.Set(key, value)

    // Write all keys back to file
    return viper.WriteConfigAs(configFilePath)
}
