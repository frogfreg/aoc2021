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

	for _, digit := range crabPositionsString {
		digitInt, _ := strconv.Atoi(digit)
		_, exists := uniquePositionsMap[digitInt]
		if !exists {
			uniquePositionsMap[digitInt] = true
		}
		crabPositions = append(crabPositions, digitInt)
	}

	var fuelCosts []int

	for uniqueCrabPos := range uniquePositionsMap {
		var fuelCost int
		for _, crabPos := range crabPositions {
			fuelCost += int(math.Abs(float64(crabPos - uniqueCrabPos)))
		}
		fuelCosts = append(fuelCosts, fuelCost)

	}

	sort.Ints(fuelCosts)

	fmt.Println(fuelCosts[0])

}
