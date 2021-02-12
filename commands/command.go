package commands

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
)

type Command interface {
	GetShortcut() string
	RegisterHandler(ctx telnet.Context, name string, args ...string) telsh.Handler
	runCommand(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error
}
