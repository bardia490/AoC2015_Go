package day17

import (
	"Aoc2015/lib/utility"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func generateContainersArray(array []int, in string, index int) (intParsingError error) {
	val, err := strconv.Atoi(in)
	if err != nil {
		return err
	}
	array[index] = val
	return nil
}

func part1(array []int) int {
	result := 0

	iterator := utility.GenerateSubsets(array)
	for subSet := range iterator {
		sum := utility.SumSlice(subSet)
		if sum == 150 {
			result += 1
		} else if sum > 150 {
			continue
		}
	}
	return result
}

func part2(array []int) int {
	result := 0
	minimum_containers := 20 // because the minimum of number of containers is definitely going to less than all of them!

	iterator := utility.GenerateSubsets(array)
	for subSet := range iterator {
		sum := utility.SumSlice(subSet)
		if sum == 150 {
			if len(subSet) == minimum_containers {
				result += 1
			} else if len(subSet) <= minimum_containers {
				// reset everything
				result = 1
				minimum_containers = len(subSet)
			}
		} else if sum > 150 {
			continue
		}
	}
	return result
	// i could have cheated and used len(subSet) == 4 for the conditions since i had known the length
	// had to be 4 by printing the previous solutions subsets (part1)
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var array [20]int // this is due the fact that my input has 20 lines, but can be changed accordingly
	for sc.Scan() {
		line := sc.Text()
		err := generateContainersArray(array[:], line, index)
		if err != nil {
			fmt.Printf("something went wrong with parsing int at line: %d, the error: %s", index, err.Error())
			return
		}
		index += 1
	}
	result := part1(array[:])

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day17 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var array [20]int // this is due the fact that my input has 20 lines, but can be changed accordingly
	for sc.Scan() {
		line := sc.Text()
		err := generateContainersArray(array[:], line, index)
		if err != nil {
			fmt.Printf("something went wrong with parsing int at line: %d, the error: %s", index, err.Error())
			return
		}
		index += 1
	}
	result := part2(array[:])

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day17 part 2 is:", result)
}
