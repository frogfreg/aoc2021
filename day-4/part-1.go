package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrix [][]string

func checkWin(m matrix) bool {
	for _, row := range m {
		if strings.Join(row, "") == "mmmmm" {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		count := 0

		for j := 0; j < 5; j++ {
			if string(m[i][j]) == "m" {
				count++
			}
		}
		if count >= 5 {
			return true
		}
	}
	return false
}

func sumBoard(m matrix) int {
	sum := 0

	for _, row := range m {
		for _, digit := range row {
			digitInt, err := strconv.Atoi(digit)
			if err == nil {
				sum += digitInt
			}
		}
	}

	return sum
}

func main() {
	f, _ := os.Open("input-1.txt")
	defer f.Close()

	bs := bufio.NewScanner(f)

	bs.Scan()
	selectedNums := strings.Split(bs.Text(), ",")
	// selectedNums := []string{"7", "4", "9", "5"}

	boards := []matrix{}

	var m matrix

	for bs.Scan() {

		row := strings.Split(bs.Text(), " ")

		for i, element := range row {
			if element == "" {
				row = append(row[:i], row[i+1:]...)
			}
		}

		if len(row) < 5 {
			continue
		}

		m = append(m, row)
		if len(m) >= 5 {
			boards = append(boards, m)
			m = matrix{}
		}

	}

	func() {
		for _, num := range selectedNums {
			for boardIndex, board := range boards {
				for rowIndex, row := range board {
					for colIndex, columnNumber := range row {
						if num == columnNumber {
							boards[boardIndex][rowIndex][colIndex] = "m"
							if checkWin(board) {
								selectedNumInt, _ := strconv.Atoi(num)

								fmt.Println(sumBoard(board) * selectedNumInt)
								return
							}
						}
					}
				}
			}
		}
	}()

}
