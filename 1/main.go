package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const chal = 2020

func main() {
	// Store all the numbers
	numbers := make(map[int]bool, 500)

	file, err := os.Open(os.Args[1])
	fatal(err)
	defer file.Close()
	fmt.Println("Challenge 1:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		numbers[i] = true
		if _, ok := numbers[chal-i]; ok {
			v := chal - i
			// Challenge 1 answer
			fmt.Printf("x = %v, y = %v\nx * y = %v\n", i, v, i*v)
		}
	}

	// This solves challenge 2
	// We need to iterate back over the items again
	fmt.Println("\nChallenge 2:")
	for i := range numbers {
		for j := range numbers {
			k := chal - i - j
			if _, ok := numbers[k]; ok {
				fmt.Printf("x = %v, y = %v, x = %v\nx * y * z = %v\n", i, j, k, i*j*k)
				return
			}
		}
	}
}
