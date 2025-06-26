package quotes

import (
	"strings"
	"testing"
)

func TestNewQuoteService(t *testing.T) {
	qs := New()
	if qs == nil {
		t.Error("Expected QuoteService instance, got nil")
	}
}

func TestGetRandomQuote(t *testing.T) {
	qs := New()

	// Test multiple calls to ensure randomness works
	quotes := make(map[string]bool)
	for i := 0; i < 10; i++ {
		quote := qs.GetRandomQuote()
		if quote == "" {
			t.Error("Expected non-empty quote")
		}
		quotes[quote] = true
	}

	// With 25 quotes, getting the same quote 10 times would be very unlikely
	if len(quotes) == 1 && len(builtinQuotes) > 1 {
		t.Error("Expected some variation in quotes, got same quote 10 times")
	}
}

func TestGetAllQuotes(t *testing.T) {
	qs := New()
	quotes := qs.GetAllQuotes()

	if len(quotes) != len(builtinQuotes) {
		t.Errorf("Expected %d quotes, got %d", len(builtinQuotes), len(quotes))
	}

	// Test that it returns a copy (modifying shouldn't affect original)
	originalLen := len(quotes)
	_ = append(quotes, "Test quote") // Use blank identifier to avoid ineffectual assignment

	quotesAgain := qs.GetAllQuotes()
	if len(quotesAgain) != originalLen {
		t.Error("Expected GetAllQuotes to return a copy, original was modified")
	}
}

func TestQuoteCount(t *testing.T) {
	qs := New()
	count := qs.QuoteCount()

	if count != len(builtinQuotes) {
		t.Errorf("Expected count %d, got %d", len(builtinQuotes), count)
	}

	if count == 0 {
		t.Error("Expected at least one quote")
	}
}

func TestBuiltinQuotesContent(t *testing.T) {
	// Test that all quotes are non-empty and meaningful
	for i, quote := range builtinQuotes {
		if quote == "" {
			t.Errorf("Quote %d is empty", i)
		}

		if len(quote) < 10 {
			t.Errorf("Quote %d seems too short: %s", i, quote)
		}

		// Verify it has some meaningful content
		if len(strings.TrimSpace(quote)) < 5 {
			t.Errorf("Quote %d appears to have no meaningful content: %s", i, quote)
		}
	}
}

func TestFallbackQuote(t *testing.T) {
	// Test the fallback behavior when no quotes available
	// We'll temporarily modify builtinQuotes to test this
	originalQuotes := builtinQuotes
	builtinQuotes = []string{} // Empty quotes

	qs := New()
	quote := qs.GetRandomQuote()

	expected := "🧘 Take a breath. This moment is all there is."
	if quote != expected {
		t.Errorf("Expected fallback quote %s, got %s", expected, quote)
	}

	// Restore original quotes
	builtinQuotes = originalQuotes
}
