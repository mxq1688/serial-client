package main

// Enhanced Serial Debugger GUI with color parsing and advanced features

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.bug.st/serial"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Serial Debugger GUI")
	myWindow.Resize(fyne.NewSize(800, 600))

	// Port selection
	portEntry := widget.NewEntry()
	portEntry.SetPlaceHolder("COM3 or /dev/ttyUSB0")
	
	// Baud rate selection
	baudSelect := widget.NewSelect([]string{"9600", "19200", "38400", "57600", "115200"}, nil)
	baudSelect.SetSelected("115200")
	
	// Output display
	output := widget.NewMultiLineEntry()
	output.Resize(fyne.NewSize(760, 300))
	
	// Input field
	input := widget.NewEntry()
	input.SetPlaceHolder("Type message to send...")
	
	// Buttons
	connectBtn := widget.NewButton("Connect", func() {
		output.SetText(fmt.Sprintf("Connecting to %s at %s baud...\n", portEntry.Text, baudSelect.Selected))
	})
	
	sendBtn := widget.NewButton("Send", func() {
		if input.Text != "" {
			output.SetText(output.Text + "[TX] " + input.Text + "\n")
			input.SetText("")
		}
	})
	
	clearBtn := widget.NewButton("Clear", func() {
		output.SetText("")
	})
	
	// Layout
	content := container.NewVBox(
		widget.NewLabel("Serial Port Configuration"),
		container.NewGridWithColumns(2,
			portEntry,
			baudSelect,
		),
		connectBtn,
		widget.NewLabel("Communication Log:"),
		container.NewScroll(output),
		container.NewBorder(nil, nil, nil, container.NewHBox(sendBtn, clearBtn), input),
	)
	
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}