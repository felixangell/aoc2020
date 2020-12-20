package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instr struct {
	opcode   string
	argument int64
}

func executeProgram(program []instr) {
	pc := int64(0)
	acc := int64(0)

	pcHistory := make([]bool, len(program))
	inc := func(by int64) {
		pcHistory[pc] = true
		nextPos := pc + by
		if pcHistory[nextPos] {
			panic("we've been here before!")
		}
		pc += by
	}

	for {
		curr := program[pc]
		jumpBy := int64(1)

		switch curr.opcode {
		case "nop":
			break
		case "acc":
			acc += curr.argument
			fmt.Println("@pc", pc, "acc=", acc)
			break
		case "jmp":
			fmt.Println("@pc", pc)
			jumpBy = curr.argument
			break
		}

		inc(jumpBy)
	}
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var program []instr
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		opcode, arg := parts[0], parts[1]

		argument, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			panic(err)
		}

		program = append(program, instr{
			opcode:   opcode,
			argument: argument,
		})
	}

	executeProgram(program)
}
