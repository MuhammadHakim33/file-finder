package main

import (
	"file-finder/internal"
	"flag"
	"fmt"
	"log"
)

const MAXOUTPUTFILE int = 100

func main() {
	// flags
	dir := flag.String("d", "", "Path to the directory to scan")
	keyword := flag.String("s", "", "Keyword to search for within files")
	ext := flag.String("e", "", "Filter files by extension, e.g., txt, png, pdf")
	flag.Parse()

	log.Println("Start")

	// initialize finder
	f, err := internal.New(*keyword, *dir, *ext)
	if err != nil {
		log.Println("Initialize failed:", err)
		return
	}

	// start finder
	result, err := f.Find()
	if err != nil {
		log.Println("Finder failed:", err)
		return
	}

	// show result
	for i, p := range result {
		if i > MAXOUTPUTFILE {
			break
		}

		fmt.Println("+ ", *p)
	}

	// response if the result exceeds the max
	exceed_response := ""
	if len(result) > MAXOUTPUTFILE {
		exceed_response = "Try a more specific keyword."
	}

	log.Printf("Found %d matching files. %s", len(result), exceed_response)
}
