package main

import (
	"bufio"
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
	LeftPadding        = 4   // Left margin for all content
	RightPadding       = 4   // Right margin buffer
	BottomPadding      = 2   // Extra spacing at bottom
	SectionSpacing     = 1   // Space between sections
)

// Breathing session configuration
type BreathingSession struct {
	Cycles       int
	ShowQuote    bool
	InhaleDur    int
	HoldDur      int
	ExhaleDur    int
	RestDur      time.Duration
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
	
	// Welcome section with consistent left padding
	printWithPadding("🧘 Welcome to your mindful moment")
	if session.Cycles == 1 {
		printWithPadding("   Find a comfortable position... Quick session (1 cycle)")
	} else {
		printWithPadding(fmt.Sprintf("   Find a comfortable position... %d breathing cycles", session.Cycles))
	}
	printWithPadding("   Press 'q' anytime to exit gracefully")
	
	// Section spacing
	addSectionSpacing()
	
	// Enhanced preparation
	printWithPadding("   When ready, press [ENTER] to begin...")
	waitForUserReady()
	
	// Gentle countdown with padding
	for i := 3; i >= 1; i-- {
		if checkForExit() {
			return
		}
		fmt.Printf("%s   Starting in %d...\r", strings.Repeat(" ", LeftPadding), i)
		time.Sleep(1 * time.Second)
	}
	printWithPadding("   Let's breathe! 🌸")
	
	addSectionSpacing()

	// Dynamic cycle headers based on session cycles
	showCycleHeaders(session.Cycles)
	
	// Reserve space for breathing visualization
	for i := 0; i < 6; i++ {
		fmt.Println()
	}
	fmt.Print("\033[6A")

	// Enhanced breathing cycles
	for cycle := 1; cycle <= session.Cycles; cycle++ {
		if checkForExit() {
			return
		}
		drawEnhancedBreathingCycle(cycle, session)
		if cycle < session.Cycles {
			showRestPause()
		}
	}

	fmt.Print("\033[6B")
	
	// Completion section with spacing
	addSectionSpacing()
	printWithPadding("   ✨ Perfect! You've completed your mindful moment")
	
	if session.ShowQuote {
		addSectionSpacing()
		printWithPadding("   Here's wisdom to carry this calm with you:")
		fmt.Println()
		quoteService := quotes.New()
		quote := quoteService.GetRandomQuote()
		displayQuoteBeautifully(quote)
	} else {
		addSectionSpacing()
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

// waitForUserReady waits for user to press ENTER
func waitForUserReady() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// checkForExit checks if user pressed 'q' to exit (simplified for now)
func checkForExit() bool {
	// For now, we'll implement a simple version
	// In a full implementation, this would use non-blocking input
	return false
}

// showCycleHeaders displays dynamic cycle headers based on session configuration
func showCycleHeaders(cycles int) {
	headerText := ""
	for i := 1; i <= cycles && i <= 5; i++ { // Limit display to 5 cycles
		if i == 1 {
			headerText += fmt.Sprintf("Cycle %d", i)
		} else {
			headerText += fmt.Sprintf("        Cycle %d", i)
		}
	}
	printWithPadding("   " + headerText)
	fmt.Println()
}

// drawEnhancedBreathingCycle draws an improved breathing cycle with progress bars
func drawEnhancedBreathingCycle(cycleNum int, session *BreathingSession) {
	// Calculate horizontal offset with proper spacing for multiple cycles
	xOffset := (cycleNum - 1) * CycleSpacing
	
	// Highlight current cycle
	highlightCurrentCycle(cycleNum)

	// Enhanced breathing phases with progress visualization
	phases := []struct {
		name        string
		emoji       string
		duration    int
		instruction string
		dot         string
	}{
		{"inhale", "🌬️", session.InhaleDur, "Breathe IN slowly and deeply...", "🔵"},
		{"hold", "⏸️", session.HoldDur, "Hold your breath gently...", "🔴"},
		{"exhale", "💨", session.ExhaleDur, "Breathe OUT slowly, release all tension...", "🟡"},
		{"hold", "⏸️", session.HoldDur, "Hold empty, stay present...", "🔴"},
	}

	for phaseIdx, phase := range phases {
		showEnhancedPhaseInstruction(phase.emoji, phase.instruction)
		
		for i := 1; i <= phase.duration; i++ {
			if checkForExit() {
				return
			}
			
			// Draw progress bar instead of just dots
			showBreathingProgress(phase.emoji, phase.name, i, phase.duration)
			
			// Still draw the box dots for visual continuity
			var row, col int
			switch phaseIdx {
			case 0: // inhale - top edge
				row, col = 0, xOffset+i*DotSpacing
			case 1: // hold - right edge  
				row, col = i, xOffset+BoxWidth
			case 2: // exhale - bottom edge
				row, col = BoxHeight, xOffset+BoxWidth-i*DotSpacing
			case 3: // hold - left edge
				row, col = BoxHeight-i, xOffset
			}
			
			drawDotAtPosition(row, col, phase.dot)
			time.Sleep(1 * time.Second)
		}
	}

	clearInstructionLine()
}

// showEnhancedPhaseInstruction shows breathing phase with better formatting
func showEnhancedPhaseInstruction(emoji, instruction string) {
	fmt.Print("\033[s") // Save cursor position
	fmt.Print("\033[7B") // Move to instruction line
	fmt.Print("\r")
	fmt.Printf("%s   %s  %s", strings.Repeat(" ", LeftPadding), emoji, instruction)
	fmt.Print(strings.Repeat(" ", ClearLineWidth))
	fmt.Print("\033[u") // Restore cursor position
}

// showBreathingProgress shows a progress bar for the current breathing phase
func showBreathingProgress(emoji, phase string, current, total int) {
	fmt.Print("\033[s") // Save cursor position
	fmt.Print("\033[8B") // Move below instruction line
	fmt.Print("\r")
	
	// Create progress bar
	filled := strings.Repeat("█", current)
	empty := strings.Repeat("░", total-current)
	progress := fmt.Sprintf("[%s%s]", filled, empty)
	
	// Show phase name and progress with left padding
	fmt.Printf("%s   %s %-8s %s %ds", strings.Repeat(" ", LeftPadding), emoji, strings.Title(phase), progress, total-current+1)
	fmt.Print(strings.Repeat(" ", 20)) // Clear rest of line
	
	fmt.Print("\033[u") // Restore cursor position
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

	// Show cycles dynamically based on what's displayed with proper padding
	fmt.Print(strings.Repeat(" ", LeftPadding))
	fmt.Print("   ") // Additional indent for cycle headers
	
	maxDisplay := 5 // Maximum cycles to display in headers
	for i := 1; i <= maxDisplay; i++ {
		if i > 1 {
			fmt.Print("        ")
		}
		switch {
		case i == currentCycle:
			fmt.Printf("▶ Cycle %d ◀", i)
		case i < currentCycle:
			fmt.Printf("✓ Cycle %d", i)
		default:
			fmt.Printf("Cycle %d", i)
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
	// Add left padding to column position
	totalCol := LeftPadding + col
	if totalCol > 0 {
		fmt.Printf("\033[%dC", totalCol) // Move right with padding
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
	fmt.Printf("%s   ", strings.Repeat(" ", LeftPadding))
	fmt.Print(left + strings.Repeat(middle, width) + right)
	fmt.Println()
}

func renderEmptyLine(boxWidth int) {
	fmt.Printf("%s   ", strings.Repeat(" ", LeftPadding))
	fmt.Print("│" + strings.Repeat(" ", boxWidth) + "│")
	fmt.Println()
}

func renderQuoteLine(line, emoji string, isFirstLine bool, boxWidth int) {
	fmt.Printf("%s   ", strings.Repeat(" ", LeftPadding))
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
