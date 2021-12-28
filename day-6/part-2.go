package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input-1.txt")
	fishMap := map[int]int{}
	for i := 0; i <= 8; i++ {
		fishMap[i] = 0
	}

	initialFish := strings.Split(string(data), ",")

	for _, digit := range initialFish {
		digitInt, _ := strconv.Atoi(digit)
		fishMap[digitInt]++
	}

	for i := 0; i < 256; i++ {
		increaseMap := map[int]int{}
		for daysIndex := 8; daysIndex >= 0; daysIndex-- {
			if fishMap[daysIndex] > 0 {
				if daysIndex == 0 {
					increaseMap[8] += fishMap[0]
					increaseMap[6] += fishMap[0]
					fishMap[0] = 0
					continue
				}
				increaseMap[daysIndex-1] += fishMap[daysIndex]
				fishMap[daysIndex] = 0
			}
		}
		for k, v := range increaseMap {
			fishMap[k] += v
		}
	}

	var totalFish int

	for _, fishCount := range fishMap {
		totalFish += fishCount
	}

	fmt.Println(totalFish)

}
