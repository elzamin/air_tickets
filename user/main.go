package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"net"

	"google.golang.org/grpc"

	pb "github.com/elzamin/air_tickets/proto/gen/go"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/config"
	"github.com/elzamin/air_tickets/user/internal/infrastructure/db"
	"github.com/elzamin/air_tickets/user/internal/repository"
	"github.com/elzamin/air_tickets/user/internal/test"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) CreateUser (_ context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Received: %v", in.GetUser())
	return &pb.CreateUserResponse{Error: &pb.Error{Message: "Hello world"}}, nil
}

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
	if (1 == 1) {
		test.TestUserDb(ctx, userRepository)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// log.Fatal(http.ListenAndServe(cfg.Server.Host, nil))
}
