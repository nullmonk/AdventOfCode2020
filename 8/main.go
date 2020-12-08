package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type programError struct {
	typ int
}

func (e *programError) Error() string {
	return fmt.Sprint("Program Error:", e.typ)
}

func (e *programError) Complete() bool {
	if e.typ == 0 {
		return true
	}
	return false
}

func (e *programError) Repeat() bool {
	if e.typ == 1 {
		return true
	}
	return false
}

func ProgramError(typ int) *programError {
	return &programError{
		typ: typ,
	}
}

type State struct {
	pc  int
	acc int
	ran map[int]bool // Instructions that have ran already
}

type Program struct {
	lines      []string
	prevStates []*State
	overRides  int
	State
}

func (p *Program) Reset() {
	p.lines = make([]string, 0)
	p.ran = make(map[int]bool)
	p.overRides = -1
}

func (p *Program) Next() *programError {
	line := p.lines[p.pc]
	instr := strings.Split(line, " ")
	val, _ := strconv.Atoi(instr[1])
	if _, ok := p.ran[p.pc]; ok {
		return ProgramError(1) // Repeated Instruction
	}

	if p.pc == p.overRides {
		p.Printf("Using override\n")
		if instr[0] == "nop" {
			instr[0] = "jmp"
		} else {
			instr[0] = "nop"
		}
	}
	//fmt.Printf("%v: '%v'; %v %v\n", p.pc, line, instr[0], val)
	switch instr[0] {
	case "nop":
		p.Printf("%v %v (saving state)\n", instr[0], val)
		if p.overRides == -1 {
			p.SaveState()
		}
		p.ran[p.pc] = true
		p.pc++
		break
	case "jmp":
		p.Printf("%v %v (saving state)\n", instr[0], val)
		if p.overRides == -1 {
			p.SaveState()
		}
		p.ran[p.pc] = true
		p.pc += val
		break
	case "acc":
		p.Printf("%v %v\n", instr[0], val)
		p.acc += val
		p.ran[p.pc] = true
		p.pc++
		break
	}

	if p.pc >= len(p.lines) {
		return ProgramError(0)
	}
	return ProgramError(-1)
}

func (p *Program) SwapInstruction(i int) {
	line := p.lines[i]
	p.Printf("Swapping inst. ")
	if strings.Contains(line, "jmp") {
		p.lines[i] = strings.Replace(line, "jmp", "nop", 1)
		fmt.Println("jmp->nop")
		return
	}
	if strings.Contains(line, "nop") {
		p.lines[i] = strings.Replace(line, "nop", "jmp", 1)
		fmt.Println("jmp->nop")
	}
}

func (p *Program) SaveState() *State {
	s := &State{
		pc:  p.pc,
		acc: p.acc,
		ran: make(map[int]bool, len(p.ran)),
	}
	for k, v := range p.ran {
		s.ran[k] = v
	}
	p.prevStates = append(p.prevStates, s)
	//p.Printf("Saving state\n")
	return s
}

func (p *Program) Load() {
	i := len(p.prevStates) - 1
	p.Printf("Hopping back a state: %v\n", i)
	if i < 0 {
		log.Fatal("Cannot go back any more")
	}
	s := p.prevStates[i]
	p.prevStates = p.prevStates[:i]
	p.acc = s.acc
	p.pc = s.pc
	p.ran = s.ran
}

func (p *Program) LoadGiven(s *State) {
	p.acc = s.acc
	p.pc = s.pc
	p.ran = s.ran
}

func (p *Program) AddLine(line string) {
	p.lines = append(p.lines, line)
}

func (p *Program) RunP1() {
	for true {
		err := p.Next()
		if err.Complete() || err.Repeat() {
			fmt.Println(p.acc)
			break
		}
	}
}

func (p *Program) Run() {
	for true {
		err := p.Next()
		if err.Complete() {
			p.Printf("Program complete\n")
			break
		}
		if err.Repeat() {
			//fmt.Println("Repeated instruction: Jumping back in time")
			p.Load()
			p.overRides = p.pc
		}
	}
}

func (p *Program) Printf(s string, i ...interface{}) {
	str := fmt.Sprintf(s, i...)
	fmt.Printf("pc: %v acc: %v states: %v - %v", p.pc, p.acc, len(p.prevStates), str)
}

func main() {
	file, err := os.Open(os.Args[1])
	fatal(err)
	defer file.Close()

	program := &Program{}
	program.Reset()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n.")
		program.AddLine(line)
	}

	/*start := program.SaveState()
	program.RunP1()
	program.LoadGiven(start)
	program.prevStates = make([]*State, 0)
	*/
	program.Run()
}
