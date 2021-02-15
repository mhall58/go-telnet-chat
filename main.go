package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logDirectory := os.Getenv("GO_CHAT_LOG_FILE")

	s := Server{Addr: os.Getenv("GO_CHAT_ADDR")}
	s.Start()
}
