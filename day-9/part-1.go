package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type matrix [][]int

func main() {

	f, _ := os.Open("input-1.txt")
	defer f.Close()

	fs := bufio.NewScanner(f)
	var smokeMatrix matrix

	for fs.Scan() {
		numSlice := []int{}

		for _, digit := range string(fs.Text()) {
			digitInt, _ := strconv.Atoi(string(digit))

			numSlice = append(numSlice, digitInt)
		}
		smokeMatrix = append(smokeMatrix, numSlice)
	}

	var riskSum int

	for rowIndex := 0; rowIndex < len(smokeMatrix); rowIndex++ {
		for colIndex := 0; colIndex < len(smokeMatrix[0]); colIndex++ {
			currValue := smokeMatrix[rowIndex][colIndex]

			var upValue, leftValue, rightValue, downValue int

			if rowIndex == 0 {
				upValue = currValue + 1
			} else {
				upValue = smokeMatrix[rowIndex-1][colIndex]
			}
			if rowIndex == len(smokeMatrix)-1 {
				downValue = currValue + 1
			} else {
				downValue = smokeMatrix[rowIndex+1][colIndex]
			}
			if colIndex == 0 {
				leftValue = currValue + 1
			} else {
				leftValue = smokeMatrix[rowIndex][colIndex-1]
			}
			if colIndex == len(smokeMatrix[0])-1 {
				rightValue = currValue + 1
			} else {
				rightValue = smokeMatrix[rowIndex][colIndex+1]
			}

			if currValue < upValue && currValue < downValue && currValue < leftValue && currValue < rightValue {
				riskSum += currValue + 1
			}
		}
	}

	fmt.Println(riskSum)
}
