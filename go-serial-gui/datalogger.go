package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// DataLogger handles data recording and playback
type DataLogger struct {
	filePath string
	data     []LogEntry
}

type LogEntry struct {
	Timestamp time.Time `json:"timestamp\u0022`
	Direction string    `json:"direction\u0022` // TX or RX
	Data      string    `json:"data\u0022`
	Hex       string    `json:"hex\u0022`
}

func NewDataLogger(path string) *DataLogger {
	return &DataLogger{
		filePath: path,
		data:     []LogEntry{},
	}
}

func (dl *DataLogger) LogData(direction, data string) {
	entry := LogEntry{
		Timestamp: time.Now(),
		Direction: direction,
		Data:      data,
		Hex:       fmt.Sprintf("%X", data),
	}
	dl.data = append(dl.data, entry)
}

func (dl *DataLogger) SaveToJSON() error {
	file, err := os.Create(dl.filePath + ".json")
	if err != nil {
		return err
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(dl.data)
}

func (dl *DataLogger) SaveToCSV() error {
	file, err := os.Create(dl.filePath + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()
	
	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	// Write header
	writer.Write([]string{"Timestamp", "Direction", "Data", "Hex"})
	
	// Write data
	for _, entry := range dl.data {
		writer.Write([]string{
			entry.Timestamp.Format(time.RFC3339),
			entry.Direction,
			entry.Data,
			entry.Hex,
		})
	}
	
	return nil
}

func (dl *DataLogger) LoadFromJSON(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	return decoder.Decode(&dl.data)
}

func (dl *DataLogger) Replay() []LogEntry {
	return dl.data
}