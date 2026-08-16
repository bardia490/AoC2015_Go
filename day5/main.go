package day5

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func NaughtyOrNice(in string) bool {
	sum := 0                     // for the vowels
	last_letter := ' '           // a place holder to check for the last letter
	double_letter_flag := false  // to check for a double letters
	illegal_combination := false // to check for any "ab", "cd", "pq" and "xy"

	for _, ch := range in {
		switch ch {
		case 'a':
			sum += 1
		case 'e':
			sum += 1
		case 'i':
			sum += 1
		case 'o':
			sum += 1
		case 'u':
			sum += 1
		}

		if (last_letter == 'a' && ch == 'b') || (last_letter == 'c' && ch == 'd') || (last_letter == 'p' && ch == 'q') || (last_letter == 'x' && ch == 'y') {
			illegal_combination = true
			break
		}
		if !double_letter_flag {
			if last_letter == ch {
				double_letter_flag = true
			}
		}
		last_letter = ch
	}
	return sum >= 3 && !illegal_combination && double_letter_flag
}

func NaughtyOrNice2(in string) bool {
	two_pairs := false
	repeating_letter := false
	scan_range := len(in) - 2

	for i := range scan_range {
		// look for two pairs if we havn't found them
		if !two_pairs {
			first_two := in[i : i+2]
			for j := i + 2; j <= len(in)-2; j++ {
				second_two := in[j : j+2]
				if first_two == second_two {
					two_pairs = true
				}
			}
		}
		// look for two pairs if we haven't found them
		if !repeating_letter { // && i != len(in)-2
			if in[i] == in[i+2] {
				repeating_letter = true
			}
		}

		if repeating_letter && two_pairs {
			break
		}
	}
	return repeating_letter && two_pairs
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	nice_strings_count := 0

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		if NaughtyOrNice(strings.TrimSpace(line)) {
			nice_strings_count += 1
		}
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	fmt.Println("the solution to day5 part 1 is:", nice_strings_count)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)
	nice_strings_count := 0

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		if NaughtyOrNice2(strings.TrimSpace(line)) {
			nice_strings_count += 1
		}
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	fmt.Println("the solution to day5 part 2 is:", nice_strings_count)
}
