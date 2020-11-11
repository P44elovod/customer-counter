package util

import "github.com/spf13/viper"

type Config struct {
	BindAddr string `mapstructure:"BIND_ADDR"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := NewConfig()
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil

}
