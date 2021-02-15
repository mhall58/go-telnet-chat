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

	file, err := os.OpenFile(os.Getenv("GO_CHAT_LOG_FILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	s := Server{Addr: os.Getenv("GO_CHAT_ADDR")}
	s.Start()
}
