package main

import (
	"sync"
	"time"
)

type PerformanceMonitor struct {
	mu            sync.RWMutex
	BytesReceived int64
	BytesSent     int64
	PacketsRx     int64
	PacketsTx     int64
	Errors        int64
	StartTime     time.Time
	RxRate        float64
	TxRate        float64
}

func NewPerformanceMonitor() *PerformanceMonitor {
	return &PerformanceMonitor{
		StartTime: time.Now(),
	}
}

func (pm *PerformanceMonitor) UpdateRx(bytes int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.BytesReceived += int64(bytes)
	pm.PacketsRx++
}

func (pm *PerformanceMonitor) UpdateTx(bytes int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.BytesSent += int64(bytes)
	pm.PacketsTx++
}

func (pm *PerformanceMonitor) IncrementErrors() {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.Errors++
}

func (pm *PerformanceMonitor) GetStats() map[string]int64 {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return map[string]int64{
		"bytes_rx":   pm.BytesReceived,
		"bytes_tx":   pm.BytesSent,
		"packets_rx": pm.PacketsRx,
		"packets_tx": pm.PacketsTx,
		"errors":     pm.Errors,
	}
}

func (pm *PerformanceMonitor) Reset() {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.BytesReceived = 0
	pm.BytesSent = 0
	pm.PacketsRx = 0
	pm.PacketsTx = 0
	pm.Errors = 0
	pm.StartTime = time.Now()
}