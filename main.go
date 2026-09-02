package main

import (
	"Aoc2015/day1"
	"Aoc2015/day10"
	"Aoc2015/day11"
	"Aoc2015/day12"
	"Aoc2015/day13"
	"Aoc2015/day14"
	"Aoc2015/day15"
	"Aoc2015/day16"
	"Aoc2015/day17"
	"Aoc2015/day18"
	"Aoc2015/day19"
	"Aoc2015/day2"
	"Aoc2015/day20"
	"Aoc2015/day3"
	"Aoc2015/day4"
	"Aoc2015/day5"
	"Aoc2015/day6"
	"Aoc2015/day7"
	"Aoc2015/day8"
	"Aoc2015/day9"
	"fmt"
	"os"
)

func main() {
	var solution string
	fmt.Print("Enter solution number: ")
	fmt.Scan(&solution)

	path := "inputs/day" + solution + ".txt"
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("could not read file:", path)
		return
	}
	defer f.Close()

	// this is just a hack until i add a new interface or something (ngl I'm kinda enjoying this)
	switch solution {
	case "1":
		day1.Solution1(f)
		day1.Solution2(f)
	case "2":
		day2.Solution1(f)
		day2.Solution2(f)
	case "3":
		day3.Solution1(f)
		day3.Solution2(f)
	case "4":
		day4.Solution1(f)
		day4.Solution2(f)
	case "5":
		day5.Solution1(f)
		day5.Solution2(f)
	case "6":
		day6.Solution1(f)
		day6.Solution2(f)
	case "7":
		day7.Solution1(f)
		day7.Solution2(f)
	case "8":
		day8.Solution1(f)
		day8.Solution2(f)
	case "9":
		day9.Solution1(f)
		day9.Solution2(f)
	case "10":
		day10.Solution1(f)
		day10.Solution2(f)
	case "11":
		day11.Solution1(f)
		day11.Solution2(f)
	case "12":
		day12.Solution1(f)
		day12.Solution2(f)
	case "13":
		day13.Solution1(f)
		day13.Solution2(f)
	case "14":
		day14.Solution1(f)
		day14.Solution2(f)
	case "15":
		day15.Solution1(f)
		day15.Solution2(f)
	case "16":
		day16.Solution1(f)
		day16.Solution2(f)
	case "17":
		day17.Solution1(f)
		day17.Solution2(f)
	case "18":
		day18.Solution1(f)
		day18.Solution2(f)
	case "19":
		day19.Solution1(f)
		day19.Solution2(f)
	case "20":
		day20.Solution1(f)
		day20.Solution2(f)
	}
}
