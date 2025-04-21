package main

import (
	"context"
	"log"
	"os"

	"github.com/elzamin/air_tickets/user/internal/api"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/config"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/db"
	userRepo "github.com/elzamin/air_tickets/user/internal/repository"
	userService "github.com/elzamin/air_tickets/user/internal/service"
)

func main() {
	cfg, err := config.New("config/config." + os.Getenv("ENV") + ".yml")
	if err != nil {
		log.Fatal("Failed to init a config", err)
	}
	ctx := context.Background()

	dbConnection, err := db.NewPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal("Failed to create postgres connections", err)
	}

	userRepository := userRepo.New(dbConnection)
	userService := userService.New(userRepository)
	userServer := api.New(userService)
	
	//test
	if (0 == 1) {
		userRepo.TestUserDb(ctx, userRepository)
	}

	api.RunGRPCServer(cfg.Server.Port, userServer)
	
}
