package day25

import (
	"fmt"
	"os"
)

// grid
//
//	  |    1         2         3         4         5         6		  7			 8
// ---+---------+---------+---------+---------+---------+---------+---------+---------+
// 1 | 20151125  18749137  17289845  30943339  10071777  33511524     num2      num6
// 2 | 31916031  21629792  16929656   7726640  15514188   4041754     num5
// 3 | 16080970   8057251   1601130   7981243  11661866  16474243
// 4 | 24592653  32451966  21345942   9380097  10600672  31527494
// 5 |    77061  17552253  28094349   6899651   9250759  31663883
// 6 | 33071741   6796745  25397450  24659492   1534922  27995004
// 7 |    num1     num4
// 8 |    num3

func part1() uint64 {
	//last_value1 := 20151125
	//for row := 2; row < 13; row++ {
	//	for col := 1; col != row+1; col++ {
	//		last_value1 = last_value1 * 252533 % 33554393 // update the last value
	//		fmt.Println(row-col+1, col, last_value1)
	//	}
	//}
	// generate the elements one by one by keep tracking the last one
	// don't need to keep track of the full table at all
	// every new row goes like this: row = x, col = 1
	// first element(x,1): last_element * 252533 % 33554393 => last_element = first_element
	// second element(x-1,2): last_element * 252533 % 33554393 => last_element = second_element
	// ... (repeat)
	// when x = 1, col = row => this is the end of this row
	// (move on to the next row) row += 1, col = 1
	// for every row we have the repeat the above process [row] times
	// repeat this process until we get to the desired row and col (row 2978, column 3083 for me).
	// based on my puzzle inputs I have to go til row = 2978 + 3083 - 1 = 6060

	// i wont repeat the values that i already have (the last value is at 6,6 => 27995004)
	var last_value uint64 = 27995004
	current_row, current_col, row := 6, 6, 11 // currently at row 11 (row+col-1), but our point is at 6, 6
	// get to the start of the next row
	for ; current_row != 0; current_row-- {
		last_value = last_value * 252533 % 33554393 // update the last value
	}
	// get to the desired row
	for row = 12; row != 6060; row++ { // for my puzzle input
		// every time we decrement the current_row we are adding to the current_col but it doesn't realy matter for this loop
		for current_row = row; current_row != 0; current_row-- {
			last_value = last_value * 252533 % 33554393 // update the last value
		}
	}
	// at the end of the last loop we reached our final row
	for current_row, current_col = row, 1; current_row != 2978 && current_col != 3083; current_row, current_col = current_row-1, current_col+1 {
		last_value = last_value * 252533 % 33554393 // update the last value
	}
	return last_value
}

func part2() int {
	return 0
}

func Solution1(f *os.File) {
	result := part1()
	fmt.Println("the solution to day24 part 1 is:", result)
}

func Solution2(f *os.File) {
	_ = part2()
	fmt.Println("sure enough, a lone star sits at the bottom, awaiting its friends. Looks like you need to provide 49 yourself.")
}
