package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

func GetConfigPath() (string, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s.cass", dir, string(os.PathSeparator)), nil
}

func CreateConfigPath() error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func existsPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
