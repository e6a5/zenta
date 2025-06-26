// Package storage handles all file I/O operations for zenta.
// It manages the ~/.zenta directory and provides safe access to logs and configuration.
package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/e6a5/zenta/internal/models"
)

const (
	ZentaDir   = ".zenta"
	LogsFile   = "logs.json"
	ConfigFile = "config.json"
	QuotesFile = "quotes.txt"
)

// Storage handles all file operations for zenta
type Storage struct {
	homeDir    string
	zentaDir   string
	logsPath   string
	configPath string
	quotesPath string
}

// New creates a new Storage instance
func New() (*Storage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("unable to find home directory: %w", err)
	}

	zentaDir := filepath.Join(homeDir, ZentaDir)

	s := &Storage{
		homeDir:    homeDir,
		zentaDir:   zentaDir,
		logsPath:   filepath.Join(zentaDir, LogsFile),
		configPath: filepath.Join(zentaDir, ConfigFile),
		quotesPath: filepath.Join(zentaDir, QuotesFile),
	}

	// Ensure zenta directory exists
	if err := s.ensureZentaDir(); err != nil {
		return nil, err
	}

	return s, nil
}

// ensureZentaDir creates the ~/.zenta directory if it doesn't exist
func (s *Storage) ensureZentaDir() error {
	if _, err := os.Stat(s.zentaDir); os.IsNotExist(err) {
		if err := os.MkdirAll(s.zentaDir, 0755); err != nil {
			return fmt.Errorf("unable to create zenta directory at %s: %w", s.zentaDir, err)
		}
	}
	return nil
}

// LoadLogs reads and parses the logs.json file
func (s *Storage) LoadLogs() ([]*models.LogEntry, error) {
	if _, err := os.Stat(s.logsPath); os.IsNotExist(err) {
		// File doesn't exist, return empty slice
		return []*models.LogEntry{}, nil
	}

	data, err := os.ReadFile(s.logsPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read logs file: %w", err)
	}

	var logs []*models.LogEntry
	if len(data) == 0 {
		return logs, nil
	}

	if err := json.Unmarshal(data, &logs); err != nil {
		return nil, fmt.Errorf("unable to parse logs file: %w", err)
	}

	return logs, nil
}

// SaveLogs writes the logs to logs.json file
func (s *Storage) SaveLogs(logs []*models.LogEntry) error {
	data, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal logs: %w", err)
	}

	if err := os.WriteFile(s.logsPath, data, 0644); err != nil {
		return fmt.Errorf("unable to write logs file: %w", err)
	}

	return nil
}

// AddLog appends a new log entry to the existing logs
func (s *Storage) AddLog(entry *models.LogEntry) error {
	logs, err := s.LoadLogs()
	if err != nil {
		return err
	}

	logs = append(logs, entry)
	return s.SaveLogs(logs)
}

// LoadConfig reads and parses the config.json file
func (s *Storage) LoadConfig() (*models.Config, error) {
	if _, err := os.Stat(s.configPath); os.IsNotExist(err) {
		// File doesn't exist, return default config and create file
		config := models.DefaultConfig()
		if err := s.SaveConfig(config); err != nil {
			return config, err // Return default even if save fails
		}
		return config, nil
	}

	data, err := os.ReadFile(s.configPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file: %w", err)
	}

	var config models.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unable to parse config file: %w", err)
	}

	return &config, nil
}

// SaveConfig writes the configuration to config.json file
func (s *Storage) SaveConfig(config *models.Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal config: %w", err)
	}

	if err := os.WriteFile(s.configPath, data, 0644); err != nil {
		return fmt.Errorf("unable to write config file: %w", err)
	}

	return nil
}
