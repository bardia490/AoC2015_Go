package day23

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	fileContents := "inc a\njio a, +2\ntpl a\ninc a"
	contents := strings.NewReader(fileContents)
	sc := bufio.NewScanner(contents)
	index := 0
	instructions := make([]token, 49)

	for sc.Scan() {
		line := sc.Bytes()
		_ = line
		if err := setInstructions(line, instructions, index); err != nil {
			fmt.Println("problem parsing the instructions:", err.Error())
			return
		}
		index += 1
	}
	a, b := part1(instructions)
	if a != 2 {
		t.Fatalf("a was not correct, a: %d", a)
	}
	if b != 0 {
		t.Fatalf("b was not correct, b: %d", b)
	}
}
