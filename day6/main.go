package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func UpdateLights(lights []bool, line string) {
	instructions := strings.Split(line, " ")

	if instructions[0] == "turn" {
		comma_index_begining := strings.IndexByte(instructions[2], ',')
		comma_index_end := strings.IndexByte(instructions[4], ',')
		x1, err_x := strconv.Atoi(instructions[2][:comma_index_begining])
		y1, err_y := strconv.Atoi(instructions[2][comma_index_begining+1:])

		if err_x != nil {
			println("there was a problem parsing the begining x:", instructions[2][:comma_index_begining], " at line:", line)
			println("canceling operations")
			return
		}

		if err_y != nil {
			println("there was a problem parsing the begining y:", instructions[2][comma_index_begining+1:], " at line:", line)
			println("canceling operations")
			return
		}

		x2, err_x := strconv.Atoi(instructions[4][:comma_index_end])
		y2, err_y := strconv.Atoi(instructions[4][comma_index_end+1:])

		if err_x != nil {
			println("there was a problem parsing the ending x:", instructions[2][:comma_index_end], " at line:", line)
			println("canceling operations")
			return
		}

		if err_y != nil {
			println("there was a problem parsing the end y:", instructions[2][comma_index_end+1:], " at line:", line)
			println("canceling operations")
			return
		}

		new_instruction := false
		if instructions[1] == "on" {
			new_instruction = true
		}

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				lights[1000*x+y] = new_instruction
			}
		}
	} else { // toggle case
		comma_index_begining := strings.IndexByte(instructions[1], ',')
		x1, err_x := strconv.Atoi(instructions[1][:comma_index_begining])
		y1, err_y := strconv.Atoi(instructions[1][comma_index_begining+1:])

		if err_x != nil {
			println("there was a problem parsing the begining x:", instructions[2][:comma_index_begining], " at line:", line)
			println("canceling operations")
			return
		}

		if err_y != nil {
			println("there was a problem parsing the begining y:", instructions[2][comma_index_begining+1:], " at line:", line)
			println("canceling operations")
			return
		}

		comma_index_end := strings.IndexByte(instructions[3], ',')
		x2, err_x := strconv.Atoi(instructions[3][:comma_index_end])
		y2, err_y := strconv.Atoi(instructions[3][comma_index_end+1:])

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				lights[1000*x+y] = !lights[1000*x+y]
			}
		}
	}
}

func CountLitLights(lights []bool) uint32 {
	var result uint32

	for _, val := range lights {
		if val {
			result += 1
		}
	}
	return result
}

func UpdateLights2(lights []uint32, line string) {
	instructions := strings.Split(line, " ")

	if instructions[0] == "turn" {
		comma_index_begining := strings.IndexByte(instructions[2], ',')
		comma_index_end := strings.IndexByte(instructions[4], ',')
		x1, err_x := strconv.Atoi(instructions[2][:comma_index_begining])
		y1, err_y := strconv.Atoi(instructions[2][comma_index_begining+1:])

		if err_x != nil {
			println("there was a problem parsing the begining x:", instructions[2][:comma_index_begining], " at line:", line)
			println("canceling operations")
			return
		}

		if err_y != nil {
			println("there was a problem parsing the begining y:", instructions[2][comma_index_begining+1:], " at line:", line)
			println("canceling operations")
			return
		}

		x2, err_x := strconv.Atoi(instructions[4][:comma_index_end])
		y2, err_y := strconv.Atoi(instructions[4][comma_index_end+1:])

		if err_x != nil {
			println("there was a problem parsing the ending x:", instructions[2][:comma_index_end], " at line:", line)
			println("canceling operations")
			return
		}

		if err_y != nil {
			println("there was a problem parsing the end y:", instructions[2][comma_index_end+1:], " at line:", line)
			println("canceling operations")
			return
		}

		next_instruction := -1
		if instructions[1] == "on" {
			next_instruction = 1
		}
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if next_instruction > 0 || lights[1000*x+y] > 0 {
					lights[1000*x+y] += uint32(next_instruction)
				}
			}
		}
	} else { // toggle case
		comma_index_begining := strings.IndexByte(instructions[1], ',')
		x1, err_x := strconv.Atoi(instructions[1][:comma_index_begining])
		y1, err_y := strconv.Atoi(instructions[1][comma_index_begining+1:])

		if err_x != nil {
			println("there was a problem parsing the begining x:", instructions[2][:comma_index_begining], " at line:", line)
			println("canceling operations")
			return
		}

		if err_y != nil {
			println("there was a problem parsing the begining y:", instructions[2][comma_index_begining+1:], " at line:", line)
			println("canceling operations")
			return
		}

		comma_index_end := strings.IndexByte(instructions[3], ',')
		x2, err_x := strconv.Atoi(instructions[3][:comma_index_end])
		y2, err_y := strconv.Atoi(instructions[3][comma_index_end+1:])

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				lights[1000*x+y] += 2
			}
		}
	}
}

func CountLitLights2(lights []uint32) uint32 {
	var result uint32

	for _, val := range lights {
		result += val
	}
	return result
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	lights := make([]bool, 1000*1000)

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		UpdateLights(lights, line)
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	fmt.Println("the solution to day6 part 1 is:", CountLitLights(lights))
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	lights := make([]uint32, 1000*1000)

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		UpdateLights2(lights, line)
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	fmt.Println("the solution to day6 part 2 is:", CountLitLights2(lights))
}
