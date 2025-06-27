// Package cli provides command-line interface handling for zenta.
// It manages command routing, help display, and user interaction.
package cli

import (
	"fmt"
	"os"

	"github.com/e6a5/zenta/internal/breathing"
	"github.com/e6a5/zenta/internal/quotes"
	"github.com/e6a5/zenta/internal/version"
)

const MinArgs = 2

// ShowHelp displays the main help message
func ShowHelp(programName string) {
	fmt.Printf("%s - mindfulness for terminal users\n", programName)
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Printf("  %s now [options]         Take a mindful breathing moment\n", programName)
	fmt.Printf("  %s reflect               End-of-day reflection on thought patterns\n", programName)
	fmt.Printf("  %s help                  Show this help message\n", programName)
	fmt.Println()
	fmt.Println("NOW OPTIONS:")
	fmt.Println("  --quick, -q                 Quick 1-cycle session")
	fmt.Println("  --extended, -e              Extended 5-cycle session")
	fmt.Println("  --silent, -s                Breathing only, skip the quote")
	fmt.Println("  --simple                    Simple line animation (for terminal compatibility)")
	fmt.Println("  --complex                   More complex animation (default except on Apple")
	fmt.Println("                              Terminal, screen, and tmux)")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Printf("  %s now                   Standard 3-cycle breathing session\n", programName)
	fmt.Printf("  %s now --quick           Quick 1-cycle breathing break\n", programName)
	fmt.Printf("  %s now --extended        Extended 5-cycle session\n", programName)
	fmt.Printf("  %s now --silent          Breathing without quote\n", programName)
	fmt.Printf("  %s now --simple          Simple animation (terminal compatibility)\n", programName)
	fmt.Printf("  %s reflect               Gentle end-of-day reflection\n", programName)
	fmt.Println()
	fmt.Println("MINDFUL ALIASES:")
	fmt.Printf("  alias breath='%s now --quick'\n", programName)
	fmt.Printf("  alias breathe='%s now'\n", programName)
	fmt.Printf("  alias reflect='%s reflect'\n", programName)
	fmt.Println()
	fmt.Println("Learn more: https://github.com/e6a5/zenta")
}

// HandleNow handles the 'now' command for breathing sessions
func HandleNow(args []string) {
	session := breathing.NewSession()
	session.ParseArgs(args)

	session.Start()

	if session.ShouldShowQuote() {
		quoteService := quotes.New()
		quote := quoteService.GetRandomQuote()
		quotes.DisplayBeautifully(quote)
	} else {
		breathing.PrintWithPadding("   Carry this calm with you throughout your day 🙏")
	}

	breathing.AddBottomPadding()
}

// HandleReflect handles the 'reflect' command for mindful reflection
func HandleReflect(args []string) {
	breathing.PrintWithPadding("🕯️  Evening Reflection")
	breathing.AddSectionSpacing()

	breathing.PrintWithPadding("   Close your eyes for a moment...")
	breathing.PrintWithPadding("   Take three deep breaths...")
	breathing.AddSectionSpacing()

	breathing.PrintWithPadding("   📝 Gentle reflection:")
	breathing.PrintWithPadding("      • What thoughts kept pulling you away today?")
	breathing.PrintWithPadding("      • Were there moments when you were truly present?")
	breathing.PrintWithPadding("      • What patterns do you notice in your mind?")
	breathing.AddSectionSpacing()

	breathing.PrintWithPadding("   These are just thoughts. They come and go like clouds.")
	breathing.PrintWithPadding("   The noticing itself is the practice. 🙏")
	breathing.AddBottomPadding()
}

// HandleVersion handles version display
func HandleVersion(programName string) {
	fmt.Printf("%s version %s\n", programName, version.Version)
}

// HandleUnknownCommand handles unknown commands
func HandleUnknownCommand(command string, programName string) {
	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
	fmt.Fprintf(os.Stderr, "Run '%s help' for available commands.\n", programName)
	os.Exit(1)
}
