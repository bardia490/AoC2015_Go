package main

import (
	"fmt"
	"local/Aoc2015/day1"
	"local/Aoc2015/day2"
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

	// this is just a hack until i add a new interface or something
	switch solution {
	case "1":
		day1.Solution1(f)
		day1.Solution2(f)
	case "2":
		day2.Solution1(f)
		day2.Solution2(f)
	}
}
