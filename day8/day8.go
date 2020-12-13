package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	jmp = "jmp"
	nop = "nop"
	acc = "acc"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	out, _ := runProgram(lines)
	fmt.Println(out)

	newProgram := make([]string, len(lines))

	for i := range lines {
		copy(newProgram, lines)

		if strings.HasPrefix(newProgram[i], jmp) {
			newProgram[i] = strings.Replace(newProgram[i], jmp, nop, 1)
		} else if strings.HasPrefix(newProgram[i], nop) {
			newProgram[i] = strings.Replace(newProgram[i], nop, jmp, 1)
		} else {
			continue
		}

		out, crash := runProgram(newProgram)
		if crash == false {
			fmt.Println(out)
			break
		}
	}
}

func runProgram(program []string) (int, bool) {
	infiniteLoop := false
	visitedLines := map[int]bool{}
	cursor := 1
	accumulator := 0

	for {
		if visitedLines[cursor] || cursor == len(program) {
			infiniteLoop = visitedLines[cursor]
			break
		}

		visitedLines[cursor] = true
		lineSplit := strings.Split(program[cursor], " ")

		val, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			panic("invalid argument")
		}

		switch lineSplit[0] {
		case nop:
			cursor += 1
		case acc:
			accumulator += val
			cursor += 1
		case jmp:
			cursor += val
		}
	}
	return accumulator, infiniteLoop
}
