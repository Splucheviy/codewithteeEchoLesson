package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv    string `mapstructure:"APP_ENV"`
	AppPort   string `mapstructure:"APP_PORT"`
	AppAdress string `mapstructure:"APP_ADDRESS"`
	DBName    string `mapstructure:"DB_NAME"`
	DBHost    string `mapstructure:"DB_HOST"`
	DBPort    string `mapstructure:"DB_PORT"`
	DBUser    string `mapstructure:"DB_USER"`
	DBPass    string `mapstructure:"DB_PASS"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Panic("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Printf("The App is running in development env on %s", env.AppAdress)
	}

	return &env
}
