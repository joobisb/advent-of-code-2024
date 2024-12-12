package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var score int
var ratings int

func main() {

	data, _ := readFile()

	rowLen := len(data)
	colLen := len(data[0])

	for i, row := range data {
		for j, _ := range row {
			if data[i][j] == "0" {
				//using this to mark already visited 9s
				tempVisited := make(map[string]bool)
				checkScores(tempVisited, rowLen, colLen, 0, i, j, data)
				checkRatings(tempVisited, rowLen, colLen, 0, i, j, data)
			}
		}
	}

	fmt.Printf("score: %d\n", score)
	fmt.Printf("rating: %d\n", ratings)
}

func checkScores(tempVisited map[string]bool, rowLen, colLen, prev, i, j int, data [][]string) {
	if prev == 9 {
		if !tempVisited[strconv.Itoa(i)+strconv.Itoa(j)] {
			score++
		}
		tempVisited[strconv.Itoa(i)+strconv.Itoa(j)] = true
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

}

// this is basically the same as checkScores only diff being calc. ratings
// no need to check if its being already visited or not
func checkRatings(tempVisited map[string]bool, rowLen, colLen, prev, i, j int, data [][]string) {
	if prev == 9 {
		ratings++
		return
	}
	if i < 0 || j < 0 || i > rowLen || j > colLen {
		if prev == 9 {
			ratings++
		}
		return
	}

	if i < rowLen-1 {
		shiftDownVal, _ := strconv.Atoi(data[i+1][j])
		if shiftDownVal == prev+1 {
			checkRatings(tempVisited, rowLen, colLen, prev+1, i+1, j, data)
		}
	}
	if j > 0 {
		shiftLeftVal, _ := strconv.Atoi(data[i][j-1])
		if shiftLeftVal == prev+1 {
			checkRatings(tempVisited, rowLen, colLen, prev+1, i, j-1, data)
		}
	}

	if j < colLen-1 {
		shiftRightVal, _ := strconv.Atoi(data[i][j+1])
		if shiftRightVal == prev+1 {
			checkRatings(tempVisited, rowLen, colLen, prev+1, i, j+1, data)
		}
	}

	if i > 0 {
		shiftUpVal, _ := strconv.Atoi(data[i-1][j])
		if shiftUpVal == prev+1 {
			checkRatings(tempVisited, rowLen, colLen, prev+1, i-1, j, data)
		}
	}

}

func readFile() ([][]string, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var array [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}

		array = append(array, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return array, nil
}
