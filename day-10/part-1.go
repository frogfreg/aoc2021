package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, _ := os.Open("input-1.txt")

	defer f.Close()

	fs := bufio.NewScanner(f)

	var syntaxErrorScore int

	for fs.Scan() {
		char := findFirstIllegal(string(fs.Text()))

		switch char {
		case ")":
			syntaxErrorScore += 3
		case "]":
			syntaxErrorScore += 57
		case "}":
			syntaxErrorScore += 1197
		case ">":
			syntaxErrorScore += 25137
		}

	}
	fmt.Println(syntaxErrorScore)

}

func findFirstIllegal(line string) string {

	stack := []string{}

	for _, char := range line {

		charStr := string(char)

		switch charStr {
		case "(":
			stack = append(stack, ")")
		case "{":
			stack = append(stack, "}")
		case "[":
			stack = append(stack, "]")
		case "<":
			stack = append(stack, ">")
		case ")", "}", "]", ">":
			expected := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if charStr != expected {
				return charStr
			}
		}

	}
	return ""
}
