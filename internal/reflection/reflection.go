// Package reflection provides prompts for mindful reflection sessions.
package reflection

// PromptSet contains the text for a reflection session.
type PromptSet struct {
	Title        string
	Instructions []string
	PromptTitle  string
	Prompts      []string
	Closing      []string
}

// GetDefaultPrompts returns the default set of reflection prompts.
func GetDefaultPrompts() PromptSet {
	return PromptSet{
		Title: "🕯️  Evening Reflection",
		Instructions: []string{
			"   Close your eyes for a moment...",
			"   Take three deep breaths...",
		},
		PromptTitle: "   📝 Gentle reflection:",
		Prompts: []string{
			"      • What thoughts kept pulling you away today?",
			"      • Were there moments when you were truly present?",
			"      • What patterns do you notice in your mind?",
		},
		Closing: []string{
			"   These are just thoughts. They come and go like clouds.",
			"   The noticing itself is the practice. 🙏",
		},
	}
}
