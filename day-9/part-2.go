package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type matrix [][]int
type coord struct {
	row, col int
}

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

	var lowPoints []coord

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
				lowPoints = append(lowPoints, coord{rowIndex, colIndex})
			}
		}
	}

	var basinSizes []int

	for _, point := range lowPoints {

		currentBasin := []coord{point}

		for pi := 0; pi < len(currentBasin); pi++ {
			basinPoint := currentBasin[pi]
			ri, ci := basinPoint.row, basinPoint.col
			smokeMatrix[ri][ci] = -1

			if ri > 0 && smokeMatrix[ri-1][ci] > smokeMatrix[ri][ci] && smokeMatrix[ri-1][ci] < 9 && smokeMatrix[ri-1][ci] >= 0 {
				currentBasin = append(currentBasin, coord{ri - 1, ci})
				smokeMatrix[ri-1][ci] = -1
			}
			if ri < len(smokeMatrix)-1 && smokeMatrix[ri+1][ci] > smokeMatrix[ri][ci] && smokeMatrix[ri+1][ci] < 9 && smokeMatrix[ri+1][ci] >= 0 {
				currentBasin = append(currentBasin, coord{ri + 1, ci})
				smokeMatrix[ri+1][ci] = -1

			}
			if ci > 0 && smokeMatrix[ri][ci-1] > smokeMatrix[ri][ci] && smokeMatrix[ri][ci-1] < 9 && smokeMatrix[ri][ci-1] >= 0 {
				currentBasin = append(currentBasin, coord{ri, ci - 1})
				smokeMatrix[ri][ci-1] = -1

			}
			if ci < len(smokeMatrix[0])-1 && smokeMatrix[ri][ci+1] > smokeMatrix[ri][ci] && smokeMatrix[ri][ci+1] < 9 && smokeMatrix[ri][ci+1] >= 0 {
				currentBasin = append(currentBasin, coord{ri, ci + 1})
				smokeMatrix[ri][ci+1] = -1

			}
		}
		basinSizes = append(basinSizes, len(currentBasin))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])

}
