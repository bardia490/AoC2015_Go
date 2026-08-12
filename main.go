package main

import (
	"fmt"
	"local/Aoc2015/day1"
	"os"
)

func main() {
	path := "inputs/day1.txt"
	fileContents, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("could not read file:", path)
		return
	}
	fmt.Println("the day solution is:", day1.Solution1(fileContents))
	day1_solution2, err := day1.Solution2(fileContents)

	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return
	}
	fmt.Println("the day solution is:", day1_solution2)
}
