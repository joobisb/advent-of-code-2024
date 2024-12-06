package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	data, err := readFile()
	if err != nil {
		return
	}

	var totalSum int
	for i := 0; i < len(data); {
		result, isValid := checkPatternFromIndex(data, i)
		if isValid {
			totalSum = totalSum + result
			i += 8
		}
		i = i + 1

	}

	fmt.Printf("result: %d", totalSum)

}

func checkPatternFromIndex(data string, index int) (int, bool) {
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
			return 0, false
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
				return 0, false
			}
			if string(data[k]) != ")" {
				return 0, false
			}
			y, _ = strconv.Atoi(strings.Join(digits, ""))
			return x * y, true
		}
	}
	return 0, false
}

func readFile() (string, error) {
	filePath := "data.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}
