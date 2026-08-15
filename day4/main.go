package day4

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func FindCorrectHash1(in []byte) int {
	result := 0

	for {
		val := fmt.Appendf(in, "%d", result)
		hash := md5.Sum(val)
		s := hex.EncodeToString(hash[:])
		if s[0:5] == "00000" {
			//fmt.Println("the val and s are:", string(val), s)
			break
		}
		result += 1
	}
	return result
}

func FindCorrectHash2(in []byte) int {
	result := 0

	for {
		val := fmt.Appendf(in, "%d", result)
		hash := md5.Sum(val)
		s := hex.EncodeToString(hash[:])
		if s[0:6] == "000000" {
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
