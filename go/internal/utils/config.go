package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config interface {
}

func CastConfig(configFilePath string, configType Config) error {
	file, err := os.Open(configFilePath)
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(configType); err != nil {
		return err
	}

	return nil
}
