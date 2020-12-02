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

func isValidPassword(min, max int, letter, password string) bool {
	c := strings.Count(password, letter)
	if c < min || c > max {
		return false
	}
	return true
}

func isValidTobogganPassword(min, max int, letter, password string) bool {
	count := 0
	if string(password[min-1]) == letter {
		count++
	}
	if string(password[max-1]) == letter {
		count++
	}
	return count == 1
}

func main() {
	var validPasswords, validTobogganPasswords int
	file, err := os.Open(os.Args[1])
	fatal(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var min, max int
		var letter, password string
		_, err := fmt.Sscanf(line, "%d-%d %1s: %s\n", &min, &max, &letter, &password)
		fatal(err)
		if isValidPassword(min, max, letter, password) {
			validPasswords++
		}
		if isValidTobogganPassword(min, max, letter, password) {
			validTobogganPasswords++
		}
	}
	fmt.Println("Challenge 1:")
	fmt.Println(validPasswords)
	fmt.Println("Challenge 2:")
	fmt.Println(validTobogganPasswords)
}
