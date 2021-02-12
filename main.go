package main

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"go-telnet-chat/commands"
)

func main() {

	shellHandler := telsh.NewShellHandler()
	shellHandler.WelcomeMessage = "Welcome to GoChat! type '/help' for a list of commands."

	registerCommand(commands.HelpCommand{}, shellHandler)

	addr := ":5555"
	if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
		panic(err)
	}
}

func registerCommand(command commands.Command, shellHandler *telsh.ShellHandler) {
	_ = shellHandler.Register("/help", telsh.ProducerFunc(command.RegisterHandler))
}
