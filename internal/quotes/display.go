// Package quotes handles quote display and rendering.
// It provides beautiful quote formatting with typing animation.
package quotes

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// DisplayBeautifully displays a quote with beautiful formatting and typing animation
func DisplayBeautifully(quote string) {
	emoji, quoteText := parseQuoteEmoji(quote)
	lines := wrapQuoteText(quoteText)
	boxWidth := calculateBoxWidth(lines)
	renderQuoteBoxWithTyping(lines, emoji, boxWidth)
}

// parseQuoteEmoji extracts emoji from quote if present
func parseQuoteEmoji(quote string) (string, string) {
	parts := strings.Fields(quote)
	if len(parts) > 0 && isEmoji(rune(parts[0][0])) {
		return parts[0], strings.Join(parts[1:], " ")
	}
	return "💭", quote
}

// wrapQuoteText wraps text to fit within terminal width
func wrapQuoteText(quoteText string) []string {
	maxWidth := 50 // Maximum line width for quotes
	words := strings.Fields(quoteText)
	var lines []string
	var currentLine []string
	currentLength := 0

	for _, word := range words {
		wordLength := len(word)
		if currentLength+wordLength+len(currentLine) > maxWidth && len(currentLine) > 0 {
			lines = append(lines, strings.Join(currentLine, " "))
			currentLine = []string{word}
			currentLength = wordLength
		} else {
			currentLine = append(currentLine, word)
			currentLength += wordLength
		}
	}

	if len(currentLine) > 0 {
		lines = append(lines, strings.Join(currentLine, " "))
	}

	return lines
}

// calculateBoxWidth calculates the width needed for the quote box
func calculateBoxWidth(lines []string) int {
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	return maxLen + 8 // Add padding for borders and spacing
}

// renderQuoteBoxWithTyping renders the quote box with typing animation
func renderQuoteBoxWithTyping(lines []string, emoji string, boxWidth int) {
	leftPadding := 4
	padding := strings.Repeat(" ", leftPadding)

	fmt.Printf("%s┌%s┐\n", padding, strings.Repeat("─", boxWidth-2))

	for i, line := range lines {
		renderQuoteLineSimpleWithTyping(line, emoji, i == 0)
	}

	fmt.Printf("%s└%s┘\n", padding, strings.Repeat("─", boxWidth-2))
}

// renderQuoteLineSimpleWithTyping renders a single line with typing animation
func renderQuoteLineSimpleWithTyping(line, emoji string, isFirstLine bool) {
	leftPadding := 4
	padding := strings.Repeat(" ", leftPadding)

	fmt.Printf("%s│ ", padding)

	if isFirstLine {
		fmt.Printf("%s ", emoji)
	} else {
		fmt.Print("  ") // Two spaces to align with emoji width
	}

	// Type out the line character by character
	for _, char := range line {
		fmt.Print(string(char))
		if !unicode.IsSpace(char) {
			time.Sleep(15 * time.Millisecond) // Slightly faster typing
		}
	}

	// Add trailing spaces to fill the box width (simplified)
	spaces := 40 - len(line) // Approximate trailing space
	if spaces > 0 {
		fmt.Print(strings.Repeat(" ", spaces))
	}

	fmt.Println(" │")
}

// isEmoji checks if a rune is an emoji character
func isEmoji(r rune) bool {
	// Simple check for emoji ranges
	return (r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
		(r >= 0x1F300 && r <= 0x1F5FF) || // Misc Symbols
		(r >= 0x1F680 && r <= 0x1F6FF) || // Transport
		(r >= 0x2600 && r <= 0x26FF) || // Misc symbols
		(r >= 0x2700 && r <= 0x27BF) || // Dingbats
		(r >= 0xFE00 && r <= 0xFE0F) || // Variation Selectors
		(r >= 0x1F900 && r <= 0x1F9FF) // Supplemental Symbols
}
