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

func main() {
	file, err := os.Open(os.Args[1])
	fatal(err)
	defer file.Close()
	fmt.Println("Challenge 1:")
	scanner := bufio.NewScanner(file)

	var route2, route1, route3, route4, route5 int
	var r1, r2, r3, r4, r5 int // Keep track of the index for each route
	r5Toggle := false
	activated := false
	cont := 0
	for scanner.Scan() {
		if !activated {
			_ = strings.Trim(scanner.Text(), " \n")
			//fmt.Printf("%v: %v\n", cont, line)
			cont++
			activated = true
			continue
		}
		line := strings.Trim(scanner.Text(), " \n")
		//fmt.Printf("%v: %v\n", cont, line)
		cont++

		r1 = (r1 + 1) % len(line)
		if line[r1] == '#' {
			route1++
		}
		r2 = (r2 + 3) % len(line)
		if line[r2] == '#' {
			route2++
		}
		r3 = (r3 + 5) % len(line)
		if line[r3] == '#' {
			route3++
		}
		r4 = (r4 + 7) % len(line)
		if line[r4] == '#' {
			route4++
		}
		fmt.Println(line)

		if r5Toggle {
			r5 = (r5 + 1) % len(line)
			fmt.Printf("%*v\n", r5+1, "^")
			if line[r5] == '#' {
				route5++
			}
		}
		r5Toggle = !r5Toggle
	}
	fmt.Println("Part 1:", route2)
	fmt.Printf("Part 2: %v, %v, %v, %v, %v = %v\n", route1, route2, route3, route4, route5, route1*route2*route3*route4*route5)
}
