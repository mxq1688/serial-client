package main

import "strings"

type LogLevel int

const (
	LogDebug LogLevel = iota
	LogInfo
	LogWarning
	LogError
	LogFatal
)

type LogEntry struct {
	Level   LogLevel
	Message string
	Color   string
}

type LogParser struct{}

func NewLogParser() *LogParser {
	return &LogParser{}
}

func (p *LogParser) ParseLine(line string) LogEntry {
	entry := LogEntry{
		Message: line,
		Level:   LogInfo,
		Color:   "#FFFFFF",
	}

	upper := strings.ToUpper(line)
	if strings.Contains(upper, "ERROR") || strings.Contains(upper, " E ") {
		entry.Level = LogError
		entry.Color = "#FF5555"
	} else if strings.Contains(upper, "WARN") || strings.Contains(upper, " W ") {
		entry.Level = LogWarning
		entry.Color = "#FFAA00"
	} else if strings.Contains(upper, "DEBUG") || strings.Contains(upper, " D ") {
		entry.Level = LogDebug
		entry.Color = "#888888"
	} else if strings.Contains(upper, "FATAL") || strings.Contains(upper, " F ") {
		entry.Level = LogFatal
		entry.Color = "#FF0000"
	} else if strings.Contains(upper, "INFO") || strings.Contains(upper, " I ") {
		entry.Level = LogInfo
		entry.Color = "#00AA00"
	}

	return entry
}

func (p *LogParser) GetColorForLevel(level LogLevel) string {
	colors := map[LogLevel]string{
		LogDebug:   "#888888",
		LogInfo:    "#00AA00",
		LogWarning: "#FFAA00",
		LogError:   "#FF5555",
		LogFatal:   "#FF0000",
	}
	if color, ok := colors[level]; ok {
		return color
	}
	return "#FFFFFF"
}