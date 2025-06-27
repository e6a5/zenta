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
	ProgressBarWidth   = 8

	// Layout padding constants
	LeftPadding    = 4 // Left margin for all content
	RightPadding   = 4 // Right margin buffer
	BottomPadding  = 2 // Extra spacing at bottom
	SectionSpacing = 1 // Space between sections
)

// Breathing session configuration
type BreathingSession struct {
	Cycles    int
	ShowQuote bool
	InhaleDur int
	HoldDur   int
	ExhaleDur int
	RestDur   time.Duration
}

func main() {
	if len(os.Args) < MinArgs {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "now":
		handleNow(os.Args[2:])
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
	fmt.Println("  zenta now [options]          Take a mindful breathing moment")
	fmt.Println("  zenta log [-t type] <reason> Log a moment of distraction, reflection, or insight")
	fmt.Println("  zenta stats [period]         View analytics from logs")
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
	fmt.Println("  zenta log \"Scrolled social media instead of coding\"")
	fmt.Println("  zenta log -t reflection \"Noticed I was feeling anxious\"")
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

func handleNow(args []string) {
	session := parseNowArgs(args)

	// Add top spacing
	fmt.Println()

	// Section spacing

	// Go straight into breathing - no interruptions
	printWithPadding("   Let's breathe 🌸")

	addSectionSpacing()

	// Reserve dedicated space for breathing visualization (12 lines)
	fmt.Println() // Guidance text line
	for i := 0; i < 10; i++ {
		fmt.Println() // Circle area
	}
	fmt.Println() // Bottom buffer

	// Move cursor back to start of breathing area
	fmt.Print("\033[12A")

	// One lung, multiple breaths
	drawContinuousBreathingSession(session)

	// Move cursor to end of breathing area
	fmt.Print("\033[12B")

	// Completion section with minimal design
	addSectionSpacing()

	if session.ShowQuote {
		quoteService := quotes.New()
		quote := quoteService.GetRandomQuote()
		displayQuoteBeautifully(quote)
	} else {
		printWithPadding("   Carry this calm with you throughout your day 🙏")
	}

	// Add bottom padding
	addBottomPadding()
}

// parseNowArgs parses command line arguments for the now command
func parseNowArgs(args []string) *BreathingSession {
	session := &BreathingSession{
		Cycles:    3,
		ShowQuote: true,
		InhaleDur: 4,
		HoldDur:   4,
		ExhaleDur: 4,
		RestDur:   RestDuration,
	}

	for _, arg := range args {
		switch arg {
		case "--quick", "-q":
			session.Cycles = 1
		case "--extended", "-e":
			session.Cycles = 5
		case "--silent", "-s":
			session.ShowQuote = false
		}
	}

	return session
}

// checkForExit checks if user pressed 'q' to exit (simplified for now)
func checkForExit() bool {
	// For now, we'll implement a simple version
	// In a full implementation, this would use non-blocking input
	return false
}

// drawContinuousBreathingSession draws one lung breathing continuously through multiple cycles
func drawContinuousBreathingSession(session *BreathingSession) {
	// Breathing phases that feel natural
	phases := []struct {
		name        string
		emoji       string
		duration    int
		instruction string
		breathType  string
	}{
		{"inhale", "🌬️", session.InhaleDur, "Breathe in gently, let your body expand...", "expand"},
		{"hold", "✨", session.HoldDur, "Hold softly, feel the fullness...", "full"},
		{"exhale", "🌸", session.ExhaleDur, "Release slowly, let everything go...", "contract"},
		{"rest", "🕯️", session.HoldDur, "Rest in the emptiness, be present...", "empty"},
	}

	// One lung, breathing continuously through all cycles
	for cycle := 1; cycle <= session.Cycles; cycle++ {
		for _, phase := range phases {
			if checkForExit() {
				return
			}

			// Show gentle guidance for each phase
			showBreathingGuidance(phase.emoji, phase.name, phase.instruction)

			// Animate the same breathing circle for this phase
			animateBreathingCircle(phase.breathType, phase.duration)
		}

		// Brief pause between breathing cycles (not a separate "rest" phase)
		if cycle < session.Cycles {
			showBreathingGuidance("💫", "rest", "Feel the rhythm... continuing...")
			time.Sleep(session.RestDur)
		}
	}

	clearBreathingDisplay()
}

// showBreathingGuidance shows gentle, non-technical breathing guidance
func showBreathingGuidance(emoji, phase, instruction string) {
	fmt.Print("\033[s") // Save cursor position
	// Move to guidance line (current line)
	fmt.Print("\r") // Go to beginning of line
	fmt.Printf("%s   %s %s", strings.Repeat(" ", LeftPadding), emoji, instruction)
	fmt.Print(strings.Repeat(" ", 20)) // Clear rest of line
	fmt.Print("\033[u")                // Restore cursor position
}

// animateBreathingCircle creates an organic breathing circle that expands/contracts
func animateBreathingCircle(breathType string, duration int) {
	// Position circle in the center of the reserved area (relative to current cursor)
	centerRowOffset := 5          // 5 lines down from guidance text
	centerCol := LeftPadding + 25 // Centered position

	for second := 1; second <= duration; second++ {
		if checkForExit() {
			return
		}

		// Calculate circle size based on breath type and progress
		var circleSize int
		var circleChar string

		switch breathType {
		case "expand": // Inhale - circle grows
			progress := float64(second) / float64(duration)
			circleSize = int(1 + progress*3) // Size 1-4
			circleChar = "○"
		case "full": // Hold full - circle stays large with gentle pulse
			circleSize = 4
			if second%2 == 0 {
				circleChar = "●"
			} else {
				circleChar = "○"
			}
		case "contract": // Exhale - circle shrinks
			progress := float64(duration-second+1) / float64(duration)
			circleSize = int(1 + progress*3) // Size 4-1
			circleChar = "○"
		case "empty": // Hold empty - small circle with gentle pulse
			circleSize = 1
			if second%2 == 0 {
				circleChar = "·"
			} else {
				circleChar = "○"
			}
		}

		// Clear previous circle using relative positioning
		clearCircleAreaRelative(centerRowOffset, centerCol)

		// Draw the breathing circle using relative positioning
		drawBreathingCircleRelative(centerRowOffset, centerCol, circleSize, circleChar)

		time.Sleep(1 * time.Second)
	}
}

// clearCircleAreaRelative clears the area where the breathing circle is drawn using relative positioning
func clearCircleAreaRelative(rowOffset, centerCol int) {
	fmt.Print("\033[s") // Save cursor position
	for row := rowOffset - 4; row <= rowOffset+4; row++ {
		fmt.Printf("\033[%dB", row)        // Move to row
		fmt.Print("\r")                    // Go to beginning of line
		fmt.Print(strings.Repeat(" ", 80)) // Clear entire line
		fmt.Print("\033[u")                // Restore position
		fmt.Print("\033[s")                // Save again for next iteration
	}
	fmt.Print("\033[u") // Final restore
}

// drawBreathingCircleRelative draws a circular breathing pattern using relative positioning
func drawBreathingCircleRelative(rowOffset, centerCol, size int, char string) {
	if size <= 1 {
		// Small circle - just center point
		drawAtPositionRelative(rowOffset, centerCol, char)
		return
	}

	// Draw concentric circles for larger sizes
	for radius := 1; radius <= size; radius++ {
		// Calculate positions for circle points
		points := []struct{ row, col int }{
			{rowOffset - radius, centerCol},   // top
			{rowOffset + radius, centerCol},   // bottom
			{rowOffset, centerCol - radius*2}, // left (wider for terminal)
			{rowOffset, centerCol + radius*2}, // right
		}

		// Add diagonal points for larger circles
		if radius > 1 {
			diag := int(float64(radius) * 0.7) // Approximate diagonal distance
			points = append(points, []struct{ row, col int }{
				{rowOffset - diag, centerCol - diag}, // top-left
				{rowOffset - diag, centerCol + diag}, // top-right
				{rowOffset + diag, centerCol - diag}, // bottom-left
				{rowOffset + diag, centerCol + diag}, // bottom-right
			}...)
		}

		// Draw all points for this radius
		circleChar := char
		if radius < size {
			circleChar = "·" // Fainter for inner circles
		}

		for _, point := range points {
			drawAtPositionRelative(point.row, point.col, circleChar)
		}
	}
}

// drawAtPositionRelative draws a character at a position relative to current cursor
func drawAtPositionRelative(rowOffset, col int, char string) {
	fmt.Print("\033[s") // Save cursor position

	// Move relative to current position
	if rowOffset > 0 {
		fmt.Printf("\033[%dB", rowOffset) // Move down
	} else if rowOffset < 0 {
		fmt.Printf("\033[%dA", -rowOffset) // Move up
	}

	// Move to column position
	fmt.Print("\r") // Go to beginning of line
	if col > 0 {
		fmt.Printf("\033[%dC", col) // Move right
	}

	fmt.Print(char)
	fmt.Print("\033[u") // Restore cursor position
}

// clearBreathingDisplay clears the breathing visualization area
func clearBreathingDisplay() {
	// Clear the guidance line
	fmt.Print("\033[s")                // Save cursor position
	fmt.Print("\r")                    // Go to beginning of line
	fmt.Print(strings.Repeat(" ", 80)) // Clear line
	fmt.Print("\033[u")                // Restore cursor position
}

// Padding and spacing helper functions
func printWithPadding(text string) {
	fmt.Printf("%s%s\n", strings.Repeat(" ", LeftPadding), text)
}

func addSectionSpacing() {
	for i := 0; i < SectionSpacing; i++ {
		fmt.Println()
	}
}

func addBottomPadding() {
	for i := 0; i < BottomPadding; i++ {
		fmt.Println()
	}
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
	renderQuoteBoxWithTyping(lines, emoji, boxWidth)
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

func renderQuoteBoxWithTyping(lines []string, emoji string, boxWidth int) {
	fmt.Println()

	// Simple quote display without box borders
	for i, line := range lines {
		renderQuoteLineSimpleWithTyping(line, emoji, i == 0)
	}

	fmt.Println()
}

func renderQuoteLineSimpleWithTyping(line, emoji string, isFirstLine bool) {
	// Start the line with proper padding
	fmt.Printf("%s   ", strings.Repeat(" ", LeftPadding))

	// Type out the emoji first (if any)
	if isFirstLine && emoji != "" {
		fmt.Print(emoji)
		time.Sleep(200 * time.Millisecond)
		fmt.Print(" ")
		time.Sleep(100 * time.Millisecond)
	}

	// Type out each word with a pause
	words := strings.Fields(line)
	for i, word := range words {
		if i > 0 {
			fmt.Print(" ")
			time.Sleep(50 * time.Millisecond)
		}

		// Type each character in the word
		for _, char := range word {
			fmt.Print(string(char))
			time.Sleep(80 * time.Millisecond)
		}

		// Brief pause between words
		time.Sleep(150 * time.Millisecond)
	}

	fmt.Println()
}

func isEmoji(r rune) bool {
	// Simple emoji detection for common ranges
	return (r >= 0x1F300 && r <= 0x1F6FF) || // Miscellaneous Symbols and Pictographs
		(r >= 0x1F900 && r <= 0x1F9FF) || // Supplemental Symbols and Pictographs
		(r >= 0x2600 && r <= 0x26FF) || // Miscellaneous Symbols
		(r >= 0x2700 && r <= 0x27BF) // Dingbats
}
