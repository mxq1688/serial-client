package main

import (
	"math"
	"strconv"
)

type WaveformData struct {
	Points  []float64
	Max     float64
	Min     float64
	Average float64
}

func NewWaveformData() *WaveformData {
	return &WaveformData{
		Points: []float64{},
	}
}

func (w *WaveformData) AddPoint(value float64) {
	w.Points = append(w.Points, value)
	if len(w.Points) > 1000 {
		w.Points = w.Points[1:]
	}
	w.updateStats()
}

func (w *WaveformData) updateStats() {
	if len(w.Points) == 0 {
		return
	}
	w.Max = w.Points[0]
	w.Min = w.Points[0]
	sum := 0.0
	for _, p := range w.Points {
		if p > w.Max {
			w.Max = p
		}
		if p < w.Min {
			w.Min = p
		}
		sum += p
	}
	w.Average = sum / float64(len(w.Points))
}

func (w *WaveformData) GetStats() string {
	if len(w.Points) == 0 {
		return "No data"
	}
	return "Max: " + strconv.FormatFloat(w.Max, 'f', 2, 64) +
		" Min: " + strconv.FormatFloat(w.Min, 'f', 2, 64) +
		" Avg: " + strconv.FormatFloat(w.Average, 'f', 2, 64)
}

func (w *WaveformData) GetStdDev() float64 {
	if len(w.Points) < 2 {
		return 0
	}
	variance := 0.0
	for _, p := range w.Points {
		diff := p - w.Average
		variance += diff * diff
	}
	variance /= float64(len(w.Points))
	return math.Sqrt(variance)
}