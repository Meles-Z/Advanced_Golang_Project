package configs

import "github.com/spf13/viper"

type Confgration struct {
	DB     DatabaseConfigruration
	Server Server
	Auth   Websecret
}
type Server struct {
	Host string
	Port int
}

type DatabaseConfigruration struct {
	Host     string
	Name     string
	Port     int
	Username string
	Password string
}

type Websecret struct {
	Secrete string
}

func LoadConfig() (*Confgration, error) {
	var cfg Confgration
	var err error
	viper.AddConfigPath("../")
	viper.SetConfigName(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.BindEnv("db_host", "DB_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_name", "DB_NAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_port", "DB_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_username", "DB_USERNAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db_password", "DB_PASSWORD"); err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
