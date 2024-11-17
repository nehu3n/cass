package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

func GetConfigFile() (string, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%sconfig.toml", configPath, string(os.PathSeparator)), nil
}

func CreateConfigFile() error {
	configFilePath, err := GetConfigFile()
	if err != nil {
		return err
	}

	err = CreateConfigPath()
	if err != nil {
		return err
	}
	
	if existsPath(configFilePath) {
		return nil
	}

	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func WriteConfigFile() error {
	config := ConfigStruct{
		DryRun: true,
	}

	b, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	if err := CreateConfigFile(); err != nil {
		return err
	}

	configFilePath, err := GetConfigFile()
	if err != nil {
		return err
	}

	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	file.Write(b)

	return nil
}
