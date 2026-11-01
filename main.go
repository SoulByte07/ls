package main

import (
	"os"
)

func main() {
	currentDir, err := os.ReadDir(".")
	if err!=nil {
		print("Failed to Read current Dir")
	}
	print(currentDir)
}
