package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"go-telnet-chat/commands"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	shellHandler := telsh.NewShellHandler()
	shellHandler.WelcomeMessage = "Welcome to GoChat! type '/help' for a list of commands."

	registerCommand(commands.HelpCommand{}, shellHandler)

	serverAddress := os.Getenv("GO_CHAT_ADDR")

	fmt.Println("GO CHAT SERVER STARTED")
	fmt.Println("Address is ", serverAddress )

	if err := telnet.ListenAndServe(serverAddress, shellHandler); nil != err {
		panic(err)
	}
}

func registerCommand(command commands.Command, shellHandler *telsh.ShellHandler) {
	_ = shellHandler.Register("/help", telsh.ProducerFunc(command.RegisterHandler))
}
