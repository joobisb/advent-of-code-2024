package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	data, err := readFile()
	if err != nil {
		return
	}

	var count int
	//this was unncessary, misunderstood the problem initially, anyway went with this
	orderedArray := []string{"X", "M", "A", "S"}
	backwardsArray := []string{"S", "A", "M", "X"}

	//iterate horizontally
	for i, row := range data {
		k := 0
		for j, _ := range row {

			if data[i][j] == orderedArray[k] {
				tempj := j
				for k < len(orderedArray) && tempj < len(row) {
					if data[i][tempj] != orderedArray[k] {
						break
					}
					tempj++
					k++
				}
				if k == len(orderedArray) {
					count++
				}
				k = 0
			}

			if data[i][j] == backwardsArray[k] {
				tempj := j
				for k < len(backwardsArray) && tempj < len(row) {
					if data[i][tempj] != backwardsArray[k] {
						break
					}
					tempj++
					k++
				}
				if k == len(backwardsArray) {
					count++
				}
				k = 0
			}
		}
	}

	rowLen := len(data)
	colLen := len(data[0])

	//iterate vertically
	for col := 0; col < colLen; col++ {
		k := 0
		for row := 0; row < rowLen; row++ {

			if data[row][col] == orderedArray[k] {
				tempRow := row
				for k < len(orderedArray) && tempRow < rowLen {
					if data[tempRow][col] != orderedArray[k] {
						break
					}
					tempRow++
					k++
				}
				if k == len(orderedArray) {
					count++
				}
				k = 0
			}

			if data[row][col] == backwardsArray[k] {
				tempRow := row
				for k < len(backwardsArray) && tempRow < rowLen {
					if data[tempRow][col] != backwardsArray[k] {
						break
					}
					tempRow++
					k++
				}
				if k == len(backwardsArray) {
					count++
				}
				k = 0
			}
		}
	}

	// diagonal
	// kind of hacky, but works
	//spend a lot of time here with a diff solution :-
	for i, row := range data {
		k := 0
		for j, _ := range row {
			topLeftI := i
			topLeftJ := j
			for {
				if (topLeftI >= 0 && topLeftJ >= 0) && k < len(orderedArray) {
					if data[topLeftI][topLeftJ] == orderedArray[k] {
						k++
					} else {
						break
					}
					topLeftJ--
					topLeftI--
				} else {
					break
				}
			}
			if k == len(orderedArray) {
				count++

			}
			k = 0
			topRightI := i
			topRightJ := j

			for {
				if topRightI >= 0 && topRightJ < rowLen && k < len(orderedArray) {
					if data[topRightI][topRightJ] == orderedArray[k] {
						k++
					} else {
						break
					}
					topRightI--
					topRightJ++
				} else {
					break
				}
			}
			if k == len(orderedArray) {
				count++
			}
			k = 0
			bottomRightI := i
			bottomRightJ := j

			for {
				if bottomRightI < colLen && bottomRightJ < rowLen && k < len(orderedArray) {
					if data[bottomRightI][bottomRightJ] == orderedArray[k] {
						k++
					} else {
						break
					}
					bottomRightI++
					bottomRightJ++
				} else {
					break
				}
			}
			if k == len(orderedArray) {
				count++
			}
			k = 0
			bottomLeftI := i
			bottomLeftJ := j

			for {
				if bottomLeftI < colLen && bottomLeftJ >= 0 && k < len(orderedArray) {
					if data[bottomLeftI][bottomLeftJ] == orderedArray[k] {
						k++
					} else {
						break
					}
					bottomLeftI++
					bottomLeftJ--
				} else {
					break
				}
			}
			if k == len(orderedArray) {
				count++
			}
			k = 0

		}
	}
	fmt.Printf("count: %d", count)
}

func readFile() ([][]string, error) {
	filePath := "data.txt"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil

}