package breathing

import (
	"testing"
)

func TestNewSession(t *testing.T) {
	s := NewSession()

	if s.Cycles != 3 {
		t.Errorf("Expected default cycles to be 3, got %d", s.Cycles)
	}
	if s.ShowQuote != true {
		t.Errorf("Expected ShowQuote to be true by default, got %v", s.ShowQuote)
	}
	if s.InhaleDur != 4 {
		t.Errorf("Expected InhaleDur to be 4, got %d", s.InhaleDur)
	}
	if s.HoldDur != 4 {
		t.Errorf("Expected HoldDur to be 4, got %d", s.HoldDur)
	}
	if s.ExhaleDur != 4 {
		t.Errorf("Expected ExhaleDur to be 4, got %d", s.ExhaleDur)
	}
	if s.SimpleMode != shouldUseSimpleAnimation() {
		t.Errorf("Expected SimpleMode to match default from shouldUseSimpleAnimation()")
	}
}

func TestParseArgs(t *testing.T) {
	defaultSimple := shouldUseSimpleAnimation()

	testCases := []struct {
		name           string
		args           []string
		expectedCycles int
		expectedQuote  bool
		expectedSimple bool
	}{
		{"no args", []string{}, 3, true, defaultSimple},
		{"quick", []string{"--quick"}, 1, true, defaultSimple},
		{"extended", []string{"--extended"}, 5, true, defaultSimple},
		{"silent", []string{"--silent"}, 3, false, defaultSimple},
		{"simple", []string{"--simple"}, 3, true, true},
		{"complex", []string{"--complex"}, 3, true, false},
		{"-q", []string{"-q"}, 1, true, defaultSimple},
		{"-e", []string{"-e"}, 5, true, defaultSimple},
		{"-s", []string{"-s"}, 3, false, defaultSimple},
		{"combo", []string{"--quick", "--silent", "--simple"}, 1, false, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSession()
			s.ParseArgs(tc.args)

			if s.Cycles != tc.expectedCycles {
				t.Errorf("For args %v, expected cycles to be %d, got %d", tc.args, tc.expectedCycles, s.Cycles)
			}
			if s.ShowQuote != tc.expectedQuote {
				t.Errorf("For args %v, expected ShowQuote to be %v, got %v", tc.args, tc.expectedQuote, s.ShowQuote)
			}
			if s.SimpleMode != tc.expectedSimple {
				t.Errorf("For args %v, expected SimpleMode to be %v, got %v", tc.args, tc.expectedSimple, s.SimpleMode)
			}
		})
	}
}

func TestShouldShowQuote(t *testing.T) {
	s := NewSession()

	s.ShowQuote = true
	if !s.ShouldShowQuote() {
		t.Error("ShouldShowQuote should return true when s.ShowQuote is true")
	}

	s.ShowQuote = false
	if s.ShouldShowQuote() {
		t.Error("ShouldShowQuote should return false when s.ShowQuote is false")
	}
}
