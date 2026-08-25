package day15

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Ingredients struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func generateIngredientsSlice(ingredients []Ingredients, in string, index int) error {
	contents := strings.Split(in, " ")

	// get capacity
	capacity := contents[2][:len(contents[2])-1]
	val, err := strconv.Atoi(capacity)

	if err != nil {
		return fmt.Errorf("could not get capacity: %s", capacity)
	}
	ingredients[index].capacity = val

	// get durability
	durability := contents[4][:len(contents[4])-1]
	val, err = strconv.Atoi(durability)

	if err != nil {
		return fmt.Errorf("could not get durability: %s", durability)
	}
	ingredients[index].durability = val

	// get flavor
	flavor := contents[6][:len(contents[6])-1]
	val, err = strconv.Atoi(flavor)

	if err != nil {
		return fmt.Errorf("could not get flavor: %s", flavor)
	}
	ingredients[index].flavor = val

	// get texture
	texture := contents[8][:len(contents[8])-1]
	val, err = strconv.Atoi(texture)

	if err != nil {
		return fmt.Errorf("could not get texture: %s", texture)
	}
	ingredients[index].texture = val

	// get calories (this time there is no , at the end of the number)
	calories := contents[10]
	val, err = strconv.Atoi(calories)

	if err != nil {
		return fmt.Errorf("could not get calories: %s", calories)
	}
	ingredients[index].calories = val

	return nil
}

func part1(ingredients []Ingredients) int {
	result := 0
	for i := 1; i < 98; i++ {
		for j := 1; j < 98-i; j++ {
			for k := 1; k < 98-i-j; k++ {
				if i+j+k > 97 {
					continue
				}
				last_ingredient := 100 - i - j - k
				total_capacity := i*ingredients[0].capacity + j*ingredients[1].capacity + k*ingredients[2].capacity + last_ingredient*ingredients[3].capacity
				total_durability := i*ingredients[0].durability + j*ingredients[1].durability + k*ingredients[2].durability + last_ingredient*ingredients[3].durability
				total_flavor := i*ingredients[0].flavor + j*ingredients[1].flavor + k*ingredients[2].flavor + last_ingredient*ingredients[3].flavor
				total_texture := i*ingredients[0].texture + j*ingredients[1].texture + k*ingredients[2].texture + last_ingredient*ingredients[3].texture
				if total_capacity < 0 || total_durability < 0 || total_flavor < 0 || total_texture < 0 {
					continue
				}
				temp := total_capacity * total_durability * total_flavor * total_texture

				if temp > result {
					result = temp
				}
			}
		}
	}
	return result
}

func part2(ingredients []Ingredients) int {
	result := 0
	for i := 1; i < 98; i++ {
		for j := 1; j < 98-i; j++ {
			for k := 1; k < 98-i-j; k++ {
				if i+j+k > 97 {
					continue
				}
				last_ingredient := 100 - i - j - k
				total_calories := i*ingredients[0].calories + j*ingredients[1].calories + k*ingredients[2].calories + last_ingredient*ingredients[3].calories
				if total_calories != 500 {
					continue
				}
				total_capacity := i*ingredients[0].capacity + j*ingredients[1].capacity + k*ingredients[2].capacity + last_ingredient*ingredients[3].capacity
				total_durability := i*ingredients[0].durability + j*ingredients[1].durability + k*ingredients[2].durability + last_ingredient*ingredients[3].durability
				total_flavor := i*ingredients[0].flavor + j*ingredients[1].flavor + k*ingredients[2].flavor + last_ingredient*ingredients[3].flavor
				total_texture := i*ingredients[0].texture + j*ingredients[1].texture + k*ingredients[2].texture + last_ingredient*ingredients[3].texture
				if total_capacity < 0 || total_durability < 0 || total_flavor < 0 || total_texture < 0 {
					continue
				}
				temp := total_capacity * total_durability * total_flavor * total_texture

				if temp > result {
					result = temp
				}
			}
		}
	}
	return result
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	ingredients := [4]Ingredients{}

	index := 0
	for sc.Scan() {
		line := sc.Text()
		err := generateIngredientsSlice(ingredients[:], line, index)
		if err != nil {
			panic(fmt.Sprintf("there was a problem: %s", err.Error()))
		}
		index += 1
	}
	result := part1(ingredients[:])

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day15 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)
	ingredients := [4]Ingredients{}

	index := 0
	for sc.Scan() {
		line := sc.Text()
		err := generateIngredientsSlice(ingredients[:], line, index)
		if err != nil {
			panic(fmt.Sprintf("there was a problem: %s", err.Error()))
		}
		index += 1
	}
	result := part2(ingredients[:])

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day15 part 2 is:", result)
}
