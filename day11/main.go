package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, _ := readFile()
	fmt.Println("data", data)
	blink(data, 0)
}

func blink(data []int, blinkCount int) {
	if blinkCount == 25 {
		fmt.Printf("stones after 25 blinks: %d \n", len(data))
		return
	}
	var blinkedArray []int
	var blinkedArrayIdx = 0
	for _, v := range data {
		if v == 0 {
			blinkedArray = append(blinkedArray, 1)
			blinkedArrayIdx++
			continue
		}
		noOfDigits := countDigits(v)
		if noOfDigits%2 == 0 {
			var temArr []int
			//find the digits and put it into a tempArr
			for v > 0 {
				rem := v % 10
				temArr = append(temArr, rem)
				v = v / 10
			}
			//split the temp array into 2 half
			//by concatenation the left and right parts of the array
			var tempString string
			for j := len(temArr) - 1; j >= 0; j-- {

				if j >= len(temArr)/2 {

					tempString = tempString + strconv.Itoa(temArr[j])
				}
				if j == len(temArr)/2 {
					tempStrToInt, _ := strconv.Atoi(tempString)
					blinkedArray = append(blinkedArray, tempStrToInt)
					tempString = ""
				}
				if j >= 0 && j < len(temArr)/2 {
					tempString = tempString + strconv.Itoa(temArr[j])
				}
				if j == 0 {
					tempStrToInt, _ := strconv.Atoi(tempString)
					blinkedArray = append(blinkedArray, tempStrToInt)
					tempString = ""
				}

			}

			blinkedArrayIdx = blinkedArrayIdx + len(temArr) + 1
			continue
		}
		blinkedArray = append(blinkedArray, v*2024)
		blinkedArrayIdx++
	}
	blink(blinkedArray, blinkCount+1)
}

func countDigits(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}
func readFile() ([]int, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return nil, err
			}
			numbers = append(numbers, num)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return numbers, nil

}
