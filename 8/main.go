package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// State ...
type State struct {
	pc  int
	acc int
	ran map[int]bool // Instructions that have run already
}

// Program ...
type Program struct {
	lines      []string
	prevStates []*State
	override   int // Override the instruction at this location
	State
}

// Init the program
func (p *Program) Init(filename string) {
	//p.lines = make([]string, 0)
	p.ran = make(map[int]bool)
	p.override = -1 // Very important

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.lines = strings.Split(string(data), "\n")
	for p.lines[len(p.lines)-1] == "" {
		p.lines = p.lines[:len(p.lines)-1]
	}
	//fmt.Printf("Loaded %v lines from %v\n", len(p.lines), filename)
}

func (p *Program) next() int {
	line := p.lines[p.pc]
	instr := strings.Split(line, " ")
	val, _ := strconv.Atoi(instr[1])
	if _, ok := p.ran[p.pc]; ok {
		return 1 // 1 indicates repeat
	}

	ovrstr := ""
	if p.pc == p.override {
		if instr[0] == "nop" {
			instr[0] = "jmp"
		} else {
			instr[0] = "nop"
		}
		ovrstr = "*"
	}
	switch instr[0] {
	case "nop":
		p.printf("%v%v %v\n", instr[0], ovrstr, val)
		if p.override == -1 {
			p.saveState()
		}
		p.ran[p.pc] = true
		p.pc++
		break
	case "jmp":
		p.printf("%v%v %v\n", ovrstr, instr[0], val)
		if p.override == -1 {
			p.saveState()
		}

		p.ran[p.pc] = true
		p.pc += val
		break
	case "acc":
		p.printf("%v %v\n", instr[0], val)
		p.acc += val
		p.ran[p.pc] = true
		p.pc++
		break
	}

	if p.pc >= len(p.lines) {
		return 0 // 0 indicated program has completed
	}
	return -1 // nO status currently. keep processing
}

func (p *Program) saveState() *State {
	s := &State{
		pc:  p.pc,
		acc: p.acc,
		ran: make(map[int]bool, len(p.ran)),
	}
	for k, v := range p.ran {
		s.ran[k] = v
	}
	p.prevStates = append(p.prevStates, s)
	p.printf("Saving state #%v\n", len(p.prevStates))
	return s
}

func (p *Program) load() {
	i := len(p.prevStates) - 1
	p.printf("revert to state: #%v\n", i)
	s := p.prevStates[i]
	p.prevStates = p.prevStates[:i]
	p.acc = s.acc
	p.pc = s.pc
	p.ran = s.ran
}

// Run the program and print the outputs
func (p *Program) Run() {
	p1Solved := false
	for true {
		ret := p.next()
		if ret == 0 {
			p.printf("Program complete\n")
			fmt.Printf("Part 2: %v\n", p.acc)
			break
		}
		if ret == 1 {
			if !p1Solved {
				fmt.Printf("Part 1: %v\n", p.acc)
				p1Solved = true
			}
			p.load()
			p.override = p.pc
		}
	}
}

func (p *Program) printf(s string, i ...interface{}) {
	/*
		str := fmt.Sprintf(s, i...)
		fmt.Printf("pc: %v acc: %v states: %v - %v", p.pc, p.acc, len(p.prevStates), str)
	*/
}

func main() {
	program := &Program{}
	program.Init(os.Args[1])
	program.Run()
}
