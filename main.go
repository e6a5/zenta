package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/e6a5/zenta/internal/models"
	"github.com/e6a5/zenta/internal/quotes"
	"github.com/e6a5/zenta/internal/stats"
	"github.com/e6a5/zenta/internal/storage"
	"github.com/e6a5/zenta/internal/version"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "now":
		handleNow()
	case "log":
		handleLog(os.Args[2:])
	case "stats":
		handleStats(os.Args[2:])
	case "help":
		showHelp()
	case "version", "--version", "-v":
		handleVersion()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		fmt.Fprintf(os.Stderr, "Run 'zenta help' for available commands.\n")
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("zenta - mindfulness for terminal users")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  zenta now                    Show a mindfulness quote")
	fmt.Println("  zenta log [-t type] <reason> Log a moment of distraction, reflection, or insight")
	fmt.Println("  zenta stats [period]         View analytics from logs")
	fmt.Println("  zenta help                   Show this help message")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  zenta now")
	fmt.Println("  zenta log \"Scrolled social media instead of coding\"")
	fmt.Println("  zenta log -t reflection \"Noticed I was feeling anxious about the deadline\"")
	fmt.Println("  zenta log -t insight \"Deep work is easier in the morning\"")
	fmt.Println("  zenta stats")
	fmt.Println("  zenta stats week")
	fmt.Println()
	fmt.Println("LOG TYPES:")
	fmt.Println("  distraction (default), reflection, insight")
	fmt.Println("  Short forms: d, r, i")
	fmt.Println()
	fmt.Println("PERIODS:")
	fmt.Println("  today, week, month, all")
	fmt.Println()
	fmt.Println("Learn more: https://github.com/e6a5/zenta")
}

func handleNow() {
	fmt.Println("🧘 Let's take a moment to breathe together...")
	fmt.Println("   Follow the rhythm: Inhale → Hold → Exhale → Hold")
	fmt.Println()

	// Perform 3 cycles of 4-4-4-4 breathing
	for cycle := 1; cycle <= 3; cycle++ {
		fmt.Printf("   Cycle %d/3:\n", cycle)

		// Inhale phase (4 seconds)
		fmt.Print("   🌬️  Inhale ")
		for i := 0; i < 4; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print("●")
		}
		fmt.Println()

		// Hold phase (4 seconds)
		fmt.Print("   ⏸️  Hold   ")
		for i := 0; i < 4; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print("○")
		}
		fmt.Println()

		// Exhale phase (4 seconds)
		fmt.Print("   💨 Exhale ")
		for i := 0; i < 4; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print("●")
		}
		fmt.Println()

		// Hold phase (4 seconds)
		fmt.Print("   ⏸️  Hold   ")
		for i := 0; i < 4; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print("○")
		}
		fmt.Println()

		if cycle < 3 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println("✨ Beautiful. Now, here's a moment of wisdom:")
	fmt.Println()

	quoteService := quotes.New()
	quote := quoteService.GetRandomQuote()
	fmt.Println(quote)
}

func handleLog(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: Please provide a reason for logging\n")
		fmt.Fprintf(os.Stderr, "Usage: zenta log [-t type] <reason>\n")
		fmt.Fprintf(os.Stderr, "Types: distraction (default), reflection, insight\n")
		os.Exit(1)
	}

	// Parse arguments for type flag
	logType := models.LogTypeDistraction // default
	var reason string

	if len(args) >= 3 && args[0] == "-t" {
		// Format: zenta log -t <type> <reason>
		typeStr := args[1]
		switch typeStr {
		case "distraction", "d":
			logType = models.LogTypeDistraction
		case "reflection", "r":
			logType = models.LogTypeReflection
		case "insight", "i":
			logType = models.LogTypeInsight
		default:
			fmt.Fprintf(os.Stderr, "Error: Invalid log type '%s'\n", typeStr)
			fmt.Fprintf(os.Stderr, "Valid types: distraction, reflection, insight\n")
			os.Exit(1)
		}
		reason = strings.Join(args[2:], " ")
	} else {
		// Format: zenta log <reason> (default to distraction)
		reason = strings.Join(args, " ")
	}

	if strings.TrimSpace(reason) == "" {
		fmt.Fprintf(os.Stderr, "Error: Please provide a reason for logging\n")
		os.Exit(1)
	}

	// Initialize storage
	store, err := storage.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing storage: %v\n", err)
		os.Exit(1)
	}

	// Create new log entry with specified type
	entry := models.NewLogEntry(logType, reason)

	// Save the log entry
	if err := store.AddLog(entry); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving log entry: %v\n", err)
		os.Exit(1)
	}

	// Show appropriate message based on log type
	var emoji, message string
	switch logType {
	case models.LogTypeDistraction:
		emoji = "🔴"
		message = "Take a moment to breathe and return to the present."
	case models.LogTypeReflection:
		emoji = "🤔"
		message = "Reflection noted. What did you learn from this moment?"
	case models.LogTypeInsight:
		emoji = "💡"
		message = "Insight captured. Wisdom grows through awareness."
	case models.LogTypeSessionStart:
		emoji = "⏰"
		message = "Session started."
	case models.LogTypeSessionEnd:
		emoji = "✅"
		message = "Session completed."
	default:
		emoji = "✓"
		message = "Entry logged."
	}

	fmt.Printf("%s Logged (%s): %s\n", emoji, logType, reason)
	fmt.Printf("  %s\n", message)
}

func handleStats(args []string) {
	// Determine period
	period := "week" // default period
	if len(args) > 0 {
		period = args[0]
	}

	// Validate period
	validPeriods := map[string]bool{
		"today": true,
		"week":  true,
		"month": true,
		"all":   true,
	}

	if !validPeriods[period] {
		fmt.Fprintf(os.Stderr, "Error: Invalid period '%s'\n", period)
		fmt.Fprintf(os.Stderr, "Valid periods: today, week, month, all\n")
		os.Exit(1)
	}

	// Initialize storage and load logs
	store, err := storage.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing storage: %v\n", err)
		os.Exit(1)
	}

	logs, err := store.LoadLogs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading logs: %v\n", err)
		os.Exit(1)
	}

	// Generate and display statistics
	analyzer := stats.New(logs)
	result := analyzer.GenerateStats(period)
	output := analyzer.FormatStats(result, period)

	fmt.Print(output)
}

func handleVersion() {
	versionInfo := version.Get()
	fmt.Println(versionInfo.String())
}
