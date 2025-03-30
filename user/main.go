package main

import (
	user_grpc "github.com/elzamin/air_tickets/proto/gen/go/user"
	"github.com/elzamin/air_tickets/user/internal/entity"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/config"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/db"
	user_repo "github.com/elzamin/air_tickets/user/internal/repository/user"
	"context"
	"log"
	"os"
	"fmt"
)

func main() {
	cfg, err := config.New("config/config." + os.Getenv("ENV") + ".yml")
	if err != nil {
		log.Fatal("Failed to init a config", err)
	}

	a := user_grpc.GetUserRequest{
		Id: "1",
	}
	fmt.Println(a.Id)

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

	err = userRepository.Create(ctx, entity.User{Id: "1", FirstName: "Elzamin", LastName: "Usubaliev"})
	if err != nil {
		log.Fatal("Failed to create user", err)
	}

	user, err := userRepository.Get(ctx, "1")
	if err != nil {
		log.Fatal("Failed to get user", err)
	}
	log.Println(user)

	err = userRepository.Update(ctx, entity.User{Id: "1", FirstName: "Vasya", LastName: "Veseliy"})
	if err != nil {
		log.Fatal("Failed to update user", err)
	}

	user, err = userRepository.Get(ctx, "1")
	if err != nil {
		log.Fatal("Failed to get user", err)
	}
	log.Println(user)

	err = userRepository.Create(ctx, entity.User{Id: "2", FirstName: "Elzamin", LastName: "Usubaliev"})
	if err != nil {
		log.Fatal("Failed to create user", err)
	}

	users, _ := userRepository.GetAll(ctx)
	log.Println(users)

	err = userRepository.Delete(ctx, "1")
	if err != nil {
		log.Fatal("Failed to delete user", err)
	}

	// log.Fatal(http.ListenAndServe(cfg.Server.Host, nil))
}
