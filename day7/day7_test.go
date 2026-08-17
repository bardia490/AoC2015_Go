package day7

import (
	"testing"
)

func TestProcessInstructions(t *testing.T) {
	instruction_notes := [8]string{"x AND y -> d", "x OR y -> e", "x LSHIFT 2 -> f", "123 -> x", "456 -> y", "y RSHIFT 2 -> g", "NOT x -> h", "NOT y -> i"}
	instructions := make(map[string]instruction, 8)
	wires := make(map[string]uint16, 10)

	for _, note := range instruction_notes {
		ParseInstructions(instructions, wires, note)
	}
	ProcessInstructions(instructions, wires)
	answers := map[string]uint16{"d": 72, "e": 507, "f": 492, "g": 114, "h": 65412, "i": 65079, "x": 123, "y": 456}

	for wire, answer := range answers {
		if wires[wire] != answer {
			t.Fatalf("Wanted: %d, got %d, with key: %s", answer, wires[wire], wire)

		}
	}
}
