package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/elzamin/air_tickets/proto/gen/go"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/config"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/db"
	"github.com/elzamin/air_tickets/user/internal/repository"
	"github.com/elzamin/air_tickets/user/internal/test"
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
	ctx := context.Background()

	//test
	if (0 == 1) {
		test.TestUserDb(ctx, userRepository)
	}

	a := pb.GetUserRequest{
		Id: "1",
	}
	fmt.Println(a.Id)

	

	// log.Fatal(http.ListenAndServe(cfg.Server.Host, nil))
}
