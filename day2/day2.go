package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CalcSurfaceArea(width, height, length uint32) uint32 {
	var total uint32 = 0
	// calc areas
	first := width * height
	second := length * height
	third := width * length
	minimum := min(first, second, third)
	total += 2*(first+second+third) + minimum
	return total
}

func CalcRibonLength(width, height, length uint32) uint32 {
	var total uint32 = 0
	// calc smallest perimeter
	first := width + height
	second := length + height
	third := width + length
	minimum := min(first, second, third)
	total += width*height*length + minimum*2
	return total
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	var result uint32 = 0
	dims := []string{"", "", ""}

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		dims = strings.Split(line, "x")
		l, _ := strconv.Atoi(dims[0])
		w, _ := strconv.Atoi(dims[1])
		h, _ := strconv.Atoi(dims[2])
		result += CalcSurfaceArea(uint32(l), uint32(w), uint32(h))
	}
	fmt.Println("the solution to day2 part 1 is:", result)

	if err := sc.Err(); err != nil {
		panic(err)
	}
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)
	var result uint32 = 0
	dims := []string{"", "", ""}

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		dims = strings.Split(line, "x")
		l, _ := strconv.Atoi(dims[0])
		w, _ := strconv.Atoi(dims[1])
		h, _ := strconv.Atoi(dims[2])
		result += CalcRibonLength(uint32(l), uint32(w), uint32(h))
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}
	fmt.Println("the solution to day2 part 2 is:", result)
}
