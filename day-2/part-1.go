package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var horizontalPosition, depth int

	f, err := os.Open("input.txt")

	if err == nil {
		bufferedReader := bufio.NewScanner(f)

		for bufferedReader.Scan() {
			instructions := strings.Fields(bufferedReader.Text())
			direction := instructions[0]
			amount, _ := strconv.Atoi(instructions[1])

			if direction == "forward" {
				horizontalPosition += amount
			}
			if direction == "down" {
				depth += amount
			}
			if direction == "up" {
				depth -= amount
			}
		}

		fmt.Println(horizontalPosition * depth)
	}

}
