package configs

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	DB     *gorm.DB
	Server ServerConfig
	Auth   AuthenticationConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

type ServerConfig struct {
	Host string
	Port int
}

type AuthenticationConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {
	var cfg Config
	var err error

	viper.AddConfigPath("./")
	viper.SetConfigName(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.BindEnv("server.host", "SERVER_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("server.port", "SERVER_PORT"); err != nil {
		return nil, err
	}

	if err = viper.BindEnv("db.host", "DB_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.port", "DB_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.name", "DB_NAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.username", "DB_USERNAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.password", "DB_PASSWORD"); err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
