package commands

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
)

type ChatCommand struct{}

func (ChatCommand) GetShortcut() string {
	return "SECRET_CHAT_COMMAND_STRING"
}

func (ChatCommand) RegisterHandler(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(ChatCommand{}.runCommand)
}

func (ChatCommand) runCommand(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "Chat is working bro!")
	return nil
}
