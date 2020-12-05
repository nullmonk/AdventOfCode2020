package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getRowColID(line string) int {
	// Convert the line to an int using 'F' and 'L' as a binary 0 and 'B' and 'R' as a binary 1
	n := 0
	for _, c := range line {
		n <<= 1
		if c == 'B' || c == 'R' {
			n |= 1
		}
	}
	// Return the id
	// Row is everything except the last 3 bits (n>>3)
	// Col is only teh last three bits (n & 7)
	return (n>>3)*8 + (n & 7)
}
func main() {
	file, err := os.Open(os.Args[1])
	fatal(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	max := 0
	var ids = make([]bool, 1028)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n")
		id := getRowColID(line)
		if id > max {
			max = id
		}

		ids[id] = true
	}

	myID := 0
	for i, exists := range ids {
		// Dont care about the first one
		if i == 0 {
			continue
		}
		// If both seats next to mine are filled, and this seat is not, then its mine
		if !exists && ids[i-1] && ids[i+1] {
			myID = i
		}
	}
	fmt.Printf("Part 1: %v\nPart 2: %v\n", max, myID)
}
