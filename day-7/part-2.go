package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input-1.txt")
	crabPositionsString := strings.Split(string(data), ",")
	crabPositions := []int{}
	uniquePositionsMap := map[int]bool{}

	var maxPos, minPos int

	for _, digit := range crabPositionsString {
		digitInt, _ := strconv.Atoi(digit)
		if digitInt > maxPos {
			maxPos = digitInt
		}
		if digitInt < minPos {
			minPos = digitInt
		}
		crabPositions = append(crabPositions, digitInt)
	}

	for i := minPos; i <= maxPos; i++ {
		uniquePositionsMap[i] = true
	}

	var fuelCosts []int

	for uniqueCrabPos := range uniquePositionsMap {
		var fuelCost int
		for _, crabPos := range crabPositions {
			diff := int(math.Abs(float64(crabPos - uniqueCrabPos)))
			fuelCost += int((diff * (diff + 1)) / 2)
		}
		fuelCosts = append(fuelCosts, fuelCost)
	}

	sort.Ints(fuelCosts)

	fmt.Println(fuelCosts[0])
}
