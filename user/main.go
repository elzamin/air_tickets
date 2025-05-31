package main

import (
	"log"
	"os"

	"github.com/elzamin/air_tickets/user/internal/api"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/config"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/db"
	"github.com/elzamin/air_tickets/user/internal/repository"
	"github.com/elzamin/air_tickets/user/internal/service"
)

func main() {
	cfg, err := config.New("config/config." + os.Getenv("ENV") + ".yml")
	if err != nil {
		log.Fatal("Failed to init a config", err)
	}

	dbConnection, err := db.NewPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal("Failed to create postgres connections", err)
	}

	userRepository := repository.New(dbConnection)
	userService := service.New(userRepository)
	userServer := api.New(userService)

	api.RunGRPCServer(cfg.Server.Port, userServer)
}
