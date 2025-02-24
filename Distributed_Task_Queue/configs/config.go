package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB   DatabaseConfigration
	User Authentication
}
type DatabaseConfigration struct {
	Host     string
	Name     string
	Port     int
	UserName string
	Password string
}
type Authentication struct {
	Secret string
}
type Server struct {
	Port string
	Host string
}

func LoadConfig() (*Configuration, error) {
	var cfg Configuration
	var err error

	viper.AddConfigPath("../")
	viper.SetConfigName(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.BindEnv("db_host", "DB_HOST")
	if err != nil {
		fmt.Println("Err:", err)
		return nil, err
	}
	err = viper.BindEnv("db_name", "DB_NAME")
	if err != nil {
		fmt.Println("Err:", err)
		return nil, err
	}
	err = viper.BindEnv("db_port", "DB_PORT")
	if err != nil {
		fmt.Println("Err:", err)
		return nil, err
	}
	err = viper.BindEnv("db_username", "DB_USERNAME")
	if err != nil {
		fmt.Println("Err:", err)
		return nil, err
	}
	err = viper.BindEnv("db_password", "DB_PASSWORD")
	if err != nil {
		fmt.Println("Err:", err)
		return nil, err
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	return &cfg, nil
}
