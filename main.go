package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// parse the flags
	showLong := flag.Bool("l", false, "Show Long format")
	showAll := flag.Bool("a", false, "Show Hidden files")
	showHelp := flag.Bool("h", false, "Show Help Menu")
	flag.Parse()

	if *showHelp {
		fmt.Println(flag.Usage)
		os.Exit(0)
	}

	// get the arguments
	arguments := flag.Args()

	// default to current dir
	if len(arguments) == 0 {
		arguments = append(arguments, ".")
	}

	// get the current dir entries
	entries, err := os.ReadDir(arguments[0])
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		os.Exit(1)
	}

	// print
	for _, entry := range entries {
		// showAll
		if !*showAll && entry.Name()[0] == '.' && len(entry.Name()) > 0 {
			continue
		}
		// showlong
		if *showLong {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}
			mode := info.Mode()
			modTime := info.ModTime().Format("Jan 01 13:00")
			size := info.Size()
			fmt.Printf("%s\t%s\t%d\t%s\n", entry.Name(), mode, size, modTime)
		}
		fmt.Printf("%s\n", entry.Name())
	}
	os.Exit(0)
}





















