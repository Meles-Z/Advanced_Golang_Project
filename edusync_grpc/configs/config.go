package configs

import "github.com/spf13/viper"

type Confgration struct {
	DB     DatabaseConfigruration
	Server ServerConfigration
	Auth   Websecret
}
type ServerConfigration struct {
	Host      string
	Port      int
	GRPC_PORT int
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
	if err = viper.BindEnv("server.grpc_port", "GRPC_PORT"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.name", "DB_NAME"); err != nil {
		return nil, err
	}

	if err = viper.BindEnv("db.host", "DB_HOST"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.port", "DB_PORT"); err != nil {
		return nil, err
	}

	if err = viper.BindEnv("db.username", "DB_USERNAME"); err != nil {
		return nil, err
	}
	if err = viper.BindEnv("db.password", "DB_PASSWORD"); err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
