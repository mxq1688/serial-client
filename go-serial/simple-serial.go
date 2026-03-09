package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Logcat Serial Debugger v1.0")
	fmt.Println("----------------------------")
	
	if len(os.Args) > 1 {
		port := os.Args[1]
		fmt.Printf("Port: %s\n", port)
	} else {
		fmt.Println("Usage: serial-debugger <port>")
		fmt.Println("Example: serial-debugger COM3")
		fmt.Println("Example: serial-debugger /dev/ttyUSB0")
	}
	
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}