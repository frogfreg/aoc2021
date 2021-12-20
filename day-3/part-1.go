package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type counts struct {
	one, zero int
}

func main() {

	f, _ := os.Open("input-1.txt")

	countsSlice := []counts{}

	bufferedReader := bufio.NewScanner(f)

	for bufferedReader.Scan() {

		if len(countsSlice) == 0 {
			countsSlice = make([]counts, len(bufferedReader.Text()))
		}

		for index, letter := range bufferedReader.Text() {

			if string(letter) == "1" {
				countsSlice[index].one++
			} else {
				countsSlice[index].zero++
			}
		}
	}

	gamma := ""
	epsilon := ""

	for _, count := range countsSlice {
		if count.one > count.zero {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaNum, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonNum, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gammaNum * epsilonNum)

}
