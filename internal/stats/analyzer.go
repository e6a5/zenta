// Package stats provides statistical analysis functionality for zenta log entries.
// It generates analytics, charts, and insights from user mindfulness data.
package stats

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/e6a5/zenta/internal/models"
)

const (
	HoursPerDay    = 24
	StatsLineWidth = 40
	MaxBarWidth    = 20
)

// Analyzer handles statistical analysis of log entries
type Analyzer struct {
	logs []*models.LogEntry
}

// New creates a new Analyzer with the provided logs
func New(logs []*models.LogEntry) *Analyzer {
	return &Analyzer{
		logs: logs,
	}
}

// StatsResult holds various statistics about the logs
type StatsResult struct {
	TotalEntries       int
	DistractionCount   int
	ReflectionCount    int
	InsightCount       int
	SessionCount       int
	TimeRange          string
	HourlyDistribution map[int]int    // hour -> count
	DailyDistribution  map[string]int // date -> count
}

// GenerateStats creates comprehensive statistics for the given period
func (a *Analyzer) GenerateStats(period string) *StatsResult {
	filteredLogs := a.filterByPeriod(period)

	result := &StatsResult{
		TotalEntries:       len(filteredLogs),
		HourlyDistribution: make(map[int]int),
		DailyDistribution:  make(map[string]int),
	}

	if len(filteredLogs) == 0 {
		result.TimeRange = "No data available"
		return result
	}

	// Count by type and collect time distributions
	for _, log := range filteredLogs {
		switch log.Type {
		case models.LogTypeDistraction:
			result.DistractionCount++
		case models.LogTypeReflection:
			result.ReflectionCount++
		case models.LogTypeInsight:
			result.InsightCount++
		case models.LogTypeSessionStart:
			result.SessionCount++
		case models.LogTypeSessionEnd:
			// SessionEnd doesn't increment SessionCount as SessionStart already does
		}

		// Hour distribution (0-23)
		hour := log.Timestamp.Hour()
		result.HourlyDistribution[hour]++

		// Daily distribution
		date := log.Timestamp.Format("2006-01-02")
		result.DailyDistribution[date]++
	}

	// Set time range
	if len(filteredLogs) > 0 {
		earliest := filteredLogs[0].Timestamp
		latest := filteredLogs[len(filteredLogs)-1].Timestamp
		result.TimeRange = fmt.Sprintf("%s to %s",
			earliest.Format("Jan 2"),
			latest.Format("Jan 2, 2006"))
	}

	return result
}

// filterByPeriod filters logs based on the specified time period
func (a *Analyzer) filterByPeriod(period string) []*models.LogEntry {
	if period == "all" || period == "" {
		return a.logs
	}

	now := time.Now()
	var cutoff time.Time

	switch period {
	case "today":
		cutoff = now.Truncate(HoursPerDay * time.Hour)
	case "week":
		// Go back to start of week (Monday)
		weekday := int(now.Weekday())
		if weekday == 0 { // Sunday
			weekday = 7
		}
		cutoff = now.AddDate(0, 0, -(weekday - 1)).Truncate(HoursPerDay * time.Hour)
	case "month":
		cutoff = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	default:
		// Unknown period, return all
		return a.logs
	}

	filtered := make([]*models.LogEntry, 0)
	for _, log := range a.logs {
		if log.Timestamp.After(cutoff) || log.Timestamp.Equal(cutoff) {
			filtered = append(filtered, log)
		}
	}

	return filtered
}

// FormatStats returns a formatted string representation of the statistics
func (a *Analyzer) FormatStats(result *StatsResult, period string) string {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("📊 Zenta Statistics (%s)\n", period))
	output.WriteString(strings.Repeat("─", StatsLineWidth))
	output.WriteString("\n\n")

	if result.TotalEntries == 0 {
		output.WriteString("No logs found for this period.\n")
		output.WriteString("Use 'zenta log \"reason\"' to start tracking.\n")
		return output.String()
	}

	// Summary
	output.WriteString(fmt.Sprintf("📅 Time Range: %s\n", result.TimeRange))
	output.WriteString(fmt.Sprintf("📝 Total Entries: %d\n\n", result.TotalEntries))

	// Breakdown by type
	output.WriteString("Entry Types:\n")
	output.WriteString(fmt.Sprintf("  🔴 Distractions: %d\n", result.DistractionCount))
	output.WriteString(fmt.Sprintf("  🤔 Reflections:  %d\n", result.ReflectionCount))
	output.WriteString(fmt.Sprintf("  💡 Insights:     %d\n", result.InsightCount))
	output.WriteString(fmt.Sprintf("  ⏰ Sessions:     %d\n\n", result.SessionCount))

	// Hourly distribution (simple ASCII chart)
	if len(result.HourlyDistribution) > 0 {
		output.WriteString("Hourly Activity:\n")
		output.WriteString(a.formatHourlyChart(result.HourlyDistribution))
		output.WriteString("\n")
	}

	return output.String()
}

// formatHourlyChart creates a simple ASCII bar chart for hourly distribution
func (a *Analyzer) formatHourlyChart(hourly map[int]int) string {
	var output strings.Builder

	// Find max count for scaling
	maxCount := 0
	for _, count := range hourly {
		if count > maxCount {
			maxCount = count
		}
	}

	if maxCount == 0 {
		return "  No activity data available\n"
	}

	// Create chart for hours with activity
	type hourData struct {
		hour  int
		count int
	}

	var hours []hourData
	for hour, count := range hourly {
		if count > 0 {
			hours = append(hours, hourData{hour, count})
		}
	}

	// Sort by hour
	sort.Slice(hours, func(i, j int) bool {
		return hours[i].hour < hours[j].hour
	})

	// Generate bars (max width of 20 chars)
	for _, h := range hours {
		barWidth := (h.count * MaxBarWidth) / maxCount
		if barWidth == 0 && h.count > 0 {
			barWidth = 1 // Ensure at least one character for non-zero counts
		}

		bar := strings.Repeat("█", barWidth)
		output.WriteString(fmt.Sprintf("  %02d:00 %s %d\n", h.hour, bar, h.count))
	}

	return output.String()
}
