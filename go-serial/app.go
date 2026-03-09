package main

import (
	"context"
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetSerialPorts() []string {
	return []string{"/dev/ttyUSB0", "/dev/ttyUSB1", "COM1", "COM2"}
}

func (a *App) Connect(port string, baudRate int) string {
	return fmt.Sprintf("Connected to %s at %d baud", port, baudRate)
}

func (a *App) Disconnect() string {
	return "Disconnected"
}

func (a *App) SendCommand(cmd string) string {
	return fmt.Sprintf("Sent: %s", cmd)
}

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "Logcat Serial Debugger - Go",
		Width:  1200,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 30, G: 30, B: 30, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}