package config

import (
	"github.com/elzamin/air_tickets/user/internal/infrastructure/model"
	"github.com/spf13/viper"
)

func New(path string) (model.Config, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return model.Config{}, err
	}

	return model.Config{
		Server: model.Server{
			Port: viper.GetString("server.port"),
		},
		Postgres: model.Postgres{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetString("postgres.port"),
			Username: viper.GetString("postgres.username"),
			Database: viper.GetString("postgres.dbname"),
			Password: viper.GetString("postgres.password"),
		},
	}, nil
}
