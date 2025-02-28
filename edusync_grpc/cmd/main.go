package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/meles-z/edusysnc_grpc/configs"
	"github.com/meles-z/edusysnc_grpc/internal/app"
)

func main() {

	// if we retun this in local machine we have to load env.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load config file: %s", err.Error())
	}
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Could not load config file:", err)
	}
	server := app.NewServer(*cfg)
	err = server.Start()
	if err != nil {
		log.Fatalf("Server failed to start: %s", err.Error())
	}
}
