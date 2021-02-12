package main

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"go-telnet-chat/commands"
	"io"
)

/** Interfaces **/

type ChatCommand interface {
	runCommand (stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error
	registerHandler(ctx telnet.Context, name string, args ...string) telsh.Handler
}

func main() {

	shellHandler := telsh.NewShellHandler()
	shellHandler.WelcomeMessage = "Welcome To Go Chat, type '/help' for a list of commands."
	_ = shellHandler.Register("/help", telsh.ProducerFunc(commands.HelpCommand{}.RegisterHandler))



	addr := ":5555"
	if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
		panic(err)
	}
}
