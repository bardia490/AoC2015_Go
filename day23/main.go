package day23

import (
	"Aoc2015/lib/utility"
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type instruction int

const (
	hlf instruction = iota
	tpl
	inc
	jmp
	jie
	jio
)

type token struct {
	inst     instruction
	value    int
	variable string // for instructions such as inc or jio or jie
}

func setInstructions(in []byte, instructions []token, index int) (parsing_error error) {
	contents := bytes.Split(in, []byte(" "))

	if bytes.Equal(contents[0], []byte("hlf")) {
		instructions[index].inst = hlf
		instructions[index].variable = string(contents[1])
	} else if bytes.Equal(contents[0], []byte("tpl")) {
		instructions[index].inst = tpl
		instructions[index].variable = string(contents[1])
	} else if bytes.Equal(contents[0], []byte("inc")) {
		instructions[index].inst = inc
		instructions[index].variable = string(contents[1])
	} else if bytes.Equal(contents[0], []byte("jmp")) {
		instructions[index].inst = jmp
		val, err := utility.BytesToIntSigned(contents[1])
		if err != nil {
			return err
		}
		instructions[index].value = val
	} else if bytes.Equal(contents[0], []byte("jie")) {
		instructions[index].inst = jie
		variable := contents[1]
		instructions[index].variable = string(variable[:len(variable)-1]) // because of the comma
		val, err := utility.BytesToIntSigned(contents[2])
		if err != nil {
			return err
		}
		instructions[index].value = val
	} else if bytes.Equal(contents[0], []byte("jio")) {
		instructions[index].inst = jio
		variable := contents[1]
		instructions[index].variable = string(variable[:len(variable)-1]) // because of the comma
		val, err := utility.BytesToIntSigned(contents[2])
		if err != nil {
			return err
		}
		instructions[index].value = val
	}
	return nil
}

func runInstruction(instruction token, a, b, instruction_number *int) {
	switch instruction.inst {
	case hlf:
		if instruction.variable == "a" {
			*a /= 2
		} else {
			*b /= 2
		}
		*instruction_number += 1
	case tpl:
		if instruction.variable == "a" {
			*a *= 3
		} else {
			*b *= 3
		}
		*instruction_number += 1
	case inc:
		if instruction.variable == "a" {
			*a += 1
		} else {
			*b += 1
		}
		*instruction_number += 1
	case jmp:
		*instruction_number += instruction.value
	case jie: //even
		if instruction.variable == "a" && *a%2 == 0 {
			*instruction_number += instruction.value
		} else if instruction.variable == "b" && *b%2 == 0 {
			*instruction_number += instruction.value
		} else {
			*instruction_number += 1
		}
	case jio: // one
		if instruction.variable == "a" && *a == 1 {
			*instruction_number += instruction.value
		} else if instruction.variable == "b" && *b == 1 {
			*instruction_number += instruction.value
		} else {
			*instruction_number += 1
		}
	}
}

func part1(instructions []token) (a, b int) {
	len_instructions := len(instructions)
	instruction_number := 0
	a = 0
	b = 0

	for instruction_number < len_instructions {
		runInstruction(instructions[instruction_number], &a, &b, &instruction_number)
	}
	return a, b
}

func part2(instructions []token) (a, b int) {
	len_instructions := len(instructions)
	instruction_number := 0
	a = 1
	b = 0

	for instruction_number < len_instructions {
		runInstruction(instructions[instruction_number], &a, &b, &instruction_number)
	}
	return a, b
}

var instructions []token = make([]token, 49)

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	index := 0

	for sc.Scan() {
		line := sc.Bytes()
		_ = line
		if err := setInstructions(line, instructions, index); err != nil {
			fmt.Println("problem parsing the instructions:", err.Error())
			return
		}
		index += 1
	}
	_, result := part1(instructions)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day23 part 1 is:", result)
}

func Solution2(f *os.File) {
	// we don't even need to parse the file again since the boss is a global variable
	_, result := part2(instructions)

	fmt.Println("the solution to day23 part 2 is:", result)
}
