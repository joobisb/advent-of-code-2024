package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	filename := "data.txt"
	leftList, rightList, err := readColumnsFromFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	fmt.Printf("distance: %d\n", int(distance(leftList, rightList)))
	fmt.Printf("similarity: %d\n", similarity(leftList, rightList))
}

func distance(leftList, rightList []int) float64 {
	var distance float64
	for i := 0; i < len(leftList); i++ {
		distance = distance + math.Abs(float64(leftList[i]-(rightList[i])))
	}
	return distance
}

func similarity(leftList, rightList []int) int {
	var totalSimilarity int
	var existingMap = make(map[int]int)
	var k int

	for i := 0; i < len(leftList); i++ {
		var similarityOccurence int
		val, ok := existingMap[leftList[i]]
		if ok {
			totalSimilarity = totalSimilarity + val
			continue
		}
		for k < int(len(leftList)) {
			if leftList[i] == rightList[k] {
				similarityOccurence++
				k++
			} else {
				if rightList[k] < leftList[i] {
					k++
				} else {
					break
				}
			}
		}
		totalSimilarity = totalSimilarity + leftList[i]*similarityOccurence
		existingMap[leftList[i]] = leftList[i] * similarityOccurence
	}
	return totalSimilarity
}

func readColumnsFromFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, "   ")

		for i, val := range columns {
			num, err := strconv.Atoi(strings.TrimSpace(val))
			if err != nil {
				return nil, nil, fmt.Errorf("error converting value to integer: %v", err)
			}
			if i%2 == 0 {
				col1 = append(col1, num)
			} else {
				col2 = append(col2, num)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return col1, col2, nil
}
