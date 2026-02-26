package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	DefaultPort     string   `json:"default_port\u0022`
	DefaultBaudRate string   `json:"default_baud_rate\u0022`
	WindowWidth     float32  `json:"window_width\u0022`
	WindowHeight    float32  `json:"window_height\u0022`
	RecentPorts     []string `json:"recent_ports\u0022`
	ColorTheme      string   `json:"color_theme\u0022`
	AutoReconnect   bool     `json:"auto_reconnect\u0022`
	LogToFile       bool     `json:"log_to_file\u0022`
	MaxLogLines     int      `json:"max_log_lines\u0022`
}

func NewConfig() *Config {
	return &Config{
		DefaultPort:     "",
		DefaultBaudRate: "115200",
		WindowWidth:     800,
		WindowHeight:    600,
		RecentPorts:     []string{},
		ColorTheme:      "dark",
		AutoReconnect:   true,
		LogToFile:       false,
		MaxLogLines:     1000,
	}
}

func (c *Config) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, c)
}

func (c *Config) Save(filename string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (c *Config) AddRecentPort(port string) {
	for i, p := range c.RecentPorts {
		if p == port {
			c.RecentPorts = append(c.RecentPorts[:i], c.RecentPorts[i+1:]...)
			break
		}
	}
	c.RecentPorts = append([]string{port}, c.RecentPorts...)
	if len(c.RecentPorts) > 5 {
		c.RecentPorts = c.RecentPorts[:5]
	}
}