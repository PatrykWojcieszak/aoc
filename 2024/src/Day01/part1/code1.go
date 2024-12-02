package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var arr1, arr2 []int
    var diffValues []int

	for scanner.Scan() {
		line := scanner.Text() 
        values := strings.Fields(line)

        intOneVal, err1 := strconv.Atoi(values[0])
        intTwoVal, err2 := strconv.Atoi(values[1])

        if err1 != nil || err2 != nil {
			fmt.Printf("Error parsing integers in line: %s\n", line)
			return
		}

        arr1 = append(arr1, intOneVal);
        arr2 = append(arr2, intTwoVal);
	}
    
	if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }

    sort.Ints(arr1)
    sort.Ints(arr2)

    for index, value := range arr1 {
        var diff = value - arr2[index]
        // fmt.Println(value, arr2[index])
        diffValues = append(diffValues, int(absInt(diff)))
    }
    
    var sum = 0

    for _, value := range diffValues {
        sum = sum + value;
    }

    fmt.Println(sum)
}