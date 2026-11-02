package main

import (
	"fmt"
	"os"
)

func main() {
	// get the arguments
	arguments := os.Args[1:]

	// default to current dir
	if len(arguments)==0 {
		arguments = append(arguments, ".")
	}

	// get the current dir entries
	entries, err := os.ReadDir(arguments[0])
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s path doesn't exists", arguments[0])
		} else {
			println(err)
		}
		os.Exit(0)
	}
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
