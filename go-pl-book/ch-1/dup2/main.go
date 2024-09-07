// Dup2 can read from the standard input as well as any named files entered as arguments while running the program
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	count := make(map[string]int)
	fileNames := os.Args[1:]

	if len(fileNames) == 0 {
		// call a function that basically does the same thing as Dup1
		fmt.Println("Press Ctrl + D to exit")
		countLines(os.Stdin, count)
	} else {
		for _, name := range fileNames {
			// open each of those files and read their contents
			file, err := os.Open(name)
			if err != nil {
				log.Fatal("Error opening the file ", name)
			}

			// Send that file to the countLines() function
			countLines(file, count)
			file.Close()
		}
	}
}

func countLines(r *os.File, count map[string]int) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		count[input.Text()]++
	}

	for line, freq := range count {
		if freq > 1 {
			fmt.Printf("Line = %s, Freq = %d\n", line, freq)
		}
	}
}
