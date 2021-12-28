package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("input-1.txt")
	var lanternFish []int

	for _, digit := range string(data) {
		digitInt, err := strconv.Atoi(string(digit))
		if err == nil {
			lanternFish = append(lanternFish, digitInt)
		}
	}

	for i := 0; i < 80; i++ {
		var count int
		currentLength := len(lanternFish)

		for li := 0; li < currentLength; li++ {
			lanternFish[li]--
			if lanternFish[li] < 0 {
				count++
				lanternFish[li] = 6
			}
		}
		for j := 0; j < count; j++ {
			lanternFish = append(lanternFish, 8)
		}
	}

	fmt.Printf("%#v\n", len(lanternFish))

}
