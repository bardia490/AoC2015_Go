package day16

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Option[T any] struct {
	value T
	ok    bool
}

type Sue struct {
	children    Option[int]
	cats        Option[int]
	samoyeds    Option[int]
	pomeranians Option[int]
	akitas      Option[int]
	vizslas     Option[int]
	goldfish    Option[int]
	trees       Option[int]
	cars        Option[int]
	perfumes    Option[int]
}

func setSueProperty(sues []Sue, property string, property_value int, index int) {
	switch property {
	case "children":
		sues[index].children = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "cats":
		sues[index].cats = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "samoyeds":
		sues[index].samoyeds = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "pomeranians":
		sues[index].pomeranians = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "akitas":
		sues[index].akitas = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "vizslas":
		sues[index].vizslas = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "goldfish":
		sues[index].goldfish = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "trees":
		sues[index].trees = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "cars":
		sues[index].cars = Option[int]{
			value: property_value,
			ok:    true,
		}
	case "perfumes":
		sues[index].perfumes = Option[int]{
			value: property_value,
			ok:    true,
		}
	}
}

func generateSueProperties(sues []Sue, in string, index int) error {
	contents := strings.Split(in, " ")

	for i := 2; i < len(contents)-2; i += 2 {
		property := contents[i][:len(contents[i])-1]
		property_value := contents[i+1][:len(contents[i+1])-1]
		property_value_int, err := strconv.Atoi(property_value)

		if err != nil {
			return err
		}

		setSueProperty(sues, property, property_value_int, index)
	}

	// last property
	last_index := len(contents) - 2
	property := contents[last_index][:len(contents[last_index])-1]
	property_value := contents[last_index+1]
	property_value_int, err := strconv.Atoi(property_value)

	if err != nil {
		return err
	}

	setSueProperty(sues, property, property_value_int, index)

	return nil
}

func part1(sues []Sue) int {
	maximum := 0
	result_index := 0

	for index, sue := range sues {
		temp := 0
		if opt := sue.children; opt.ok && opt.value == 3 {
			temp += 1
		}
		if opt := sue.cats; opt.ok && opt.value == 7 {
			temp += 1
		}
		if opt := sue.samoyeds; opt.ok && opt.value == 2 {
			temp += 1
		}
		if opt := sue.pomeranians; opt.ok && opt.value == 3 {
			temp += 1
		}
		if opt := sue.akitas; opt.ok && opt.value == 0 {
			temp += 1
		}
		if opt := sue.vizslas; opt.ok && opt.value == 0 {
			temp += 1
		}
		if opt := sue.goldfish; opt.ok && opt.value == 5 {
			temp += 1
		}
		if opt := sue.trees; opt.ok && opt.value == 3 {
			temp += 1
		}
		if opt := sue.cars; opt.ok && opt.value == 2 {
			temp += 1
		}
		if opt := sue.perfumes; opt.ok && opt.value == 1 {
			temp += 1
		}
		if temp > maximum {
			result_index = index
			maximum = temp
		}
	}
	return result_index + 1
}

func part2(sues []Sue) int {
	maximum := 0
	result_index := 0

	for index, sue := range sues {
		temp := 0
		if opt := sue.children; opt.ok && opt.value == 3 {
			temp += 1
		}
		if opt := sue.cats; opt.ok && opt.value > 7 {
			temp += 1
		}
		if opt := sue.samoyeds; opt.ok && opt.value == 2 {
			temp += 1
		}
		if opt := sue.pomeranians; opt.ok && opt.value < 3 {
			temp += 1
		}
		if opt := sue.akitas; opt.ok && opt.value == 0 {
			temp += 1
		}
		if opt := sue.vizslas; opt.ok && opt.value == 0 {
			temp += 1
		}
		if opt := sue.goldfish; opt.ok && opt.value < 5 {
			temp += 1
		}
		if opt := sue.trees; opt.ok && opt.value > 3 {
			temp += 1
		}
		if opt := sue.cars; opt.ok && opt.value == 2 {
			temp += 1
		}
		if opt := sue.perfumes; opt.ok && opt.value == 1 {
			temp += 1
		}
		if temp > maximum {
			result_index = index
			maximum = temp
		}
	}
	return result_index + 1
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	sues := [500]Sue{}

	index := 0
	for sc.Scan() {
		line := sc.Text()
		err := generateSueProperties(sues[:], line, index)
		if err != nil {
			panic(fmt.Sprintf("there was a problem: %s, at index: %d", err.Error(), index))
		}
		index += 1
	}
	result := part1(sues[:])

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day16 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)
	sues := [500]Sue{}

	index := 0
	for sc.Scan() {
		line := sc.Text()
		err := generateSueProperties(sues[:], line, index)
		if err != nil {
			panic(fmt.Sprintf("there was a problem: %s", err.Error()))
		}
		index += 1
	}
	result := part2(sues[:])

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day16 part 2 is:", result)
}
