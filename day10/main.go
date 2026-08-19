package day10

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func StepForward(in string) string {
	current := in[0]
	counter := 0

	result := make([]byte, 0, len(in)*2)

	for index := 0; index < len(in); index++ {
		if current == in[index] {
			counter++
		} else {
			result = strconv.AppendInt(result, int64(counter), 10)
			result = append(result, current)

			current = in[index]
			counter = 1
		}
	}
	result = strconv.AppendInt(result, int64(counter), 10)
	result = append(result, current)
	return string(result)
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	sc.Scan()
	result := sc.Text()

	for range 40 {
		result = StepForward(result)
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day10 part 1 is:", len(result))
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)
	sc.Scan()
	result := sc.Text()

	for range 50 {
		result = StepForward(result)
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day10 part 2 is:", len(result))
}
