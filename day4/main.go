package day4

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
)

func FindCorrectHash1(in []byte) int {
	result := 0
	buffer := make([]byte, len(in), len(in)+20)
	copy(buffer, in)

	for {
		buffer := strconv.AppendInt(buffer, int64(result), 10)
		hash := md5.Sum(buffer)
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 0x10 {
			break
		}
		result += 1
	}
	return result
}

func FindCorrectHash2(in []byte) int {
	result := 282749
	buffer := make([]byte, len(in), len(in)+20)
	copy(buffer, in)

	for {
		buffer := strconv.AppendInt(buffer[:len(in)], int64(result), 10)
		hash := md5.Sum(buffer)
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			break
		}
		result += 1
	}
	return result
}

func Solution1(f *os.File) {
	data, err := io.ReadAll(f)
	data = bytes.TrimSpace(data)
	f.Seek(0, io.SeekStart)

	if err != nil {
		fmt.Println("something was wrong with the file input4, solution failed")
	}
	result := FindCorrectHash1(data)
	fmt.Println("the solution to day4 part 1 is:", result)
}

func Solution2(f *os.File) {
	data, err := io.ReadAll(f)
	data = bytes.TrimSpace(data)

	if err != nil {
		fmt.Println("something was wrong with the file input4, solution failed")
	}
	result := FindCorrectHash2(data)
	fmt.Println("the solution to day4 part 2 is:", result)
}
