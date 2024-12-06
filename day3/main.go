package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var enabled bool = true

func main() {

	data, err := readFile()
	if err != nil {
		return
	}

	sumOfPattern(data)
	sumOfPatternWithCondition(data)

}

func sumOfPattern(data string) {
	var totalSum int
	for i := 0; i < len(data); {
		result, index, isValid := checkPatternFromIndex(data, i)
		if isValid {
			totalSum = totalSum + result
			i = index + 1
		} else {
			i = i + 1
		}

	}

	fmt.Printf("result: %d\n", totalSum)
}

func sumOfPatternWithCondition(data string) {
	var totalSum int
	for i := 0; i < len(data); {

		v, followThrough := checkMulEnabled(data, i)
		if !followThrough {
			i = v + 1
			continue
		}

		if enabled {
			result, index, isValid := checkPatternFromIndex(data, i)
			if isValid {
				totalSum = totalSum + result
				i = index + 1
			} else {
				i = i + 1
			}
		} else {
			i = i + 1
		}
	}

	fmt.Printf("result with condition: %d\n", totalSum)
}

func checkMulEnabled(data string, index int) (int, bool) {

	if string(data[index]) == "d" && string(data[index+1]) == "o" &&
		string(data[index+2]) == "n" && string(data[index+3]) == "'" && string(data[index+4]) == "t" &&
		string(data[index+5]) == "(" && string(data[index+6]) == ")" {
		enabled = false
		return index + 6, false
	}

	if string(data[index]) == "d" && string(data[index+1]) == "o" &&
		string(data[index+2]) == "(" && string(data[index+3]) == ")" {
		enabled = true
		return index + 3, false
	}

	return index, true
}

func checkPatternFromIndex(data string, index int) (int, int, bool) {
	if string(data[index]) == "m" && string(data[index+1]) == "u" && string(data[index+2]) == "l" && string(data[index+3]) == "(" {
		k := index + 4
		var digits []string
		var x int
		var y int
		for {
			if !unicode.IsDigit(rune(data[k])) {
				break
			}
			digits = append(digits, string(data[k]))
			k++
		}
		if len(digits) == 0 {
			return 0, k, false
		}
		x, _ = strconv.Atoi(strings.Join(digits, ""))
		digits = digits[:0]
		if string(data[k]) == "," {
			k = k + 1
			for {
				if !unicode.IsDigit(rune(data[k])) {
					break
				}
				digits = append(digits, string(data[k]))
				k++
			}
			if len(digits) == 0 {
				return 0, k, false
			}
			if string(data[k]) != ")" {
				return 0, k, false
			}
			y, _ = strconv.Atoi(strings.Join(digits, ""))
			return x * y, k, true
		}
	}
	return 0, index, false
}

func readFile() (string, error) {
	filePath := "data.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}
