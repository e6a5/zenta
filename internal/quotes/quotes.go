// Package quotes provides mindfulness quotes and quote management functionality.
// It offers a collection of built-in quotes inspired by Zen, Stoicism, and mindfulness practices.
package quotes

import (
	"crypto/rand"
	"math/big"
)

// Built-in mindfulness quotes inspired by Zen, Stoicism, and mindfulness practices
var builtinQuotes = []string{
	"🧘 Take a breath. This moment is all there is.",
	"🌱 What you resist persists. What you accept transforms.",
	"⭐ The present moment is the only time over which we have dominion. - Thich Nhat Hanh",
	"🍃 Wherever you are, be there totally. - Eckhart Tolle",
	"🌊 You have power over your mind—not outside events. Realize this, and you will find strength. - Marcus Aurelius",
	"🎯 The best way to take care of the future is to take care of the present moment.",
	"🌸 Peace comes from within. Do not seek it without. - Buddha",
	"🕯️ Between stimulus and response there is a space. In that space is our power to choose our response.",
	"🌿 Mindfulness is about being fully awake in our lives.",
	"⚡ This too shall pass. Notice what arises, and let it go.",
	"🎋 The mind is everything. What you think you become. - Buddha",
	"🌅 Each morning we are born again. What we do today is what matters most.",
	"🪨 Be like water making its way through cracks. - Bruce Lee",
	"🌊 Flow with whatever may happen and let your mind be free.",
	"⭐ The quieter you become, the more you are able to hear.",
	"🌱 In the beginner's mind there are many possibilities, in the expert's mind there are few. - Shunryu Suzuki",
	"🕊️ Let go or be dragged. - Zen Proverb",
	"🌸 The only way out is through.",
	"🎯 Focus on the step in front of you, not the whole staircase.",
	"🌿 Breathe in calm, breathe out chaos.",
	"⚖️ Balance is not something you find, it's something you create.",
	"🌊 When you realize nothing is lacking, the whole world belongs to you. - Lao Tzu",
	"🪷 Muddy water is best cleared by leaving it alone. - Alan Watts",
	"🌅 Every moment is a fresh beginning. - T.S. Eliot",
	"🎋 Simplicity is the ultimate sophistication.",
}

// QuoteService handles quote retrieval and management
type QuoteService struct{}

// New creates a new QuoteService
func New() *QuoteService {
	return &QuoteService{}
}

// GetRandomQuote returns a random quote from the built-in collection
func (qs *QuoteService) GetRandomQuote() string {
	if len(builtinQuotes) == 0 {
		return "🧘 Take a breath. This moment is all there is."
	}

	// Use crypto/rand for better randomness
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(builtinQuotes))))
	if err != nil {
		// Fallback to first quote if crypto/rand fails
		return builtinQuotes[0]
	}

	return builtinQuotes[n.Int64()]
}

// GetAllQuotes returns all built-in quotes (useful for testing or exporting)
func (qs *QuoteService) GetAllQuotes() []string {
	// Return a copy to prevent external modification
	quotes := make([]string, len(builtinQuotes))
	copy(quotes, builtinQuotes)
	return quotes
}

// QuoteCount returns the number of available quotes
func (qs *QuoteService) QuoteCount() int {
	return len(builtinQuotes)
}
