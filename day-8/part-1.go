package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input-1.txt")

	defer f.Close()

	fs := bufio.NewScanner(f)

	var specialDigitCount int

	for fs.Scan() {
		entry := strings.Split(fs.Text(), "|")

		patternString, outputString := entry[0], entry[1]

		_ = patternString
		for _, digitString := range strings.Fields(outputString) {
			if len(digitString) == 2 || len(digitString) == 7 || len(digitString) == 4 || len(digitString) == 3 {
				specialDigitCount++
			}
		}
	}

	fmt.Println(specialDigitCount)
}
