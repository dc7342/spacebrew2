package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Telegram struct {
		Token   string
		Timeout int
		Channel string
	}

	Bot struct {
		Title string
	}

	DB struct {
		Host    string
		Port    int
		User    string
		Pass    string
		DBName  string
		SSLMode string
	}
}

func Read() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/goalkeeper/")
	viper.AddConfigPath("$HOME/.goalkeeper")
	viper.AddConfigPath("/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	var c Config

	err = viper.Unmarshal(&c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
