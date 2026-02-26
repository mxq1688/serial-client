package main

import (
	"fmt"
	"strings"
)

type ProtocolParser struct {
	protocolType string
}

func NewProtocolParser(pType string) *ProtocolParser {
	return &ProtocolParser{protocolType: pType}
}

func (p *ProtocolParser) Parse(data []byte) string {
	switch p.protocolType {
	case "modbus":
		return p.parseModbus(data)
	case "mqtt":
		return p.parseMQTT(data)
	default:
		return fmt.Sprintf("Raw: %X", data)
	}
}

func (p *ProtocolParser) parseModbus(data []byte) string {
	if len(data) < 4 {
		return "Invalid Modbus frame"
	}
	var result strings.Builder
	result.WriteString(fmt.Sprintf("Modbus RTU:\n"))
	result.WriteString(fmt.Sprintf("  Slave: %d\n", data[0]))
	result.WriteString(fmt.Sprintf("  Function: %d\n", data[1]))
	return result.String()
}

func (p *ProtocolParser) parseMQTT(data []byte) string {
	if len(data) < 2 {
		return "Invalid MQTT frame"
	}
	packetType := (data[0] >> 4) & 0x0F
	return fmt.Sprintf("MQTT Packet Type: %d\n", packetType)
}