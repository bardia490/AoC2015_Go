package day8

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountNumberOfCharacters(in string) int {
	return len(in)
}

func CountNumberOfInMemoryCharacters(in string) int {
	result := 0
	index := 1

	for index < len(in)-1 { // -2 is because every string contains a " at the begining and end of it
		if in[index] == '\\' {
			if in[index+1] == 'x' {
				index += 4
			} else {
				index += 2
			}
			result += 1
		} else {
			index += 1
			result += 1
		}
	}
	return result
}

func CountNumberOfCharacterEncodings(in string) int {
	result := 6 // because every " at the begining and end is going to be replaced by "\". e.g "" => "\"\""
	index := 1

	for index < len(in)-1 { // -2 is because every string contains a " at the begining and end of it
		if in[index] == '\\' {
			if in[index+1] == 'x' {
				index += 4
				result += 5
			} else {
				index += 2 // because the \" (or similar things) turn into \\\"
				result += 4
			}
		} else {
			index += 1
			result += 1
		}
	}
	return result
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	number_of_characters := 0
	number_of_in_memory_characters := 0
	for sc.Scan() {
		line := sc.Text()
		number_of_characters += CountNumberOfCharacters(line)
		number_of_in_memory_characters += CountNumberOfInMemoryCharacters(line)
	}
	if err := sc.Err(); err != nil {
		fmt.Printf("something went wrong with file, %s", err.Error())
	}
	fmt.Println("the solution to day8 part 1 is:", number_of_characters-number_of_in_memory_characters)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	number_of_characters := 0
	number_of_encoded_characters := 0
	for sc.Scan() {
		line := sc.Text()
		number_of_characters += CountNumberOfCharacters(line)
		number_of_encoded_characters += CountNumberOfCharacterEncodings(line)
	}
	if err := sc.Err(); err != nil {
		fmt.Printf("something went wrong with file, %s", err.Error())
	}
	fmt.Println("the solution to day8 part 2 is:", number_of_encoded_characters-number_of_characters)
}
