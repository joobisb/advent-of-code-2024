package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	filename := "data.txt"

	data, err := readInputFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var unsafe int
	var unsafeWithoutTolerance int
	for _, v := range data {
		var isAsc bool
		if len(v) < 2 {
			continue
		}
		if v[1] > v[0] {
			isAsc = true
		}
		if !validateAdjacent(v, 0) {
			unsafe++
			unsafeWithoutTolerance++
			for i := 0; i < len(v); i++ {
				if checkIfTolerant(v, i) {
					unsafe--
					break
				}
			}
			continue
		}
		var checkPointFlag int
		for i := 1; i < len(v)-1; i++ {
			if isAsc {
				if v[i+1] < v[i] {
					checkPointFlag = i
					unsafe++
					unsafeWithoutTolerance++
					break
				}
			} else {
				if v[i+1] > v[i] {
					checkPointFlag = i
					unsafe++
					unsafeWithoutTolerance++
					break
				}
			}
			if !validateAdjacent(v, i) {
				checkPointFlag = i
				unsafe++
				unsafeWithoutTolerance++
				break
			}
		}
		if checkPointFlag != 0 {
			for i := 0; i < len(v); i++ {
				if checkIfTolerant(v, i) {
					unsafe--
					break
				}
			}
		}
	}
	fmt.Printf("safeWithoutTolerance: %d\n", len(data)-unsafeWithoutTolerance)
	fmt.Printf("safeWithTolerance: %d\n", len(data)-unsafe)
}

func validateAdjacent(v []int, index int) bool {
	if (math.Abs(float64(v[index+1]-v[index]))) < 1 || math.Abs(float64(v[index+1]-v[index])) > 3 {
		return false
	}
	return true
}

func checkIfTolerant(slice []int, index int) bool {
	v := make([]int, len(slice))
	copy(v, slice)

	v = append(v[:index], v[index+1:]...)

	var unsafe int

	var isAsc bool
	if len(v) < 2 {
		return true
	}
	if v[1] > v[0] {
		isAsc = true
	}
	if !validateAdjacent(v, 0) {
		unsafe++
		return false
	}
	for i := 1; i < len(v)-1; i++ {
		if isAsc {
			if v[i+1] < v[i] {
				unsafe++
				break
			}
		} else {
			if v[i+1] > v[i] {
				unsafe++
				break
			}
		}
		if !validateAdjacent(v, i) {
			unsafe++
			break
		}
	}
	if unsafe > 0 {
		return false
	}
	return true
}

func readInputFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numStrings := strings.Fields(line)

		var row []int
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("error converting %q to int: %v", numStr, err)
			}
			row = append(row, num)
		}

		result = append(result, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return result, nil
}
