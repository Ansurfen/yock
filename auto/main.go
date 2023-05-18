package main

import (
	"fmt"
	"os"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Failed to get executable path: %v\n", err)
		return
	}
	fmt.Printf("Executable path: %s\n", exePath)
}
