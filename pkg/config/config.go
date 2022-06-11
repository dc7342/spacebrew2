package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Telegram struct {
		Token           string
		Timeout         int
		ChannelUsername string
	}

	Page struct {
		TasksPerPage int
	}

	Text struct {
		Title                string
		WelcomeMessage       string
		UnknownMessage       string
		NewTaskName          string
		NewDescriptionName   string
		NewTaskDone          string
		EditTaskChoice       string
		EditDescription      string
		EditName             string
		Confirmation         string
		ConfirmationNegative string
		Done                 string
	}

	Button struct {
		Menu            string
		AllTasks        string
		AddTask         string
		Cancel          string
		EditDescription string
		EditTitle       string
		Yes             string
		No              string
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
