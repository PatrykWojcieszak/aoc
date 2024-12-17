package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

func getRegValue(register map[string]uint64, operand uint64) uint64 {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return register["A"]
	} else if operand == 5 {
		return register["B"]
	} else if operand == 6 {
		return register["C"]
	}

	return 0
}

func runProgram(program []uint64, seed uint64) []uint64 {
	register := make(map[string]uint64)
	register["A"] = seed
	register["B"] = 0
	register["C"] = 0

	result := []uint64{}
	var instrCounter uint64 = 0

	for instrCounter < uint64(len(program)-1) {
		instr := program[instrCounter]
		instrCounter++
		operand := program[instrCounter]
		operandValue := getRegValue(register, operand)

		if instr == 0 {
			val := math.Floor(float64(register["A"] / uint64(math.Pow(2, float64(operandValue)))))
			register["A"] = uint64(val)
		} else if instr == 1 {
			register["B"] = register["B"] ^ operand
		} else if instr == 2 {
			register["B"] = operandValue % 8
		} else if instr == 3 {
			if register["A"] != 0 {
				instrCounter = operand
				continue
			}
		} else if instr == 4 {
			register["B"] = register["B"] ^ register["C"]
		} else if instr == 5 {
			value := operandValue % 8

			for _, digit := range fmt.Sprintf("%v", value) {
				num, err := strconv.Atoi(string(digit))

				if err == nil {
					result = append(result, uint64(num))
				}
			}
		} else if instr == 6 {
			val := math.Floor(float64(register["A"] / uint64(math.Pow(2, float64(operandValue)))))
			register["B"] = uint64(val)
		} else if instr == 7 {
			val := math.Floor(float64(register["A"] / uint64(math.Pow(2, float64(operandValue)))))
			register["C"] = uint64(val)
		}

		instrCounter++
	}

	return result
}

func partTwo(program []uint64) (seed uint64) {
	for itr := len(program) - 1; itr >= 0; itr-- {
		seed <<= 3
		for !slices.Equal(runProgram(program, seed), program[itr:]) {
			seed++
		}
	}
	return
}

func main() {
	fmt.Println("Part Two:", partTwo([]uint64{2, 4, 1, 1, 7, 5, 1, 5, 4, 3, 0, 3, 5, 5, 3, 0}))
}
