package main

import (
	"context"
	"log"
	"os"

	"github.com/elzamin/air_tickets/gateway/internal/api"
	"github.com/elzamin/air_tickets/gateway/internal/client"
	//"github.com/elzamin/air_tickets/gateway/internal/entity"
	"github.com/elzamin/air_tickets/gateway/internal/infrastructure/config"
	//pb "github.com/elzamin/air_tickets/proto/gen/go"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New("config/config." + os.Getenv("ENV") + ".yml")
	if err != nil {
		log.Fatal("Failed to init a config", err)
	}

	uc := client.ConnectToUserGRPC(cfg.Client_user.Port)
	gateClient := client.New(uc)
	gateServer := api.New(gateClient)
	gateServer.RunHTTPServer(ctx)
}
