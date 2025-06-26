package models

import (
	"testing"
	"time"
)

func TestNewLogEntry(t *testing.T) {
	note := "Test log entry"
	entry := NewLogEntry(LogTypeDistraction, note)

	if entry.ID == "" {
		t.Error("Expected ID to be generated, got empty string")
	}

	if entry.Type != LogTypeDistraction {
		t.Errorf("Expected type %s, got %s", LogTypeDistraction, entry.Type)
	}

	if entry.Note != note {
		t.Errorf("Expected note %s, got %s", note, entry.Note)
	}

	if entry.Timestamp.IsZero() {
		t.Error("Expected timestamp to be set, got zero time")
	}

	// Timestamp should be recent (within last second)
	if time.Since(entry.Timestamp) > time.Second {
		t.Error("Expected timestamp to be recent")
	}
}

func TestNewSessionLogEntry(t *testing.T) {
	note := "Session start"
	sessionID := "test-session-123"
	duration := 1800 // 30 minutes

	entry := NewSessionLogEntry(LogTypeSessionStart, note, sessionID, duration)

	if entry.SessionID != sessionID {
		t.Errorf("Expected session ID %s, got %s", sessionID, entry.SessionID)
	}

	if entry.Duration != duration {
		t.Errorf("Expected duration %d, got %d", duration, entry.Duration)
	}

	if entry.Type != LogTypeSessionStart {
		t.Errorf("Expected type %s, got %s", LogTypeSessionStart, entry.Type)
	}
}

func TestLogTypes(t *testing.T) {
	tests := []struct {
		logType LogType
		want    string
	}{
		{LogTypeDistraction, "distraction"},
		{LogTypeReflection, "reflection"},
		{LogTypeInsight, "insight"},
		{LogTypeSessionStart, "session_start"},
		{LogTypeSessionEnd, "session_end"},
	}

	for _, tt := range tests {
		t.Run(string(tt.logType), func(t *testing.T) {
			if string(tt.logType) != tt.want {
				t.Errorf("Expected %s, got %s", tt.want, string(tt.logType))
			}
		})
	}
}
