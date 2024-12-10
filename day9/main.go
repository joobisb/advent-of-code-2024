package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"
)

func main() {

	numbers, err := readFile()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Numbers: %v\n", numbers)

	var idSlice []string
	var idCount int
	for i, v := range numbers {
		var isBlock bool
		for k := 0; k < v; k++ {
			if i%2 == 0 {
				idSlice = append(idSlice, strconv.Itoa(idCount))
				isBlock = true
			} else {
				idSlice = append(idSlice, ".")
			}
		}

		if isBlock {
			idCount++
		}
	}

	for i, j := 0, len(idSlice)-1; i < j; {
		_, errLeft := strconv.Atoi(idSlice[i])
		_, errRight := strconv.Atoi(idSlice[j])

		if idSlice[i] == "." && errRight == nil {
			temp := idSlice[i]
			idSlice[i] = idSlice[j]
			idSlice[j] = temp
			i++
			j--
		} else if errLeft == nil && idSlice[j] == "." {
			i++
			j--
		} else if idSlice[i] == "." && idSlice[j] == "." {
			j--
		} else if errLeft == nil && errRight == nil {
			i++
		}
	}

	var totalSum int
	for i, v := range idSlice {
		intV, _ := strconv.Atoi(v)
		totalSum += (i * intV)
	}

	fmt.Printf("totalSum: %d\n", totalSum)

}

func readFile() ([]int, error) {
	filePath := "data.txt"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': %w", filePath, err)
	}

	content := string(data)
	var numbers []int

	for _, r := range content {
		if unicode.IsDigit(r) {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("failed to convert rune '%c' to integer: %w", r, err)
			}
			numbers = append(numbers, num)
		}
	}

	return numbers, nil

}
