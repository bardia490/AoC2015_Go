package day18

import (
	"Aoc2015/lib/utility"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func generateLightsArray(array []int, in string, row int) {
	columns_count := len(in)

	for index, val := range in {
		value := 0 // if val == '.'
		if val == '#' {
			value = 1
		}
		array[row*columns_count+index] = value
	}
}

// checks to see if the neighbor is a valid index based on
// the current index and number of rows
func isIndexValid(neighbor, index, rows_count int) bool {
	if neighbor < 0 || neighbor >= rows_count*rows_count {
		return false
	}

	row, col := index/rows_count, index%rows_count
	neighborRow, neighborCol := neighbor/rows_count, neighbor%rows_count

	return math.Abs(float64(row-neighborRow)) <= 1 &&
		math.Abs(float64(col-neighborCol)) <= 1
}

func forwardOneStep(array []int, dest []int, rows_count int) {
	for index := range rows_count * rows_count {
		total_lights_on := 0
		// indexes from the row above
		N := index - rows_count      // North
		NW := index - rows_count - 1 // North West
		NE := index - rows_count + 1 // North East
		// indexes from left and right
		L := index - 1 // Left
		R := index + 1 // Right
		// indexes from the row below
		S := index + rows_count      // South
		SW := index + rows_count - 1 // South West
		SE := index + rows_count + 1 // South East

		if isIndexValid(N, index, rows_count) {
			total_lights_on += array[N]
		}
		if isIndexValid(NW, index, rows_count) {
			total_lights_on += array[NW]
		}
		if isIndexValid(NE, index, rows_count) {
			total_lights_on += array[NE]
		}
		if isIndexValid(L, index, rows_count) {
			//fmt.Println("at index: ", index, "L:", array[L])
			total_lights_on += array[L]
		}
		if isIndexValid(R, index, rows_count) {
			//fmt.Println("at index: ", index, "R:", array[R])
			total_lights_on += array[R]
		}
		if isIndexValid(S, index, rows_count) {
			total_lights_on += array[S]
		}
		if isIndexValid(SW, index, rows_count) {
			//fmt.Println("at index: ", index, "SW:", array[SW])
			total_lights_on += array[SW]
		}
		//fmt.Println("at index: ", index, "SE:", array[SE], isIndexValid(SE, rows_count))
		if isIndexValid(SE, index, rows_count) {
			//fmt.Println("at index: ", index, "SE:", array[SE])
			total_lights_on += array[SE]
		}
		//fmt.Println("index: ", index, "total_lights_on: ", total_lights_on, "light: ", array[index])
		dest[index] = array[index]
		if array[index] == 1 && !(total_lights_on == 2 || total_lights_on == 3) {
			//fmt.Println("index: ", index, "total_lights_on: ", total_lights_on, "light: ", array[index], "turning off")
			dest[index] = 0
		} else if array[index] == 0 && total_lights_on == 3 {
			dest[index] = 1
		}
	}
}

func part1(array []int, number_of_steps int) int {
	rows_count := int(math.Sqrt(float64(len(array))))
	buffer := make([]int, len(array))
	copy(buffer, array)

	for range number_of_steps {
		forwardOneStep(array, buffer, rows_count)
		array, buffer = buffer, array
		//array[0] = 1                                // Left  Up corner
		//array[rows_count-1] = 1                     // right Up corner
		//array[rows_count*rows_count-rows_count] = 1 // Left Down corner
		//array[rows_count*rows_count-1] = 1          // right Down corner
		//for index := 0; index < len(array); index += rows_count {
		//	fmt.Println(index, array[index:index+rows_count])
		//}
		//fmt.Println("================")
	}
	return utility.SumSlice(array)
}

func part2(array []int, number_of_steps int) int {
	rows_count := int(math.Sqrt(float64(len(array))))
	array[0] = 1                                // Left  Up corner
	array[rows_count-1] = 1                     // right Up corner
	array[rows_count*rows_count-rows_count] = 1 // Left Down corner
	array[rows_count*rows_count-1] = 1          // right Down corner
	return part1(array, number_of_steps)
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var array [100 * 100]int
	for sc.Scan() {
		line := sc.Text()
		generateLightsArray(array[:], line, index)
		index += 1
	}
	result := part1(array[:], 100)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day17 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var array [1000 * 1000]int
	for sc.Scan() {
		line := sc.Text()
		generateLightsArray(array[:], line, index)
		index += 1
	}
	result := part2(array[:], 100)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day17 part 2 is:", result)
}
