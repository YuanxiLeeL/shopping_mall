package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn string
		maxIdleConns int
   		maxOpenConns int
   		maxlifetime int
	}
}

var Appconfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	Appconfig = &Config{}

	if err := viper.Unmarshal(Appconfig); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
	InitDB()
	InitRedis()
	InitCasbin()
	InitValidate()
}
