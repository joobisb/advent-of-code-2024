package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var score int

func main() {

	data, _ := readFile()
	fmt.Println("2D Array:")
	for _, row := range data {
		fmt.Println(row)
	}
	rowLen := len(data)
	colLen := len(data[0])

	for i, row := range data {
		for j, _ := range row {
			if data[i][j] == "0" {
				tempVisited := make(map[int]int)
				checkScores(tempVisited, rowLen, colLen, 0, i, j, data)
				fmt.Println("check")
			}
		}
	}

	fmt.Printf("score: %d\n", score)
}

func checkScores(tempVisited map[int]int, rowLen, colLen, prev, i, j int, data [][]string) {
	if prev == 9 {
		if tempVisited[i] != j {
			score++
		}
		if tempVisited[i] == j {
			fmt.Println("here again")
		}
		tempVisited[i] = j
		return
	}
	if i < 0 || j < 0 || i > rowLen || j > colLen {
		if prev == 9 {
			score++
		}
		return
	}

	if i < rowLen-1 {
		shiftDownVal, _ := strconv.Atoi(data[i+1][j])
		if shiftDownVal == prev+1 {
			checkScores(tempVisited, rowLen, colLen, prev+1, i+1, j, data)
		}
	}
	if j > 0 {
		shiftLeftVal, _ := strconv.Atoi(data[i][j-1])
		if shiftLeftVal == prev+1 {
			checkScores(tempVisited, rowLen, colLen, prev+1, i, j-1, data)
		}
	}

	if j < colLen-1 {
		shiftRightVal, _ := strconv.Atoi(data[i][j+1])
		if shiftRightVal == prev+1 {
			checkScores(tempVisited, rowLen, colLen, prev+1, i, j+1, data)
		}
	}

	if i > 0 {
		shiftUpVal, _ := strconv.Atoi(data[i-1][j])
		if shiftUpVal == prev+1 {
			checkScores(tempVisited, rowLen, colLen, prev+1, i-1, j, data)
		}
	}
	return

}

func readFile() ([][]string, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Slice to hold the 2D array as strings
	var array [][]string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Treat each digit in the row as an integer
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}

		// Append the row to the 2D array
		array = append(array, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return array, nil
}
