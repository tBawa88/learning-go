// Dup1 prints the text of each line that appears more than 1 time in the standard input precedded with it's count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//A freq counter to count freq of each line entered through Stdin
	fmt.Println("Press Ctrl + D to end the program")
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	//BTW this is a while loop which will run as long as input.Scan() keeps returning True (meaning there is a new line input in the terminal)
	for input.Scan() {
		count[input.Text()]++ // storing the most recent string line in the map while updating it's freq
	}

	//Now traverse over the count map and print those keys that have value > 1
	for key, value := range count {
		if value > 1 {
			fmt.Printf("Duplicate line %s - %d\n", key, value)
		}
	}

}

// STEPS
// declare a freq counter map[string]int
// create a bufio.Scanner using os.Stdin as a pointer to os.File
// then read from that scanner and put all the lines inside a map
// finally iterate over that map and print all those lines that have freq more than 1
