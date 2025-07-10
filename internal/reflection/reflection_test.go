package reflection

import (
	"testing"
)

func TestGetDefaultPrompts(t *testing.T) {
	prompts := GetDefaultPrompts()

	if prompts.Title == "" {
		t.Error("Expected Title to not be empty")
	}
	if len(prompts.Instructions) == 0 {
		t.Error("Expected Instructions to not be empty")
	}
	if prompts.PromptTitle == "" {
		t.Error("Expected PromptTitle to not be empty")
	}
	if len(prompts.Prompts) == 0 {
		t.Error("Expected Prompts to not be empty")
	}
	if len(prompts.Closing) == 0 {
		t.Error("Expected Closing to not be empty")
	}
}
