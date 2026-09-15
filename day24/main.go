package day24

import (
	"Aoc2015/lib/utility"
	"bufio"
	"fmt"
	"math"
	"os"
)

var weights = make([]int, 28)

func setWeights(in []byte, weights []int, index int) (parsing_error error) {
	val, err := utility.BytesToIntSigned(in)
	if err != nil {
		return err
	}

	weights[index] = val
	return nil
}

func part1(weights []int) (uint64, error) {
	groups_weight := utility.SumSlice(weights) / 3 // the weight of each group (which is 508)
	// find the first group that matches the groups_weight (don't need to check the rest of the groups)
	// since the weights are odd and the target group_weight is even, then i need even amount of the groups
	len_of_weights := len(weights)
	var min_quantum_entanglement uint64 = math.MaxUint64
	for number_op_items := 2; number_op_items < len_of_weights; number_op_items += 2 {
		subsets, err := utility.GenerateSubsetsN(weights, number_op_items)
		if err != nil {
			fmt.Println("there was an error with range:", err.Error())
			return 0, err
		}
		for current_weights := range subsets {
			if utility.SumSlice(current_weights) == groups_weight {
				var current_quantum_entanglement uint64 = 1
				for _, weight := range current_weights {
					current_quantum_entanglement *= uint64(weight)
				}
				if current_quantum_entanglement < min_quantum_entanglement {
					min_quantum_entanglement = current_quantum_entanglement
				}
			}
		}
		if min_quantum_entanglement != math.MaxUint64 {
			break
		}
	}
	return min_quantum_entanglement, nil
}

func part2(weights []int) (uint64, error) {
	groups_weight := utility.SumSlice(weights) / 4 // the weight of each group
	// find the first group that matches the groups_weight (don't need to check the rest of the groups)
	len_of_weights := len(weights)
	var min_quantum_entanglement uint64 = math.MaxUint64
	for number_op_items := 2; number_op_items < len_of_weights; number_op_items += 1 {
		subsets, err := utility.GenerateSubsetsN(weights, number_op_items)
		if err != nil {
			fmt.Println("there was an error with range:", err.Error())
			return 0, err
		}
		for current_weights := range subsets {
			if utility.SumSlice(current_weights) == groups_weight {
				var current_quantum_entanglement uint64 = 1
				for _, weight := range current_weights {
					current_quantum_entanglement *= uint64(weight)
				}
				if current_quantum_entanglement < min_quantum_entanglement {
					min_quantum_entanglement = current_quantum_entanglement
				}
			}
		}
		if min_quantum_entanglement != math.MaxUint64 {
			break
		}
	}
	return min_quantum_entanglement, nil
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	index := 0

	for sc.Scan() {
		line := sc.Bytes()
		_ = line
		if err := setWeights(line, weights, index); err != nil {
			fmt.Println("problem parsing the weights:", err.Error())
			return
		}
		index += 1
	}
	result, err := part1(weights)
	if err != nil {
		fmt.Println("what error?", err.Error())
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day24 part 1 is:", result)
}

func Solution2(f *os.File) {
	// we don't even need to parse the file again
	result, _ := part2(weights)

	fmt.Println("the solution to day24 part 2 is:", result)
}
