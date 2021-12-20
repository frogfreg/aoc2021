package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var horizontalPosition, aim, depth int

	f, err := os.Open("input.txt")

	if err == nil {
		bufferedReader := bufio.NewScanner(f)

		for bufferedReader.Scan() {
			instructions := strings.Fields(bufferedReader.Text())
			direction := instructions[0]
			amount, _ := strconv.Atoi(instructions[1])

			if direction == "down" {
				aim += amount
			}
			if direction == "up" {
				aim -= amount
			}
			if direction == "forward" {
				horizontalPosition += amount
				depth += aim * amount
			}
		}

		fmt.Println(horizontalPosition * depth)
	}
}
