package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/meles-z/edusysnc_grpc/configs"
	"github.com/meles-z/edusysnc_grpc/internal/app"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func main() {
	fmt.Println("gRPC server running ...")

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
