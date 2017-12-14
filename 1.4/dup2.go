// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Make a map of string to map of counts per filename
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, matches := range counts {
		total := 0
		names := make([]string, 0)
		for name, count := range matches {
			total += count
			names = append(names, name)
		}

		if total > 1 {
			fmt.Printf("%d\t'%s'\n", total, line)
			fmt.Printf("%v\n", names)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for {
		if ok := input.Scan(); !ok {
			// If we got an actual error (not io.EOF), puke it out.
			if err := input.Err(); err != nil {
				log.Fatalf("I/O error reading '%s': %s",
					f.Name(), err)
			}

			return
		}

		text := input.Text()
		if _, ok := counts[text]; !ok {
			counts[text] = map[string]int{f.Name(): 1}
		} else {
			counts[input.Text()][f.Name()]++
		}
	}
}

//!-
