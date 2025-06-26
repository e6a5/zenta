// Package models defines core data structures and types for zenta.
// It contains configuration, logging, and other fundamental data models.
package models

// Config represents the user configuration for zenta
type Config struct {
	TimerDuration       int    `json:"timer_duration"` // in minutes
	BreakDuration       int    `json:"break_duration"` // in minutes
	SoundEnabled        bool   `json:"sound_enabled"`
	NotificationMethod  string `json:"notification_method"` // "bell", "message", "silent"
	Timezone            string `json:"timezone"`            // "auto" or specific timezone
	CustomQuotesEnabled bool   `json:"custom_quotes_enabled"`
	StatsDefaultPeriod  string `json:"stats_default_period"` // "week", "month", "all"
}

// DefaultConfig returns a configuration with sensible defaults
func DefaultConfig() *Config {
	return &Config{
		TimerDuration:       45,
		BreakDuration:       10,
		SoundEnabled:        true,
		NotificationMethod:  "message",
		Timezone:            "auto",
		CustomQuotesEnabled: false,
		StatsDefaultPeriod:  "week",
	}
}
