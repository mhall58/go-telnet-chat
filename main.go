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

	chatHandler := NewChatHandler()

	// Register Custom Command Here:
	registerCommand(commands.HelpCommand{}, chatHandler)
	registerCommand(commands.ChatCommand{}, chatHandler)

	// Load Configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverAddress := os.Getenv("GO_CHAT_ADDR")
	fmt.Println("GO CHAT SERVER STARTED")
	fmt.Println("Address is ", serverAddress)

	if err := telnet.ListenAndServe(serverAddress, chatHandler); nil != err {
		panic(err)
	}
}

func registerCommand(command commands.Command, shellHandler *ShellHandler) {
	_ = shellHandler.Register(command.GetShortcut(), telsh.ProducerFunc(command.RegisterHandler))
}
