package day7

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	first     string
	second    string
	operation string
}

// it will return the final converted value
// Note: this should only be used when you are sure the number can be
// converted to uint16 otherwise it will panic
func convertStringToUint16(num string) uint16 {
	temp, err := strconv.Atoi(num)
	if err != nil {
		panic(fmt.Sprintf("there was a error in integer detection, integer: %s aborting operation", num))
	}
	return uint16(temp)
}

func ParseInstructions(instructions map[string]instruction, wires map[string]uint16, in string) {
	splitted_string := strings.Split(in, " ")
	splitted_string_len := len(splitted_string)

	switch splitted_string_len {
	case 3: // for [a -> b] cases
		if splitted_string[0][0] >= '0' && splitted_string[0][0] <= '9' { // if the first one is a number we dont need to add it to the table
			wires[splitted_string[2]] = convertStringToUint16(splitted_string[0])
		} else { // if the first one wasn't number
			instructions[splitted_string[2]] = instruction{first: splitted_string[0], second: "", operation: ""}
		}
	case 4: // only for [NOT A -> B] cases
		instructions[splitted_string[3]] = instruction{first: splitted_string[1], second: "", operation: splitted_string[0]}
	case 5: // for [A op B -> C] cases
		instructions[splitted_string[4]] = instruction{first: splitted_string[0], second: splitted_string[2], operation: splitted_string[1]}
	}
}

// process an instruction
// final is the value at the right hand side
func ProcessInstruction(instructions map[string]instruction, wires map[string]uint16, in instruction, final string) {
	switch in.operation {
	case "": // case of [A -> B], in.first == A
		val := DetermineWire(instructions, wires, in.first)
		wires[final] = val
	case "NOT": // only for [NOT A -> B] cases
		val := DetermineWire(instructions, wires, in.first)
		wires[final] = ^val
	case "RSHIFT": // only for [NOT A -> B] cases
		var val_1 uint16 = 0
		var val_2 uint16 = 0
		if in.first[0] >= '0' && in.first[0] <= '9' { // if the first one is a number we dont need to add it to the table
			val_1 = convertStringToUint16(in.first)
		} else {
			val_1 = DetermineWire(instructions, wires, in.first)
		}
		if in.second[0] >= '0' && in.second[0] <= '9' { // if the first one is a number we dont need to add it to the table
			val_2 = convertStringToUint16(in.second)
		} else {
			val_2 = DetermineWire(instructions, wires, in.second)
		}
		wires[final] = val_1 >> val_2
	case "LSHIFT": // for [A op B -> C] cases
		var val_1 uint16 = 0
		var val_2 uint16 = 0
		if in.first[0] >= '0' && in.first[0] <= '9' { // if the first one is a number we dont need to add it to the table
			val_1 = convertStringToUint16(in.first)
		} else {
			val_1 = DetermineWire(instructions, wires, in.first)
		}
		if in.second[0] >= '0' && in.second[0] <= '9' { // if the first one is a number we dont need to add it to the table
			val_2 = convertStringToUint16(in.second)
		} else {
			val_2 = DetermineWire(instructions, wires, in.second)
		}
		wires[final] = val_1 << val_2
	case "AND": // for [A op B -> C] cases
		var val_1 uint16 = 0
		var val_2 uint16 = 0
		if in.first[0] >= '0' && in.first[0] <= '9' { // if the first one is a number we dont need to add it to the table
			val_1 = convertStringToUint16(in.first)
		} else {
			val_1 = DetermineWire(instructions, wires, in.first)
		}
		if in.second[0] >= '0' && in.second[0] <= '9' { // if the first one is a number we dont need to add it to the table
			val_2 = convertStringToUint16(in.second)
		} else {
			val_2 = DetermineWire(instructions, wires, in.second)
		}
		wires[final] = val_1 & val_2
	case "OR": // for [A op B -> C] cases
		var val_1 uint16 = 0
		var val_2 uint16 = 0
		if in.first[0] >= '0' && in.first[0] <= '9' { // if the first one is a number we dont need to add it to the table
			val_1 = convertStringToUint16(in.first)
		} else {
			val_1 = DetermineWire(instructions, wires, in.first)
		}
		if in.second[0] >= '0' && in.second[0] <= '9' { // if the first one is a number we dont need to add it to the table
			val_2 = convertStringToUint16(in.second)
		} else {
			val_2 = DetermineWire(instructions, wires, in.second)
		}
		wires[final] = val_1 | val_2
	}
	delete(instructions, final)
}

func DetermineWire(instructions map[string]instruction, wires map[string]uint16, wire string) uint16 {
	val, ok := wires[wire]

	if !ok { // the wire wasn't found yet
		new_instruction := instructions[wire]
		ProcessInstruction(instructions, wires, new_instruction, wire)
	}
	val, _ = wires[wire]
	return val
}

func ProcessInstructions(instructions map[string]instruction, wires map[string]uint16) {
	for wire, ins := range instructions {
		ProcessInstruction(instructions, wires, ins, wire)
	}
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	instructions := make(map[string]instruction, 335)
	wires := make(map[string]uint16, 100)

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		ParseInstructions(instructions, wires, line)
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	ProcessInstructions(instructions, wires)
	fmt.Println("the solution to day7 part 1 is:", wires["a"])
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	instructions := make(map[string]instruction, 335)
	wires := make(map[string]uint16, 100)

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		ParseInstructions(instructions, wires, line)
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}
	delete(instructions, "b")
	wires["b"] = 46065

	ProcessInstructions(instructions, wires)
	fmt.Println("the solution to day7 part 2 is:", wires["a"])
}
