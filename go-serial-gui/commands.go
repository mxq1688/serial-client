package main

import "strings"

type CommandTemplate struct {
	Name        string
	Command     string
	Description string
	Category    string
}

type CommandManager struct {
	Templates  []CommandTemplate
	History    []string
	MaxHistory int
}

func NewCommandManager() *CommandManager {
	return &CommandManager{
		Templates:  getDefaultTemplates(),
		History:    []string{},
		MaxHistory: 50,
	}
}

func getDefaultTemplates() []CommandTemplate {
	return []CommandTemplate{
		{Name: "AT Test", Command: "AT", Description: "Test connection", Category: "AT"},
		{Name: "Get Version", Command: "AT+VERSION", Description: "Get version", Category: "AT"},
		{Name: "Reset", Command: "AT+RST", Description: "Reset module", Category: "AT"},
		{Name: "LED ON", Command: "LED:1", Description: "Turn LED on", Category: "Arduino"},
		{Name: "LED OFF", Command: "LED:0", Description: "Turn LED off", Category: "Arduino"},
		{Name: "Read Coils", Command: "01 01 00 00 00 10", Description: "Read 16 coils", Category: "Modbus"},
		{Name: "Connect", Command: "CONNECT", Description: "MQTT Connect", Category: "MQTT"},
	}
}

func (cm *CommandManager) AddToHistory(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return
	}
	cm.History = append([]string{cmd}, cm.History...)
	if len(cm.History) > cm.MaxHistory {
		cm.History = cm.History[:cm.MaxHistory]
	}
}

func (cm *CommandManager) GetHistory() []string {
	return cm.History
}

func (cm *CommandManager) ClearHistory() {
	cm.History = []string{}
}

func (cm *CommandManager) ExecuteTemplate(name string) string {
	for _, t := range cm.Templates {
		if t.Name == name {
			return t.Command
		}
	}
	return ""
}