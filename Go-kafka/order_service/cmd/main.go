package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/meles-z/go-kafa/order_service/configs"
	"github.com/meles-z/go-kafa/order_service/internal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file.", err)
	}
	server, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Could not load enf file", err)
	}
	inter := internal.NewServer(*server)
	log.Fatal(inter.Start())
}
