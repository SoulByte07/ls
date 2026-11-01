package main

import (
	"os"
	"fmt"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		println("Failed to Read current Dir")
	}
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
