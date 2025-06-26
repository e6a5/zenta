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

// Constants for various display and timing values
const (
	MinArgs            = 2
	MaxQuoteWidth      = 50
	QuoteBoxMinWidth   = 50
	QuotePadding       = 8
	QuoteBorderPadding = 4
	CenterDivisor      = 2
	CycleSpacing       = 20
	MaxCycles          = 3
	RestDuration       = 2 * time.Second
	ClearLineWidth     = 60
	StatusLineWidth    = 80
	BoxWidth           = 12
	BoxHeight          = 4
	DotSpacing         = 3
)

func main() {
	if len(os.Args) < MinArgs {
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
	fmt.Println("🧘 Box breathing - let's breathe together...")
	fmt.Println()
	fmt.Println("   Get ready... we'll do 3 breathing cycles")
	fmt.Println()

	// Preparation countdown
	for i := 3; i >= 1; i-- {
		fmt.Printf("   Starting in %d...\r", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("   Let's breathe! 🌸    \n\n")

	// Labels for the three cycles
	fmt.Println("   Cycle 1        Cycle 2        Cycle 3")
	fmt.Println()

	// Reserve space for 3 horizontal boxes
	for i := 0; i < 6; i++ {
		fmt.Println()
	}

	// Move cursor back up to start drawing
	fmt.Print("\033[6A")

	// Draw 3 cycles horizontally with guidance
	for cycle := 1; cycle <= MaxCycles; cycle++ {
		drawBoxWithGuidance(cycle)
		if cycle < MaxCycles {
			showRestPause()
		}
	}

	// Move cursor past all boxes
	fmt.Print("\033[6B")

	// Simple, clean celebration
	fmt.Println()
	fmt.Println("   ✨ Perfect! Here's wisdom to carry this calm with you:")
	fmt.Println()

	quoteService := quotes.New()
	quote := quoteService.GetRandomQuote()
	displayQuoteBeautifully(quote)
}

func drawBoxWithGuidance(cycleNum int) {
	// Calculate horizontal offset for each box
	xOffset := (cycleNum - 1) * CycleSpacing

	// Highlight current cycle
	highlightCurrentCycle(cycleNum)

	// Top edge - Inhale with guidance
	showPhaseInstruction("🌬️  Breathe IN slowly and deeply...")
	for i := 1; i <= BoxHeight; i++ {
		drawDotAtPosition(0, xOffset+i*DotSpacing, "🔵")
		showCountdown(BoxHeight - i + 1)
		time.Sleep(1 * time.Second)
	}

	// Right edge - Hold with guidance
	showPhaseInstruction("⏸️  Hold your breath gently...")
	for i := 1; i <= BoxHeight; i++ {
		drawDotAtPosition(i, xOffset+BoxWidth, "🔴")
		showCountdown(BoxHeight - i + 1)
		time.Sleep(1 * time.Second)
	}

	// Bottom edge - Exhale with guidance
	showPhaseInstruction("💨 Breathe OUT slowly, release all tension...")
	for i := 1; i <= BoxHeight; i++ {
		drawDotAtPosition(BoxHeight, xOffset+BoxWidth-i*DotSpacing, "🟡")
		showCountdown(BoxHeight - i + 1)
		time.Sleep(1 * time.Second)
	}

	// Left edge - Hold with guidance
	showPhaseInstruction("⏸️  Hold empty, stay present...")
	for i := 1; i <= BoxHeight; i++ {
		drawDotAtPosition(BoxHeight-i, xOffset, "🔴")
		showCountdown(BoxHeight - i + 1)
		time.Sleep(1 * time.Second)
	}

	// Clear the instruction line
	clearInstructionLine()
}

func highlightCurrentCycle(currentCycle int) {
	// Save cursor position
	fmt.Print("\033[s")

	// Move to cycle labels line
	fmt.Print("\033[7A")
	fmt.Print("\r")

	// Show all cycles with current one highlighted
	for i := 1; i <= MaxCycles; i++ {
		switch {
		case i == currentCycle:
			fmt.Print("   ▶ Cycle " + fmt.Sprintf("%d", i) + " ◀    ")
		case i < currentCycle:
			fmt.Print("   ✓ Cycle " + fmt.Sprintf("%d", i) + "     ")
		default:
			fmt.Print("   Cycle " + fmt.Sprintf("%d", i) + "       ")
		}
	}

	// Restore cursor position
	fmt.Print("\033[u")
}

func showPhaseInstruction(instruction string) {
	// Save cursor position
	fmt.Print("\033[s")

	// Move to instruction line (below the boxes)
	fmt.Print("\033[7B")
	fmt.Print("\r")
	fmt.Print(instruction)
	fmt.Print(strings.Repeat(" ", ClearLineWidth)) // Clear rest of line

	// Restore cursor position
	fmt.Print("\033[u")
}

func showCountdown(count int) {
	if count > 1 {
		// Save cursor position
		fmt.Print("\033[s")

		// Move to countdown position (next to instruction)
		fmt.Print("\033[7B")
		fmt.Print("\033[50C") // Move right to countdown area
		fmt.Printf("(%d)", count-1)

		// Restore cursor position
		fmt.Print("\033[u")
	}
}

func clearInstructionLine() {
	// Save cursor position
	fmt.Print("\033[s")

	// Move to instruction line and clear it
	fmt.Print("\033[7B")
	fmt.Print("\r")
	fmt.Print(strings.Repeat(" ", StatusLineWidth))

	// Restore cursor position
	fmt.Print("\033[u")
}

func showRestPause() {
	showPhaseInstruction("💫 Rest between cycles... feel the calm...")
	time.Sleep(RestDuration)
	clearInstructionLine()
}

func drawDotAtPosition(row, col int, dot string) {
	// Save current cursor position
	fmt.Print("\033[s")

	// Move to specific position relative to current cursor
	if row > 0 {
		fmt.Printf("\033[%dB", row) // Move down
	}
	if col > 0 {
		fmt.Printf("\033[%dC", col) // Move right
	}

	// Draw the dot
	fmt.Print(dot)

	// Restore cursor position
	fmt.Print("\033[u")
}

func handleLog(args []string) {
	if len(args) == 0 {
		printLogUsage()
		os.Exit(1)
	}

	logType, reason := parseLogArgs(args)

	if strings.TrimSpace(reason) == "" {
		fmt.Fprintf(os.Stderr, "Error: Please provide a reason for logging\n")
		os.Exit(1)
	}

	saveLogEntry(logType, reason)
	printLogConfirmation(logType, reason)
}

func printLogUsage() {
	fmt.Fprintf(os.Stderr, "Error: Please provide a reason for logging\n")
	fmt.Fprintf(os.Stderr, "Usage: zenta log [-t type] <reason>\n")
	fmt.Fprintf(os.Stderr, "Types: distraction (default), reflection, insight\n")
}

func parseLogArgs(args []string) (models.LogType, string) {
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

	return logType, reason
}

func saveLogEntry(logType models.LogType, reason string) {
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
}

func printLogConfirmation(logType models.LogType, reason string) {
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

func displayQuoteBeautifully(quote string) {
	emoji, quoteText := parseQuoteEmoji(quote)
	lines := wrapQuoteText(quoteText)
	boxWidth := calculateBoxWidth(lines)
	renderQuoteBox(lines, emoji, boxWidth)
}

func parseQuoteEmoji(quote string) (string, string) {
	runes := []rune(quote)
	if len(runes) > 0 && isEmoji(runes[0]) {
		emoji := string(runes[0])
		quoteText := strings.TrimSpace(string(runes[1:]))
		return emoji, quoteText
	}
	return "", quote
}

func wrapQuoteText(quoteText string) []string {
	words := strings.Fields(quoteText)
	lines := []string{}
	currentLine := ""

	for _, word := range words {
		if len(currentLine+" "+word) <= MaxQuoteWidth {
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		} else {
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
			currentLine = word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}

func calculateBoxWidth(lines []string) int {
	longestLine := 0
	for _, line := range lines {
		if len(line) > longestLine {
			longestLine = len(line)
		}
	}

	boxWidth := longestLine + QuotePadding
	if boxWidth < QuoteBoxMinWidth {
		boxWidth = QuoteBoxMinWidth
	}
	return boxWidth
}

func renderQuoteBox(lines []string, emoji string, boxWidth int) {
	fmt.Println()

	// Top border
	renderBorderLine("╭", "─", "╮", boxWidth)

	// Empty line
	renderEmptyLine(boxWidth)

	// Quote lines with emoji
	for i, line := range lines {
		renderQuoteLine(line, emoji, i == 0, boxWidth)
	}

	// Empty line
	renderEmptyLine(boxWidth)

	// Bottom border
	renderBorderLine("╰", "─", "╯", boxWidth)
	fmt.Println()
}

func renderBorderLine(left, middle, right string, width int) {
	fmt.Print("   ")
	fmt.Print(left + strings.Repeat(middle, width) + right)
	fmt.Println()
}

func renderEmptyLine(boxWidth int) {
	fmt.Print("   ")
	fmt.Print("│" + strings.Repeat(" ", boxWidth) + "│")
	fmt.Println()
}

func renderQuoteLine(line, emoji string, isFirstLine bool, boxWidth int) {
	fmt.Print("   ")
	fmt.Print("│  ")

	// Add emoji to first line
	displayLine := line
	if isFirstLine && emoji != "" {
		displayLine = emoji + " " + line
	}

	// Center the text safely
	availableSpace := boxWidth - QuoteBorderPadding
	padding := (availableSpace - len(displayLine)) / CenterDivisor
	if padding < 0 {
		padding = 0
	}

	fmt.Print(strings.Repeat(" ", padding))
	fmt.Print(displayLine)

	// Right padding
	rightPadding := availableSpace - len(displayLine) - padding
	if rightPadding < 0 {
		rightPadding = 0
	}
	fmt.Print(strings.Repeat(" ", rightPadding))
	fmt.Print("  │")
	fmt.Println()
}

func isEmoji(r rune) bool {
	// Simple emoji detection for common ranges
	return (r >= 0x1F300 && r <= 0x1F6FF) || // Miscellaneous Symbols and Pictographs
		(r >= 0x1F900 && r <= 0x1F9FF) || // Supplemental Symbols and Pictographs
		(r >= 0x2600 && r <= 0x26FF) || // Miscellaneous Symbols
		(r >= 0x2700 && r <= 0x27BF) // Dingbats
}
