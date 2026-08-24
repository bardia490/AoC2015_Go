package day14

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Deer struct {
	speed         int
	speeding_time int
	resting_time  int
}

func generateDeersSpeedMap(deers []Deer, in string, index int) {
	contents := strings.Split(in, " ")
	speed, err := strconv.Atoi(contents[3])
	if err != nil {
		panic(fmt.Sprintf("there was an error parsing the speed for deer: %s, speed: %s, err: %s", contents[0], contents[3], err.Error()))
	}

	speeding_time, err := strconv.Atoi(contents[6])
	if err != nil {
		panic(fmt.Sprintf("there was an error parsing the speeding time for deer: %s, speeding time: %s, err: %s", contents[0], contents[6], err.Error()))
	}

	resting_time, err := strconv.Atoi(contents[13])
	if err != nil {
		panic(fmt.Sprintf("there was an error parsing the resting time for deer: %s, resting time: %s, err: %s", contents[0], contents[13], err.Error()))
	}

	deers[index] = Deer{
		speed:         speed,
		speeding_time: speeding_time,
		resting_time:  resting_time,
	}
}

// time is the total duration of the race
func part1(deers []Deer, time int) int {
	// the aglorithm works as follows:
	// if a dear has the speeding time x
	// resting time of y
	// then we conclude the following:
	// for time <= x + y => the deer does not even complete one full cycle (duration of x+y)
	// this means => if time <= x => travel distance: time * deer_speed
	// else (time > x) => travel distance: x * deer_speed
	// if the time > x + y
	// then int(time / ( x + y))  gives us the number of complete cycles
	// and time % (x + y) gives us the remaining_time for the deers cycle
	// which we can just use the rules from the first part again (time <= x + y)
	maximum := 0
	for _, deer := range deers {
		result := 0
		if time <= deer.speeding_time+deer.resting_time {
			if time <= deer.speeding_time {
				result = time * deer.speed
			} else { // (time > x)
				result = deer.speeding_time * deer.speed
			}
		} else {
			number_of_cycles := int(time / (deer.speeding_time + deer.resting_time))
			remaining_time := time % (deer.speeding_time + deer.resting_time)

			if remaining_time <= deer.speeding_time {
				result = remaining_time * deer.speed
			} else { // (remaining_time > x)
				result = deer.speeding_time * deer.speed
			}

			result += number_of_cycles * (deer.speeding_time * deer.speed)
		}
		if result > maximum {
			maximum = result
		}
	}
	return maximum
}

// this returns the index
func findMaximum(s []int) (index int, value int) {
	maximum := math.MinInt
	index_result := 0

	for index, val := range s {
		if val > maximum {
			maximum = val
			index_result = index
		}
	}
	return index_result, maximum
}

// time is the total duration of the race
func part2(deers []Deer, time int) int {
	deers_points := make([]int, len(deers))
	deers_travel_distance := make([]int, len(deers))
	is_deer_resting := make([]bool, len(deers)) // they should be all false

	t := 0
	winning_index := 0
	for t <= time {
		for index, deer := range deers {
			if !is_deer_resting[index] {
				deers_travel_distance[index] += deer.speed
			}
			remaining_time := (t + 1) % (deer.speeding_time + deer.resting_time)
			switch remaining_time {
			case 0:
				is_deer_resting[index] = false
			case deer.speeding_time:
				is_deer_resting[index] = true
			}
		}
		t += 1
		winning_index, _ = findMaximum(deers_travel_distance)
		deers_points[winning_index] += 1
	}
	_, val := findMaximum(deers_points)
	return val
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	speeds := make([]Deer, 9)

	index := 0
	for sc.Scan() {
		line := sc.Text()
		generateDeersSpeedMap(speeds, line, index)
		index += 1
	}
	result := part1(speeds, 2503)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day14 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)
	speeds := make([]Deer, 9)

	index := 0
	for sc.Scan() {
		line := sc.Text()
		generateDeersSpeedMap(speeds, line, index)
		index += 1
	}
	result := part2(speeds, 2503)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day14 part 2 is:", result)
}
