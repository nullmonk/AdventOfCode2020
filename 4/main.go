package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// return function to make sure a string is between min and mx
func between(s string, min, max int) bool {
	i, _ := strconv.Atoi(s)
	if i < min || i > max {
		return false
	}
	return true
}

func validateHgt(s string) bool {
	if strings.HasSuffix(s, "in") {
		s := strings.TrimRight(s, "in")
		return between(s, 59, 76)
	}
	if strings.HasSuffix(s, "cm") {
		s := strings.TrimRight(s, "cm")
		return between(s, 150, 193)
	}
	return false
}

func printMap(mp map[string]string) {
	for k, v := range mp {
		fmt.Printf("%v: %v, ", k, v)
	}
	fmt.Println("")
}

func strToMap(record string) map[string]string {
	record = strings.Replace(strings.Trim(record, "\n"), "\n", " ", -1)
	rcd := make(map[string]string, 8)
	for _, r := range strings.Split(record, " ") {
		rVals := strings.Split(r, ":")
		rcd[rVals[0]] = rVals[1]
	}
	return rcd
}
func isValidRecord(types map[string]checkFunc, record map[string]string, strict bool) bool {
	for n, fn := range types {
		if val, ok := record[n]; ok {
			if strict && !fn(val) {
				//fmt.Printf("Invalid option: %v:%v\n", n, val)
				return false
			}
			continue
		}
		return false
	}
	return true
}

type checkFunc func(string) bool

func main() {
	types := map[string]checkFunc{
		// Check the byr
		"byr": func(s string) bool { return between(s, 1920, 2002) },
		"iyr": func(s string) bool { return between(s, 2010, 2020) },
		"eyr": func(s string) bool { return between(s, 2020, 2030) },
		"hgt": validateHgt,
		"hcl": regexp.MustCompile("#[0-9a-f]{6}").MatchString,
		"ecl": regexp.MustCompile("(amb|blu|brn|gry|grn|hzl|oth)").MatchString,
		"pid": regexp.MustCompile("^\\d{9}$").MatchString,
		//"cid": func(s string) bool { return true },
	}
	validRecords := 0
	validRecords2 := 0

	//recordMap := make(map[string]string, 8)
	data, err := ioutil.ReadFile(os.Args[1])
	fatal(err)
	for _, v := range strings.Split(string(data), "\n\n") {
		rcd := strToMap(v)
		if isValidRecord(types, rcd, false) {
			validRecords++
			if isValidRecord(types, rcd, true) {
				validRecords2++
				printMap(rcd)

			}
		}
	}
	fmt.Println("Challenge 1:", validRecords)
	fmt.Println("Challenge 2:", validRecords2)
}
