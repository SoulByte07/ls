package main

import (
	"fmt"
	"os"
	"flag"
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

	// flag: long format 
	var showLong bool
	flag.BoolVar(&showLong, "l", false, "Shows Long format")



	for _, entry := range entries {
		fmt.Printf("%s \t", entry)
		info, err := entry.Info()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		mode := info.Mode()
		size := info.Size()
		modTime := info.ModTime()
		fmt.Printf("%s \t %d \t %s \n", mode, size, modTime)
	}
}
























