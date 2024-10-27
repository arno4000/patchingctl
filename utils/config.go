package utils

import (
	"github.com/spf13/viper"
)

func LoadConfig(configFile string) error {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
