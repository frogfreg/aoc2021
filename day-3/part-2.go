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

func filter(s []string, index int, letter string) []string {

	temp := []string{}

	for _, value := range s {
		if string(value[index]) == letter {
			temp = append(temp, value)
		}
	}

	return temp
}

func main() {

	f, _ := os.Open("input-1.txt")

	countsSlice := []counts{}
	numberSlice := []string{}

	bufferedReader := bufio.NewScanner(f)

	for bufferedReader.Scan() {

		numberSlice = append(numberSlice, bufferedReader.Text())

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

	filteredNumbersOxygen := numberSlice[:]
	filteredNumbersC := numberSlice[:]

	i := 0

	for len(filteredNumbersOxygen) != 1 && i <= len(countsSlice)-1 {
		var mostCommon string

		if countsSlice[i].one >= countsSlice[i].zero {
			mostCommon = "1"
		} else {
			mostCommon = "0"
		}

		filteredNumbersOxygen = filter(filteredNumbersOxygen, i, mostCommon)

		countsSlice = make([]counts, len(filteredNumbersOxygen[0]))

		for _, num := range filteredNumbersOxygen {
			for index, letter := range num {
				if string(letter) == "1" {
					countsSlice[index].one++
				} else {
					countsSlice[index].zero++
				}
			}
		}
		i++
	}
	i = 0
	for len(filteredNumbersC) != 1 && i <= len(countsSlice)-1 {
		var leastCommon string

		if countsSlice[i].one >= countsSlice[i].zero {
			leastCommon = "0"
		} else {
			leastCommon = "1"
		}

		filteredNumbersC = filter(filteredNumbersC, i, leastCommon)

		countsSlice = make([]counts, len(filteredNumbersC[0]))

		for _, num := range filteredNumbersC {
			for index, letter := range num {
				if string(letter) == "1" {
					countsSlice[index].one++
				} else {
					countsSlice[index].zero++
				}
			}
		}
		i++
	}

	oxy, _ := strconv.ParseInt(filteredNumbersOxygen[0], 2, 64)
	c, _ := strconv.ParseInt(filteredNumbersC[0], 2, 64)

	fmt.Println(oxy * c)

}
