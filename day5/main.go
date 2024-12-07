package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var orderMap = make(map[string][]string)
var inputSlice [][]string

func main() {
	readAndAssign()

	var sum int
	for _, inp := range inputSlice {
		var checkSlice []string
		i := 0
		checkSlice = append(checkSlice, inp[0])
		checkFlag := false
		for i = 1; i < len(inp); i++ {
			orderingsForI := orderMap[inp[i]]

			for _, chVal := range checkSlice {
				for _, ordVal := range orderingsForI {
					if ordVal == chVal {
						checkFlag = true
						break
					}
				}
			}
			checkSlice = append(checkSlice, inp[i])
		}
		if !checkFlag {
			midElem := inp[len(inp)/2]
			midElemInt, _ := strconv.Atoi(midElem)
			sum += int(midElemInt)
		}
	}
	fmt.Printf("mid element sum: %d\n", sum)
}

func readAndAssign() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isInputSliceSection := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			isInputSliceSection = true
			continue
		}

		if isInputSliceSection {
			parts := strings.Split(line, ",")
			inputSlice = append(inputSlice, parts)
		} else {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				orderMap[parts[0]] = append(orderMap[parts[0]], parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

}
