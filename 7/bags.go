package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bagArray map[string]*Bag

// Bags ...
type Bags struct {
	bags bagArray
}

// NewBags ...
func NewBags() *Bags {
	return &Bags{
		bags: make(bagArray),
	}
}

// NewBag ...
func (b *Bags) NewBag(color string) *Bag {
	color = strings.Trim(color, " \n.")

	nb := &Bag{
		color:  color,
		bags:   make(map[string]int),
		parent: b,
	}
	b.bags[color] = nb
	return nb
}

// Items ...
func (b *Bags) Items() bagArray {
	return b.bags
}

// Get a bag from the bags
func (b *Bags) Get(color string) *Bag {
	return b.bags[color]
}

// Bag ...
type Bag struct {
	color  string
	bags   map[string]int
	parent *Bags
	//bags  bagArray
}

func (b *Bag) String() string {
	return b.color
}

// AddBags ...
func (b *Bag) AddBags(bag ...string) {
	for _, bg := range bag {
		bg = strings.Trim(bg, "\n. ")
		vars := strings.Split(bg, " ")
		count, _ := strconv.Atoi(vars[0])
		b.bags[vars[1]+" "+vars[2]] = count
	}
}

// ColorInBags return true if any of the subbags contains the color bag
func (b *Bag) ColorInBags(color string) bool {
	for bg := range b.bags {
		if bg == color {
			return true
		}
		// bg is a string, we need the bag object
		bgObj := b.parent.Get(bg)
		if bgObj.ColorInBags(color) {
			return true
		}
	}
	return false
}

// CountBags returns the number of bags in the bag
func (b *Bag) CountBags() int {
	count := 0
	for color, i := range b.bags {
		count += i * (b.parent.Get(color).CountBags() + 1)
	}
	return count
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	bags := NewBags()
	file, err := os.Open(os.Args[1])
	fatal(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineD := strings.Split(line, " bags contain ")
		// ["shiny gold", "1 dark olive bag, 2 vibrant plum bags.\n"]
		b := bags.NewBag(lineD[0])
		lineD = strings.Split(lineD[1], ", ")
		if lineD[0] != "no other bags." {
			b.AddBags(lineD...)
		}
	}

	part1 := 0
	for _, v := range bags.Items() {
		if v.ColorInBags("shiny gold") {
			part1++
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", bags.Get("shiny gold").CountBags())
}
