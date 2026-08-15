package day3

import (
	"Aoc2015/lib/set"
	"fmt"
	"io"
	"os"
)

type point struct {
	x, y int
}

func Solution1(f *os.File) {
	in, err := io.ReadAll(f)
	f.Seek(0, io.SeekStart)

	if err != nil {
		fmt.Println("could not read the full file, solution failed")
	}

	h := set.Create[point](50)
	var starting_point = point{0, 0}
	h.Append(starting_point)
	for _, elem := range in {
		switch elem {
		case '>':
			starting_point.x += 1
			h.Append(starting_point)
		case '<':
			starting_point.x -= 1
			h.Append(starting_point)
		case 'v':
			starting_point.y -= 1
			h.Append(starting_point)
		case '^':
			starting_point.y += 1
			h.Append(starting_point)
		}
	}
	fmt.Println("the solution to day3 part1 was:", h.Len())
}

func Solution2(f *os.File) {
	in, err := io.ReadAll(f)
	f.Seek(0, io.SeekStart)

	if err != nil {
		fmt.Println("could not read the full file, solution failed")
	}

	h1 := set.Create[point](50) // santa
	h2 := set.Create[point](50) // robo-santa
	var starting_point1 = point{0, 0}
	var starting_point2 = point{0, 0}
	h1.Append(starting_point1)
	h2.Append(starting_point2)
	santa_turn := true
	for _, elem := range in {
		switch elem {
		case '>':
			if santa_turn {
				starting_point1.x += 1
				h1.Append(starting_point1)
				santa_turn = false
			} else {
				starting_point2.x += 1
				h2.Append(starting_point2)
				santa_turn = true
			}
		case '<':
			if santa_turn {
				starting_point1.x -= 1
				h1.Append(starting_point1)
				santa_turn = false
			} else {
				starting_point2.x -= 1
				h2.Append(starting_point2)
				santa_turn = true
			}
		case 'v':
			if santa_turn {
				starting_point1.y -= 1
				h1.Append(starting_point1)
				santa_turn = false
			} else {
				starting_point2.y -= 1
				h2.Append(starting_point2)
				santa_turn = true
			}
		case '^':
			if santa_turn {
				starting_point1.y += 1
				h1.Append(starting_point1)
				santa_turn = false
			} else {
				starting_point2.y += 1
				h2.Append(starting_point2)
				santa_turn = true
			}
		}
	}
	fmt.Println("the solution to day3 part2 was:", set.Union(h1, h2).Len())
}
