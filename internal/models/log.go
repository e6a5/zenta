package models

import (
	"time"

	"github.com/google/uuid"
)

// LogType represents the type of log entry
type LogType string

const (
	LogTypeDistraction  LogType = "distraction"
	LogTypeReflection   LogType = "reflection"
	LogTypeInsight      LogType = "insight"
	LogTypeSessionStart LogType = "session_start"
	LogTypeSessionEnd   LogType = "session_end"
)

// LogEntry represents a single log entry in the system
type LogEntry struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Type      LogType   `json:"type"`
	Note      string    `json:"note"`
	SessionID string    `json:"session_id,omitempty"`
	Duration  int       `json:"duration,omitempty"` // in seconds
}

// NewLogEntry creates a new log entry with generated ID and current timestamp
func NewLogEntry(logType LogType, note string) *LogEntry {
	return &LogEntry{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Type:      logType,
		Note:      note,
	}
}

// NewSessionLogEntry creates a log entry for session start/end with session ID
func NewSessionLogEntry(logType LogType, note string, sessionID string, duration int) *LogEntry {
	return &LogEntry{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Type:      logType,
		Note:      note,
		SessionID: sessionID,
		Duration:  duration,
	}
}
