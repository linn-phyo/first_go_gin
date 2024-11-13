package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBName        string `mapstructure:"DB_NAME"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JwtSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "SERVER_ADDRESS", "JWT_SECRET_KEY",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env : ", err)
		return config, err
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		log.Println("The App is running in development env")
		return config, err
	}

	return config, nil
}
