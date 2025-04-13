package config

import (
	"github.com/elzamin/air_tickets/gateway/internal/infrastructure/model"
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
		Client_user: model.Client{
			Port:     viper.GetString("grpc_user.port"),
		},
		Client_warehouse: model.Client{
			Port:     viper.GetString("grpc_warehouse.port"),
		},
		Client_order: model.Client{
			Port:     viper.GetString("grpc_order.port"),
		},
	}, nil
}
