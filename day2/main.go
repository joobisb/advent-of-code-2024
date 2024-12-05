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
	for _, v := range data {
		var isAsc bool
		if v[1] > v[0] {
			isAsc = true
		}
		if !validateAdjacent(v, 0) {
			unsafe++
			continue
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
	}
	fmt.Printf("safe: %d", len(data)-unsafe)

}

func validateAdjacent(v []int, index int) bool {
	if (math.Abs(float64(v[index+1]-v[index]))) < 1 || math.Abs(float64(v[index+1]-v[index])) > 3 {
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
