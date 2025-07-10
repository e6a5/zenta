// Package breathing provides mindful breathing session functionality.
// It handles breathing visualization, timing, and user interaction.
package breathing

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
)

// Constants for breathing visualization
const (
	LeftPadding    = 4 // Left margin for all content
	SectionSpacing = 1 // Space between sections
	BottomPadding  = 2 // Extra spacing at bottom
	RestDuration   = 2 * time.Second
)

// Session represents a breathing session configuration
type Session struct {
	Cycles     int
	ShowQuote  bool
	InhaleDur  int
	HoldDur    int
	ExhaleDur  int
	RestDur    time.Duration
	SimpleMode bool
}

// sigint defines the signals to listen for to restore the cursor.
// It's a variable to allow for different signals on different OSes,
// though currently it's the same for all.
var sigint = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

// sigexit handles the process of exiting gracefully.
func sigexit(s os.Signal) {
	os.Exit(0)
}

// NewSession creates a new breathing session with default settings
func NewSession() *Session {
	return &Session{
		Cycles:     3,
		ShowQuote:  true,
		InhaleDur:  4,
		HoldDur:    4,
		ExhaleDur:  4,
		RestDur:    RestDuration,
		SimpleMode: shouldUseSimpleAnimation(),
	}
}

// ParseArgs parses command line arguments and configures the session
func (s *Session) ParseArgs(args []string) {
	for _, arg := range args {
		switch arg {
		case "--quick", "-q":
			s.Cycles = 1
		case "--extended", "-e":
			s.Cycles = 5
		case "--silent", "-s":
			s.ShowQuote = false
		case "--complex":
			s.SimpleMode = false
		case "--simple":
			s.SimpleMode = true
		}
	}
}

func (s Session) HideCursor() func() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, sigint...)
	fmt.Print("\033[?25l")
	go func() {
		s := <-sig
		fmt.Println("\033[?25h")
		sigexit(s)
	}()
	return func() { fmt.Println("\033[?25h") }
}

// Start begins the breathing session with visualization
func (s *Session) Start() {
	// Add top spacing
	fmt.Println()

	// Go straight into breathing - no interruptions
	PrintWithPadding("   Let's breathe 🌸")
	AddSectionSpacing()

	// Use simple animation for better compatibility or if requested
	if s.SimpleMode {
		s.drawSimpleBreathingSession()
	} else {
		// Reserve dedicated space for breathing visualization (12 lines)
		fmt.Println() // Guidance text line
		for i := 0; i < 10; i++ {
			fmt.Println() // Circle area
		}
		fmt.Println() // Bottom buffer

		// Move cursor back to start of breathing area
		fmt.Print("\033[12A")

		// One lung, multiple breaths
		s.drawContinuousBreathingSession()

		// Move cursor to end of breathing area
		fmt.Print("\033[12B")
	}

	AddSectionSpacing()
}

// shouldUseSimpleAnimation determines if we should use simple animation for compatibility
func shouldUseSimpleAnimation() bool {
	// Check if we're on macOS Terminal which has ANSI issues
	if runtime.GOOS == "darwin" {
		if term := os.Getenv("TERM_PROGRAM"); term == "Apple_Terminal" {
			return true
		}
	}

	// Modern terminal multiplexers like tmux and screen handle complex animations fine
	return false
}

// drawSimpleBreathingSession draws a simple line-based breathing animation for compatibility
func (s *Session) drawSimpleBreathingSession() {
	for cycle := 1; cycle <= s.Cycles; cycle++ {
		// Inhale phase
		s.drawSimplePhase("🌬️", "Breathe in gently...", s.InhaleDur, []string{
			"·", "○", "○○", "●○○", "●●○○", "●●●○", "●●●●",
		})

		// Hold phase
		s.drawSimplePhase("✨", "Hold softly...", s.HoldDur, []string{
			"●●●●", "●●●●", "●●●●", "●●●●",
		})

		// Exhale phase
		s.drawSimplePhase("🌸", "Release slowly...", s.ExhaleDur, []string{
			"●●●●", "●●●○", "●●○○", "●○○○", "○○○○", "○○  ", "○   ",
		})

		// Rest phase
		s.drawSimplePhase("🕯️", "Rest in emptiness...", s.HoldDur, []string{
			"·", "·", "·", "·",
		})

		// Brief pause between cycles
		if cycle < s.Cycles {
			PrintWithPadding("   💫 Feel the rhythm... continuing...")
			time.Sleep(s.RestDur)
			fmt.Println()
		}
	}

	PrintWithPadding("   🙏 Complete")
	fmt.Println()
}

