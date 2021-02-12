package commands

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
)

type Command interface {
	// GetShortcut returns the command keyword. i.e '/help'
	GetShortcut() string

	// RegisterHandler binds the command to the session
	RegisterHandler(ctx telnet.Context, name string, args ...string) telsh.Handler

	// runCommand Performs the command
	runCommand(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error
}
