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

	var first_pair_character byte = 0
	var first_pair_found bool = false
	non_overlapping_pairs := false

	for index := 0; index < len(in)-2; index++ {
		first := in[index]
		second := in[index+1]
		third := in[index+2]

		if first == 'i' || first == 'o' || first == 'l' { // check for forbiden characters
			correctForbidenPassword(in, index)
			return false
		}
		if second == 'i' || second == 'o' || second == 'l' { // check for forbiden characters
			correctForbidenPassword(in, index+1)
			return false
		}
		if third == 'i' || third == 'o' || third == 'l' { // check for forbiden characters
			correctForbidenPassword(in, index+2)
			return false
		}
		if !three_straight_letters && first+1 == second && second+1 == third {
			three_straight_letters = true
		}
		if !non_overlapping_pairs {
			if first == second {
				if first_pair_character != first && first_pair_found {
					non_overlapping_pairs = true
				} else {
					first_pair_character = first
					first_pair_found = true
				}
			} else if second == third {
				if first_pair_character != second && first_pair_found {
					non_overlapping_pairs = true
				} else {
					first_pair_character = second
					first_pair_found = true
				}
			}
		}
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

	fmt.Println("the solution to day11 part 1 is:", new_pass)
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

	fmt.Println("the solution to day11 part 2 is:", new_pass)
}
