package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var totalSum int

func main() {

	sumSlice, sliceOfInp, _ := readFile()

	var interResult [][]int
	for sliceIndex, inp := range sliceOfInp {
		interResult = [][]int{}
		for x := 0; x < len(inp)-1; x++ {
			if len(interResult) == 0 {

				var interTempSlice []int
				interTempSlice = append(interTempSlice, inp[x]*inp[x+1])
				interTempSlice = append(interTempSlice, (inp[x] + inp[x+1]))

				interResult = append(interResult, interTempSlice)
				continue

			}
			var interTempSlice []int

			for k := 0; k < len(interResult[x-1]); k++ {
				interTempSlice = append(interTempSlice, interResult[x-1][k]*inp[x+1])
				interTempSlice = append(interTempSlice, (interResult[x-1][k] + inp[x+1]))

			}
			interResult = append(interResult, interTempSlice)

		}

		for m := 0; m < len(interResult[len(inp)-2]); m++ {
			if sumSlice[sliceIndex] == interResult[len(inp)-2][m] {
				totalSum = totalSum + sumSlice[sliceIndex]
				break
			}
		}
	}
	fmt.Printf("totalSum: %d\n", totalSum)
	solveWithConcatenateOperator(sliceOfInp, sumSlice)

}

func solveWithConcatenateOperator(sliceOfInp [][]int, sumSlice []int) {
	var interResult [][]int

	for sliceIndex, inp := range sliceOfInp {
		interResult = [][]int{}
		for x := 0; x < len(inp)-1; x++ {
			if len(interResult) == 0 {

				var interTempSlice []int
				interTempSlice = append(interTempSlice, inp[x]*inp[x+1])
				interTempSlice = append(interTempSlice, (inp[x] + inp[x+1]))

				//concantenates here, by appending
				appended := (inp[x] * appendUtil(inp[x+1])) + inp[x+1]
				interTempSlice = append(interTempSlice, appended)

				interResult = append(interResult, interTempSlice)
				continue

			}
			var interTempSlice []int

			for k := 0; k < len(interResult[x-1]); k++ {
				interTempSlice = append(interTempSlice, interResult[x-1][k]*inp[x+1])
				interTempSlice = append(interTempSlice, (interResult[x-1][k] + inp[x+1]))

				//concantenates here, by appending
				appended := interResult[x-1][k]*appendUtil(inp[x+1]) + inp[x+1]
				interTempSlice = append(interTempSlice, appended)
			}
			interResult = append(interResult, interTempSlice)

		}

		for m := 0; m < len(interResult[len(inp)-2]); m++ {
			if sumSlice[sliceIndex] == interResult[len(inp)-2][m] {
				totalSum = totalSum + sumSlice[sliceIndex]
				break
			}
		}
	}
	fmt.Printf("totalSumWithConcatenateOperator: %d\n", totalSum)

}

func appendUtil(n int) int {
	var x = 1

	for x <= n {
		x = x * 10
	}
	return x

}

func readFile() ([]int, [][]int, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}
	defer file.Close()

	var sumArray []int
	var sliceOfSlices [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		keyStr := strings.TrimSpace(parts[0])
		key, err := strconv.Atoi(keyStr)
		if err != nil {
			fmt.Println("Error parsing key:", err)
			continue
		}

		valueStrs := strings.Fields(strings.TrimSpace(parts[1]))
		var values []int
		for _, vStr := range valueStrs {
			v, err := strconv.Atoi(vStr)
			if err != nil {
				fmt.Println("Error parsing value:", err)
				continue
			}
			values = append(values, v)
		}

		sumArray = append(sumArray, key)
		sliceOfSlices = append(sliceOfSlices, values)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return sumArray, sliceOfSlices, nil

}
