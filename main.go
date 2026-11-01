package main

import (
	"os"
	"fmt"
)

func main() {
	arguments := os.Args[1:]

	entries, err := os.ReadDir(arguments[-1])
	if err != nil {
		println("Failed to Read current Dir")
	}
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
