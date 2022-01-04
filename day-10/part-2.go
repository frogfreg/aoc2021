package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var charsClosingMap = map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
var charsValue = map[string]int{")": 1, "]": 2, "}": 3, ">": 4}

func main() {

	f, _ := os.Open("input-1.txt")

	defer f.Close()

	fs := bufio.NewScanner(f)

	var scores []int

	for fs.Scan() {
		charsMap := map[string]int{}

		filteredChars := removeValidChunks(string(fs.Text()))

		for _, char := range filteredChars {
			charsMap[string(char)]++
		}

		var closingCount int

		for _, char := range ")}]>" {
			closingCount += charsMap[string(char)]
		}

		if closingCount == 0 {
			complement := []string{}

			for _, char := range filteredChars {
				charStr := string(char)
				complement = append(complement, charsClosingMap[charStr])
			}
			var singleLineScore int
			for i := len(complement) - 1; i >= 0; i-- {
				singleLineScore *= 5
				singleLineScore += charsValue[complement[i]]
			}
			scores = append(scores, singleLineScore)
		}

	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}

func removeValidChunks(line string) string {
	chars := strings.Split(line, "")
	charIndex := 0

	for charIndex < len(chars)-1 {
		if charIndex < len(chars)-1 && chars[charIndex+1] == charsClosingMap[chars[charIndex]] {
			if charIndex == len(chars)-2 {
				chars = chars[:charIndex]
			} else {
				chars = append(chars[:charIndex], chars[charIndex+2:]...)
			}
			charIndex = 0
			continue
		}
		charIndex++
	}

	return strings.Join(chars, "")
}
