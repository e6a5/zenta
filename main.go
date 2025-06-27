// Zenta - mindfulness for terminal users
// A zen-focused breathing and reflection tool for developers
package main

import (
	"os"
	"path/filepath"

	"github.com/e6a5/zenta/internal/cli"
)

func main() {
	// Get the executable name from argv[0]
	programName := filepath.Base(os.Args[0])

	if len(os.Args) < cli.MinArgs {
		cli.ShowHelp(programName)
		return
	}

	command := os.Args[1]

	switch command {
	case "now":
		cli.HandleNow(os.Args[2:])
	case "reflect":
		cli.HandleReflect(os.Args[2:])
	case "help":
		cli.ShowHelp(programName)
	case "version", "--version", "-v":
		cli.HandleVersion(programName)
	default:
		cli.HandleUnknownCommand(command, programName)
	}
}