// drawSimplePhase draws a single breathing phase with progressive animation
func (s *Session) drawSimplePhase(emoji, instruction string, duration int, patterns []string) {
	if checkForExit() {
		return
	}

	// Show the phase instruction
	PrintWithPadding(fmt.Sprintf("   %s %s", emoji, instruction))

	// Reserve space for the animation
	PrintWithPadding("      ")

	// Animate through the duration
	for second := 1; second <= duration; second++ {
		// Choose pattern based on progress through the phase
		patternIndex := (second - 1) * (len(patterns) - 1) / (duration - 1)
		if patternIndex >= len(patterns) {
			patternIndex = len(patterns) - 1
		}
		pattern := patterns[patternIndex]

		// Go back up one line and overwrite the animation line
		fmt.Print("\033[1A")
		PrintWithPadding(fmt.Sprintf("      %s", pattern))

		time.Sleep(time.Second)
	}

	fmt.Println() // Add spacing after phase
}

// ShouldShowQuote returns whether to show quote after breathing
func (s *Session) ShouldShowQuote() bool {
	return s.ShowQuote
}

// drawContinuousBreathingSession draws one lung breathing continuously through multiple cycles
func (s *Session) drawContinuousBreathingSession() {
	// Breathing phases that feel natural
	phases := []struct {
		name        string
		emoji       string
		duration    int
		instruction string
		breathType  string
	}{
		{"inhale", "🌬️", s.InhaleDur, "Breathe in gently, let your body expand...", "expand"},
		{"hold", "✨", s.HoldDur, "Hold softly, feel the fullness...", "full"},
		{"exhale", "🌸", s.ExhaleDur, "Release slowly, let everything go...", "contract"},
		{"rest", "🕯️", s.HoldDur, "Rest in the emptiness, be present...", "empty"},
	}

	// One lung, breathing continuously through all cycles
	for cycle := 1; cycle <= s.Cycles; cycle++ {
		for _, phase := range phases {
			if checkForExit() {
				return
			}

			// Show gentle guidance for each phase
			showBreathingGuidance(phase.emoji, phase.name, phase.instruction)

			// Animate the same breathing circle for this phase
			animateBreathingCircle(phase.breathType, phase.duration)
		}

		// Brief pause between breathing cycles
		if cycle < s.Cycles {
			showBreathingGuidance("💫", "rest", "Feel the rhythm... continuing...")
			time.Sleep(s.RestDur)
		}
	}

	clearBreathingDisplay()
}

// checkForExit checks if user pressed 'q' to exit (simplified for now)
func checkForExit() bool {
	// For now, we'll implement a simple version
	// In a full implementation, this would use non-blocking input
	return false
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
		case "expand":
			// Gradually expand from 1 to 4
			progress := float64(second) / float64(duration)
			circleSize = 1 + int(progress*3)
			circleChar = "○"
		case "full":
			// Stay at full size
			circleSize = 4
			circleChar = "●"
		case "contract":
			// Gradually contract from 4 to 1
			progress := float64(second) / float64(duration)
			circleSize = 4 - int(progress*3)
			circleChar = "○"
		case "empty":
			// Stay small
			circleSize = 1
			circleChar = "·"
		}

		// Clear previous circle and draw new one
		clearCircleAreaRelative(centerRowOffset, centerCol)
		drawBreathingCircleRelative(centerRowOffset, centerCol, circleSize, circleChar)

		time.Sleep(time.Second)
	}
}

// clearCircleAreaRelative clears the breathing circle area relative to current cursor position
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

// PrintWithPadding prints text with consistent left padding
func PrintWithPadding(text string) {
	fmt.Printf("%s%s\n", strings.Repeat(" ", LeftPadding), text)
}

