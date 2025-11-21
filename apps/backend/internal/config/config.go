package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig `yaml:"app"`
}

type AppConfig struct {
	Port  int  `yaml:"port"`
	Debug bool `yaml:"debug"`
}

func LoadConfig(env string) (*Config, error) {
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../config")
	viper.SetConfigType("yaml")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal config: %w", err)
	}
	return &config, nil
}
