package cli

import (
	"github.com/mattn/go-isatty"
	"os"
)

// Isatty check current device is terminal
func Isatty() bool {
	if isatty.IsTerminal(os.Stdin.Fd()) {
		return true
	}
	if isatty.IsCygwinTerminal(os.Stdin.Fd()) {
		return true
	}
	return false
}
