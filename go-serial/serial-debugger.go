package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("=====================================")
	fmt.Println("   Logcat Serial Debugger v1.0")
	fmt.Println("=====================================")
	fmt.Println()

	var port string
	var baudRate int

	// Get serial port
	fmt.Print("Enter serial port (e.g.