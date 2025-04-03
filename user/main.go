package main

import (
	"fmt"
	"log"
	"os"
	//"net"
	user_grpc "github.com/elzamin/air_tickets/proto/gen/go"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/config"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/db"
	user_repo "github.com/elzamin/air_tickets/user/internal/repository/user"
	//"github.com/elzamin/air_tickets/user/internal/test"
	"google.golang.org/grpc"
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
	_ = userRepository

	//test.TestUserDb(userRepository)

	a := user_grpc.GetUserRequest{
		Id: "1",
	}
	fmt.Println(a.Id)

	//lis, err := net.Listen("tcp", fmt.Sprint("localhost:8080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer, err:= grpc.NewClient("passthrough:///localhost:8081")
	if err != nil {
		log.Fatal("Failed to dial the server", err)
	}
	defer grpcServer.Close()

	user_grpc.NewUserClient(grpcServer)
	

	// log.Fatal(http.ListenAndServe(cfg.Server.Host, nil))
}
