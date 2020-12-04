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

func printMap(mp map[string]string) {
	for k, v := range mp {
		fmt.Printf("%v: %v, ", k, v)
	}
	fmt.Println("")
}

func main() {
	file, err := os.Open(os.Args[1])
	fatal(err)
	defer file.Close()
	types := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validRecords := 0
	scanner := bufio.NewScanner(file)
	recordMap := make(map[string]string, 8)
	i := 0
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n")
		if line == "" {
			// Process the record
			valid := true
			for _, x := range types {
				val, valid := recordMap[x]
				if !valid || val == "" {
					break
				}
			}
			if valid {
				printMap(recordMap)
				validRecords++
			}
			recordMap = make(map[string]string, 8)
			i++
			continue
		}
		for _, rec := range strings.Split(line, " ") {
			x := strings.Split(rec, ":")
			recordMap[string(x[0])] = string(x[1])
		}
	}
	fmt.Println("Challenge 1:", validRecords)
	fmt.Println("Challenge 2:")
}
