package main

import (
	"log"
	"os"

	"github.com/elzamin/air_tickets/gateway/internal/infrastructure/config"
)

func main() {
	_, err := config.New("config/config." + os.Getenv("ENV") + ".yml")
	if err != nil {
		log.Fatal("Failed to init a config", err)
	}
}