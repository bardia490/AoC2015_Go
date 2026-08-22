package day12

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// calculates the sum of the numbers in the json file
// the error is for when the json data is corrupted and cannot be parsed correctly
func sumNumbers(in []byte, solution_number int) (float64, error) {
	var f any

	err := json.Unmarshal(in, &f)
	if err != nil {
		return 0, err
	}
	if solution_number == 1 {
		return sumMaps(f.(map[string]any)), nil
	}
	return sumMaps2(f.(map[string]any)), nil
}

// sums the numbers it finds in the json (which is in map[string]any format)
func sumMaps(m map[string]any) float64 {
	var result float64 = 0

	for _, v := range m {
		switch vv := v.(type) {
		case float64:
			result += vv
		case []any:
			result += sumArrays(vv)
		case map[string]any:
			result += sumMaps(vv)
		}
	}
	return result
}

// sums the numbers it finds in the json (which is in []any format)
func sumArrays(m []any) float64 {
	var result float64 = 0

	for _, v := range m {
		switch vv := v.(type) {
		case float64:
			result += vv
		case []any:
			result += sumArrays(vv)
		case map[string]any:
			result += sumMaps(vv)
		}
	}
	return result
}

// sums the numbers it finds in the json (which is in map[string]any format)
// excluding the objects that hold any property with the value "red"
func sumMaps2(m map[string]any) float64 {
	var result float64 = 0

	for _, v := range m {
		switch vv := v.(type) {
		case string:
			if vv == "red" {
				return 0
			}
		case float64:
			result += vv
		case []any:
			result += sumArrays2(vv)
		case map[string]any:
			result += sumMaps2(vv)
		}
	}
	return result
}

// sums the numbers it finds in the json (which is in []any format)
// excluding the objects that hold any property with the value "red"
func sumArrays2(m []any) float64 {
	var result float64 = 0

	for _, v := range m {
		switch vv := v.(type) {
		case float64:
			result += vv
		case []any:
			result += sumArrays2(vv)
		case map[string]any:
			result += sumMaps2(vv)
		}
	}
	return result
}

func Solution1(f *os.File) {
	contents, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("something went wrong with reading the file: ", err.Error())
		return
	}
	var result float64 = 0
	result, err = sumNumbers(contents, 1)
	if err != nil {
		fmt.Println("something went wrong with parsing the json: ", err.Error())
	}

	fmt.Println("the solution to day12 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	contents, err := io.ReadAll(f)

	if err != nil {
		fmt.Println("something went wrong with reading the file: ", err.Error())
		return
	}
	var result float64 = 0
	result, err = sumNumbers(contents, 2)
	if err != nil {
		fmt.Println("something went wrong with parsing the json: ", err.Error())
	}

	fmt.Println("the solution to day12 part 2 is:", result)
}
