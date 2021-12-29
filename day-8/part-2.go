package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input-1.txt")

	defer f.Close()

	fs := bufio.NewScanner(f)

	var total int

	for fs.Scan() {
		entry := strings.Split(fs.Text(), "|")
		patternString, outputString := entry[0], entry[1]

		numMap := map[string]string{}

		ocurrenceMap := map[string]int{}
		patterns := strings.Fields(patternString)

		for patternIndex, digitString := range patterns {
			letters := strings.Split(digitString, "")
			sort.Strings(letters)

			sortedDigitString := strings.Join(letters, "")

			patterns[patternIndex] = sortedDigitString

			if len(sortedDigitString) == 2 {
				numMap["1"] = sortedDigitString
			}
			if len(sortedDigitString) == 3 {
				numMap["7"] = sortedDigitString
			}
			if len(sortedDigitString) == 4 {
				numMap["4"] = sortedDigitString
			}
			if len(sortedDigitString) == 7 {
				numMap["8"] = sortedDigitString
			}
			for _, letter := range letters {
				ocurrenceMap[letter]++
			}
		}
		numMap["0"], numMap["2"], numMap["3"], numMap["5"], numMap["6"], numMap["9"] = findOtherNums(patterns, ocurrenceMap, numMap)

		var trueDigitString string

		for _, outputLetters := range strings.Fields(outputString) {
			letters := strings.Split(outputLetters, "")
			sort.Strings(letters)

			sortedOutputLetters := strings.Join(letters, "")
			for num, letters := range numMap {
				if sortedOutputLetters == letters {

					trueDigitString += num
				}
			}

		}
		trueDigitInt, _ := strconv.Atoi(trueDigitString)

		total += trueDigitInt
	}

	fmt.Println(total)

}

func findOtherNums(patterns []string, ocurrenceMap map[string]int, numMap map[string]string) (string, string, string, string, string, string) {
	var notIncludedTwo, notIncludedThree, notIncludedNine string
	var zeroHints []string

	var zero, two, three, five, six, nine string

	for letter, count := range ocurrenceMap {
		if count == 9 {
			notIncludedTwo = letter
		}
		if count == 6 {
			notIncludedThree = letter
		}
		if count == 4 {
			notIncludedNine = letter
		}
		if count == 7 {
			zeroHints = append(zeroHints, letter)
		}
	}

	for _, pat := range patterns {

		isZero := len(pat) == 6 && strings.Contains(pat, notIncludedNine) && strings.Contains(pat, string(numMap["1"][0])) && strings.Contains(pat, string(numMap["1"][1]))

		if !strings.Contains(pat, notIncludedTwo) {
			two = pat
		}
		if len(pat) == 5 && !strings.Contains(pat, notIncludedThree) && (strings.Contains(pat, string(numMap["1"][0])) && strings.Contains(pat, string(numMap["1"][1]))) {
			three = pat
		}
		if len(pat) == 6 && !strings.Contains(pat, notIncludedNine) {
			nine = pat
		}
		if isZero {
			zero = pat
		}
		if len(pat) == 5 && strings.Contains(pat, notIncludedThree) {
			five = pat
		}
		if len(pat) == 6 && pat != zero && pat != nine {
			six = pat
		}
	}
	return zero, two, three, five, six, nine
}
