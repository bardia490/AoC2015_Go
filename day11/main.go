package day11

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// this function will panic wen the index == 0 and in[index] == 'z'
func forwadPassword(in []byte) {
	index := len(in) - 1

	for { // run until incrementation is complete
		if in[index] == 'z' {
			in[index] = 'a'

			if index == 0 {
				panic("reached index 0 with no password being found")
			}
			index -= 1 // move backwards one more character
			continue
		}
		// if the in[index] != 'z' we can increment safely
		in[index] += 1
		if in[index] == 'i' || in[index] == 'o' || in[index] == 'l' {
			correctForbidenPassword(in, index)
		}
		break
	}
}

func correctForbidenPassword(in []byte, index int) {
	for i := index + 1; i < len(in); i++ {
		in[i] = 'a'
	}
	in[index] += 1
}

func checkPassword(in []byte) bool {
	three_straight_letters := false
	var last_char byte = 0
	straight_letters_counter := 1

	var first_pair_character byte = 0
	non_overlapping_pairs := false

	for index, val := range in {
		if val == 'i' || val == 'o' || val == 'l' {
			correctForbidenPassword(in, index)
			return false
		}
		if !three_straight_letters {
			//fmt.Println("the letter and the last letter", string(last_char), string(val))
			if val == last_char+1 {
				straight_letters_counter++
				if straight_letters_counter == 3 {
					three_straight_letters = true
				}
			} else { // reset back to zero because they werent back to back
				//fmt.Println("the letter and the last letter were not back to back", string(last_char), string(val))
				straight_letters_counter = 1
			}
		} // three_straight_letters
		if !non_overlapping_pairs {
			if val == last_char {
				if first_pair_character == 0 {
					first_pair_character = val
				} else if first_pair_character != val {
					non_overlapping_pairs = true
				}
			}
		} // end of non_overlapping_pairs check
		last_char = val
	}
	return non_overlapping_pairs && three_straight_letters
}

func generateNewPassword(pass []byte) string {
	for !checkPassword(pass) {
		forwadPassword(pass)
	}
	return string(pass)
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	sc.Scan()
	result := sc.Bytes()
	new_pass := generateNewPassword(result)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day10 part 1 is:", new_pass)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)
	sc.Scan()
	result := sc.Bytes()
	new_pass := generateNewPassword(result)
	result = ([]byte(new_pass))
	forwadPassword(result) // to expire the password
	new_pass = generateNewPassword(result)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day10 part 2 is:", new_pass)
}
