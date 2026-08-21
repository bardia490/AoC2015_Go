package day12

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func countNumbers(in []byte, solution_number int) (float64, error) {
	var f any

	err := json.Unmarshal(in, &f)
	if err != nil {
		return 0, err
	}
	if solution_number == 1 {
		return countMaps(f.(map[string]any)), nil
	}
	return countMaps(f.(map[string]any)), nil
}

func scanForRedInMaps(m map[string]any) bool {
	for k, v := range m {
		if k == "red" {
			return true
		}
		switch vv := v.(type) {
		case []any:
			if scanForRedInArrays(vv) {
				return true
			}
		case map[string]any:
			if scanForRedInMaps(vv) {
				return true
			}
		}
	}
	return false
}

func scanForRedInArrays(m []any) bool {
	for _, v := range m {
		switch vv := v.(type) {
		case []any: // i need to think more about this one
			if scanForRedInArrays(vv) {
				return true
			}
		case map[string]any:
			if scanForRedInMaps(vv) {
				return true
			}
		}
	}
	return true
}

func countMaps(m map[string]any) float64 {
	var result float64 = 0

	for k, v := range m {
		switch vv := v.(type) {
		//case int:
		//	fmt.Println(k, "is int", vv)
		//	result += vv
		case float64:
			fmt.Println(k, "is float64", vv)
			result += vv
		case []any:
			//fmt.Println(k, "is an array:")
			result += countArrays(vv)
		case map[string]any:
			//fmt.Println(k, "is an object:")
			result += countMaps(vv)
		}
	}
	return result
}

func countArrays(m []any) float64 {
	var result float64 = 0

	for _, v := range m {
		switch vv := v.(type) {
		case float64:
			result += vv
		case []any:
			//fmt.Println(k, "is an array:")
			result += countArrays(vv)
		case map[string]any:
			//fmt.Println(k, "is an object:")
			result += countMaps(vv)
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
	result, err = countNumbers(contents, 1)
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
	result, err = countNumbers(contents, 2)
	if err != nil {
		fmt.Println("something went wrong with parsing the json: ", err.Error())
	}

	fmt.Println("the solution to day12 part 2 is:", result)
}
