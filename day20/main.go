package day20

import (
	"Aoc2015/lib/utility"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func calcHousePresents(house int) int {
	result := 0
	for divisor := range utility.GenerateNumberDivisors(house) {
		result += divisor
	}
	return result
}

func part1(target int) (house_number int) {
	target = target / 10
	initial_guess := 1000

	for {
		presents := calcHousePresents(initial_guess)
		if presents >= target {
			break
		}
		initial_guess += 20
	}

	lowest_house_number := initial_guess

	// probe the last 50 houses to find the lowest one (just in case there is something lower that we missed in the above loop)
	for guess := initial_guess; guess > initial_guess-50; guess-- {
		presents := calcHousePresents(guess)
		if presents >= target && lowest_house_number > guess {
			lowest_house_number = guess
		}
	}

	return lowest_house_number
}

func calcHousePresents2(house int) int {
	result := 0
	for divisor := range utility.GenerateNumberDivisors(house) {
		if house <= divisor*50 {
			result += divisor
		}
	}
	return result
}

func part2(target int) (house_number int) {
	target += target % 11
	target /= 11

	initial_guess := 100000

	for {
		presents := calcHousePresents2(initial_guess)
		if presents >= target {
			break
		}
		initial_guess += 50
		//fmt.Println("house:", initial_guess, "presents:", presents)
	}

	lowest_house_number := initial_guess

	// probe the last 50 houses to find the lowest one (just in case there is something lower that we missed in the above loop)
	for guess := initial_guess; guess > initial_guess-100; guess-- {
		presents := calcHousePresents2(guess)
		if presents >= target && lowest_house_number > guess {
			//fmt.Println("house:", guess, "presents:", presents)
			lowest_house_number = guess
		}
	}
	presents := calcHousePresents2(lowest_house_number)
	fmt.Println("house:", lowest_house_number, "presents:", presents)
	return lowest_house_number
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	sc.Scan()
	line := sc.Text()
	target, err := strconv.Atoi(line)
	if err != nil {
		fmt.Printf("got an error: %s\n", err.Error())
		return
	}
	result := part1(target)

	fmt.Println("the solution to day20 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	sc.Scan()
	line := sc.Text()
	target, err := strconv.Atoi(line)
	if err != nil {
		fmt.Printf("got an error: %s\n", err.Error())
		return
	}
	result := part2(target)

	fmt.Println("the solution to day20 part 2 is:", result)
}