// AddSectionSpacing adds consistent spacing between sections
func AddSectionSpacing() {
	for i := 0; i < SectionSpacing; i++ {
		fmt.Println()
	}
}

// AddBottomPadding adds consistent bottom padding
func AddBottomPadding() {
	for i := 0; i < BottomPadding; i++ {
		fmt.Println()
	}
}

// StartAnchor provides an immediate, intuitive manual breathing session.
func (s *Session) StartAnchor() {
	// Hide cursor and restore on exit
	defer s.HideCursor()()

	fmt.Println()
	PrintWithPadding("   🌸")
	PrintWithPadding("   Let the rhythm guide you. [SPACE] to switch phase, [q] to quit.")
	fmt.Println()

	if err := s.runAnchorBreathing(); err != nil {
		// If real-time mode fails, print an informative error.
		PrintWithPadding("Error: This terminal does not support this mode.")
		PrintWithPadding("The 'zenta now' command is a great alternative.")
		fmt.Println()
		return
	}

	// Clean up the final line of the visualizer and add mindful spacing.
	fmt.Print("\r" + strings.Repeat(" ", 80) + "\r")
	PrintWithPadding("   🙏 Carry this calm with you.")
	AddBottomPadding()
}

// runAnchorBreathing sets up the terminal and runs the new pacer logic.
func (s *Session) runAnchorBreathing() error {
	fd := int(os.Stdin.Fd())
	if !term.IsTerminal(fd) {
		return fmt.Errorf("stdin is not a terminal")
	}

	// Switch to raw mode to read single key presses.
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	// Ensure the terminal state is always restored.
	defer func() {
		if err := term.Restore(fd, oldState); err != nil {
			fmt.Fprintf(os.Stderr, "Error restoring terminal: %v\n", err)
		}
	}()

	var (
		breathSize int
		maxSize    = 40
		phase      = "inhale" // "inhale", "exhale", or "paused"
		keyPress   = make(chan byte, 1)
	)

	// Goroutine to read key presses.
	go func() {
		defer close(keyPress)
		for {
			buffer := make([]byte, 1)
			if count, err := os.Stdin.Read(buffer); err != nil || count == 0 {
				return
			}
			keyPress <- buffer[0]
		}
	}()

	// The main animation loop.
	for {
		// 1. Check for user input to change phase.
		select {
		case key, ok := <-keyPress:
			if !ok {
				return nil // Channel closed.
			}
			switch key {
			case ' ':
				if phase == "inhale" {
					phase = "exhale" // Switch to exhaling.
				} else if phase == "paused" {
					phase = "inhale" // Start new cycle from paused state.
				}
			case 'q', 'Q', 3: // Ctrl+C
				return nil
			}
		default:
			// No input, continue the current phase.
		}

		// 2. Update the breath size based on the current phase.
		switch phase {
		case "inhale":
			if breathSize < maxSize {
				breathSize++
			}
		case "exhale":
			if breathSize > 0 {
				breathSize--
			} else {
				phase = "paused" // Breath is empty, wait for user.
			}
		case "paused":
			// Do nothing, wait for the user to press space.
		}

		// 3. Render the visual and pause.
		s.drawBreathingVisual(breathSize, maxSize, phase)
		time.Sleep(90 * time.Millisecond) // Slower, more calming pace.
	}
}

// drawBreathingVisual renders a simple, minimalist line of dots.
func (s *Session) drawBreathingVisual(size, visualMaxWidth int, phase string) {
	var bar strings.Builder
	bar.WriteString(strings.Repeat(" ", LeftPadding)) // Indent

	// Display the current phase, padded for alignment.
	phaseText := fmt.Sprintf("%-8s", phase)
	bar.WriteString(phaseText)
	bar.WriteString(" [")

	displaySize := size
	if displaySize > visualMaxWidth {
		displaySize = visualMaxWidth
	}

	for i := 0; i < displaySize; i++ {
		bar.WriteString("●")
	}

	if displaySize < visualMaxWidth {
		bar.WriteString("○")
	}

	for i := 0; i < visualMaxWidth-displaySize-1; i++ {
		bar.WriteString("·")
	}

	bar.WriteString("]")

	// Overwrite the current line with the new visual.
	fmt.Print("\r" + bar.String() + " ")
}
