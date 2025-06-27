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
func ShowHelp() {
	fmt.Println("zenta - mindfulness for terminal users")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  zenta now [options]          Take a mindful breathing moment")
	fmt.Println("  zenta reflect                End-of-day reflection on thought patterns")
	fmt.Println("  zenta help                   Show this help message")
	fmt.Println()
	fmt.Println("NOW OPTIONS:")
	fmt.Println("  --quick, -q                  Quick 1-cycle session (1 min)")
	fmt.Println("  --extended, -e               Extended 5-cycle session (5 min)")
	fmt.Println("  --silent, -s                 Breathing only, skip the quote")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  zenta now                    Standard 3-cycle breathing session")
	fmt.Println("  zenta now --quick            Quick 1-minute breathing break")
	fmt.Println("  zenta now --extended         Extended 5-minute session")
	fmt.Println("  zenta now --silent           Breathing without quote")
	fmt.Println("  zenta reflect                Gentle end-of-day reflection")
	fmt.Println()
	fmt.Println("MINDFUL ALIASES:")
	fmt.Println("  alias breath='zenta now --silent --quick'")
	fmt.Println("  alias breathe='zenta now --silent'")
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
func HandleVersion() {
	fmt.Printf("zenta version %s\n", version.Version)
}

// HandleUnknownCommand handles unknown commands
func HandleUnknownCommand(command string) {
	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
	fmt.Fprintf(os.Stderr, "Run 'zenta help' for available commands.\n")
	os.Exit(1)
}
