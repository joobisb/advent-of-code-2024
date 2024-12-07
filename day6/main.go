package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	data, err := readFile()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rowLen := len(data)
	colLen := len(data[0])
	var count int
	for i, row := range data {
		for j, _ := range row {
			if data[i][j] == "^" {
				count = 1
				data[i][j] = "X"
			iterate:
				for i >= 0 {
					if i == 0 {
						goto output
					}
					i--
					if data[i][j] != "#" {
						if data[i][j] != "X" {
							count++
						}
						data[i][j] = "X"

					} else {
						i++
						for j < rowLen {
							if j == rowLen-1 {
								goto output
							}
							j++
							if data[i][j] != "#" {
								if data[i][j] != "X" {
									count++
								}
								data[i][j] = "X"
							} else {
								j--
								for i < colLen {
									if i == colLen-1 {
										goto output
									}
									i++
									if data[i][j] != "#" {
										if data[i][j] != "X" {
											count++
										}
										data[i][j] = "X"
									} else {
										i--
										for j >= 0 {
											if j == 0 {
												goto output
											}
											j--
											if data[i][j] != "#" {
												if data[i][j] != "X" {
													count++
												}
												data[i][j] = "X"
											} else {
												j++
												goto iterate
											}
										}
									}
								}
							}
						}

					}
				}

			}
		}
	}

output:
	fmt.Printf("count: %d\n", count)

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
