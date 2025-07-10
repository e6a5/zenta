// Package cli provides command-line interface handling for zenta.
// It manages command routing, help display, and user interaction.
package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/e6a5/zenta/internal/breathing"
	"github.com/e6a5/zenta/internal/quotes"
	"github.com/e6a5/zenta/internal/reflection"
	"github.com/e6a5/zenta/internal/version"
)

const MinArgs = 2

// ShowHelp displays the main help message
func ShowHelp(programName string) {
	fmt.Printf("%s - mindfulness for terminal users\n", programName)
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Printf("  %s now [options]         Take a mindful breathing moment\n", programName)
	fmt.Printf("  %s anchor                Guided breathing anchor\n", programName)
	fmt.Printf("  %s reflect               End-of-day reflection on thought patterns\n", programName)
	fmt.Printf("  %s help                  Show this help message\n", programName)
	fmt.Println()
	fmt.Println("NOW OPTIONS:")
	fmt.Println("  --quick, -q                 Quick 1-cycle session")
	fmt.Println("  --extended, -e              Extended 5-cycle session")
	fmt.Println("  --silent, -s                Breathing only, skip the quote")
	fmt.Println("  --simple                    Simple line animation (for terminal compatibility)")
	fmt.Println("  --complex                   Force complex animation (default except on Apple Terminal)")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Printf("  %s now                   Standard 3-cycle breathing session\n", programName)
	fmt.Printf("  %s now --quick           Quick 1-cycle breathing break\n", programName)
	fmt.Printf("  %s now --extended        Extended 5-cycle session\n", programName)
	fmt.Printf("  %s now --silent          Breathing without quote\n", programName)
	fmt.Printf("  %s now --simple          Simple animation (terminal compatibility)\n", programName)
	fmt.Printf("  %s anchor                Anchor your breath to the present moment\n", programName)
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

	defer session.HideCursor()()
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

// HandleAnchor handles the 'anchor' command for the interactive pacer.
func HandleAnchor(args []string) {
	session := breathing.NewSession()
	session.StartAnchor()

	// Show a quote after the session, unless it was silent.
	// This check is a placeholder for future flags, e.g., --breathe-silent
	if session.ShouldShowQuote() {
		quoteService := quotes.New()
		quote := quoteService.GetRandomQuote()
		quotes.DisplayBeautifully(quote)
	}
}

// HandleReflect handles the 'reflect' command for mindful reflection
func HandleReflect(args []string) {
	prompts := reflection.GetDefaultPrompts()

	// Begin the session
	breathing.PrintWithPadding(prompts.Title)
	time.Sleep(1 * time.Second)
	breathing.AddSectionSpacing()

	// Guide through initial instructions with pauses
	for i, line := range prompts.Instructions {
		breathing.PrintWithPadding(line)
		// Give more time for the last instruction (taking breaths)
		if i == len(prompts.Instructions)-1 {
			time.Sleep(5 * time.Second)
		} else {
			time.Sleep(3 * time.Second)
		}
	}
	breathing.AddSectionSpacing()

	// Introduce the reflection prompts
	breathing.PrintWithPadding(prompts.PromptTitle)
	time.Sleep(2 * time.Second)

	// Display each prompt with a long pause for contemplation
	for _, line := range prompts.Prompts {
		breathing.PrintWithPadding(line)
		time.Sleep(8 * time.Second)
	}
	breathing.AddSectionSpacing()

	// Display the closing thoughts with pauses
	for _, line := range prompts.Closing {
		breathing.PrintWithPadding(line)
		time.Sleep(3 * time.Second)
	}
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
