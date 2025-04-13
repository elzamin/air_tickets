package main

import (
	"log"
	"os"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/elzamin/air_tickets/proto/gen/go"
	"github.com/elzamin/air_tickets/gateway/internal/infrastructure/config"
)

func main() {
	cfg, err := config.New("config/config." + os.Getenv("ENV") + ".yml")
	if err != nil {
		log.Fatal("Failed to init a config", err)
	}

	conn, err := grpc.NewClient("localhost:" + cfg.Client_user.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateUser(ctx, &pb.CreateUserRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetError().Message)
}