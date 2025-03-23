package main

import (
	"airtickets/internal/entity/model"
	"airtickets/internal/infrastructure/config"
	"airtickets/internal/infrastructure/db"
	user_repo "airtickets/internal/repository/user"
	"context"
	"log"
	"os"
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

	userRepository := user_repo.New(dbConnection)

	ctx := context.Background()
	err = userRepository.TouchTable(ctx)
	if err != nil {
		log.Fatal("Failed to touch table user", err)
	}

	err = userRepository.Create(ctx, model.User{Id: "1", FirstName: "Elzamin", LastName: "Usubaliev"})
	if err != nil {
		log.Fatal("Failed to create user", err)
	}
	user, err := userRepository.Get(ctx, "1")
	if err != nil {
		log.Fatal("Failed to get user", err)
	}
	log.Println(user)


	// log.Fatal(http.ListenAndServe(cfg.Server.Host, nil))
}
