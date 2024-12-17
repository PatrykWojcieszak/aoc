package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/2024/pkg/file"
)

func getRegValue(register map[string]int, operand int) int {
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

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	register := make(map[string]int)
	program := []int{}

	for _, row := range output {
		if strings.Contains(row, "Program") {
			splitted := strings.Split(row, ":")
			values := strings.Split(strings.ReplaceAll(splitted[1], " ", ""), ",")

			for _, val := range values {
				num, err1 := strconv.Atoi(string(val))

				if err1 == nil {
					program = append(program, num)
				}
			}
		} else if strings.Contains(row, "Register") {
			splitted := strings.Split(row, ":")
			regName := string(splitted[0][len(splitted[0])-1])

			num, err1 := strconv.Atoi(strings.ReplaceAll(splitted[1], " ", ""))

			if err1 == nil {
				register[regName] = num
			}

			var reg string
			var values string
			fmt.Sscanf(row, "Register %s: %s", &reg, &values)

		}
	}

	result := []string{}
	instrCounter := 0
	for instrCounter < len(program)-1 {
		instr := program[instrCounter]
		instrCounter++
		operand := program[instrCounter]
		operandValue := getRegValue(register, operand)

		if instr == 0 {
			val := math.Floor(float64(register["A"] / int(math.Pow(2, float64(operandValue)))))
			register["A"] = int(val)
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

			for _, digit := range strconv.Itoa(value) {
				result = append(result, string(digit))
			}
		} else if instr == 6 {
			val := math.Floor(float64(register["A"] / int(math.Pow(2, float64(operandValue)))))
			register["B"] = int(val)
		} else if instr == 7 {
			val := math.Floor(float64(register["A"] / int(math.Pow(2, float64(operandValue)))))
			register["C"] = int(val)
		}

		instrCounter++
	}

	for regName, reg := range register {
		fmt.Println("Register", regName, ":", reg)
	}
	fmt.Println("Program", program)
	fmt.Println("Result", strings.Join(result, ","))
}
