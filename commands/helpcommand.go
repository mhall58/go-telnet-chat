package commands

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
)

type HelpCommand struct{}

func (HelpCommand) GetShortcut() string {
	return "/help"
}

func (HelpCommand) RegisterHandler(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(HelpCommand{}.runCommand)
}
func (HelpCommand) runCommand(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	commands := []string{
		"------------------------------------------------------------------\r\n",
		"Commands:\r\n",
		"------------------------------------------------------------------\r\n",
		"-  /list                - List channels\r\n",
		"-  /join <channel>      - join a channel\r\n",
		"-  /leave               - leave a channel\r\n",
		"-  /part                - alias for leave\r\n",
		"-  /handle <new name>   - change your chat handle\r\n",
		"-  /help                - prints this menu\r\n",
		"-  /giffy <keywords>    - inserts a random gif based on keyword\r\n",
		"-  /exit                - ends the session\r\n",
		"------------------------------------------------------------------\r\n",
		"\r\n",
	}

	for _, command := range commands {
		oi.LongWriteString(stdout, command)
	}

	return nil
}
