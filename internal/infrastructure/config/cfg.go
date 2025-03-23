package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	Postgres Postgres
}

func New(path string) (Config, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	return Config{
		Server: Server {
			Host: viper.GetString("server.host"),
		},
		Postgres: Postgres{
			Host: viper.GetString("postgres.host"),
			Port: viper.GetString("postgres.port"),
			Username: viper.GetString("postgres.username"),
			Database: viper.GetString("postgres.dbname"),
			Password: viper.GetString("postgres.password"),
		},
	}, nil
}