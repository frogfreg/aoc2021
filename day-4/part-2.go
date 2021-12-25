package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrix [][]string

func main() {
	f, _ := os.Open("input-1.txt")
	defer f.Close()

	bs := bufio.NewScanner(f)

	bs.Scan()
	selectedNums := strings.Split(bs.Text(), ",")

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

	boardMap := map[int]matrix{}
	lastWinnerIndex := 0
	var winnerNum string

	for _, num := range selectedNums {
		for boardIndex, board := range boards {
			for rowIndex, row := range board {
				for colIndex, columnNumber := range row {
					if num == columnNumber {
						boards[boardIndex][rowIndex][colIndex] = "m"
						if checkWin(board) {

							_, keyExists := boardMap[boardIndex]

							if !keyExists {

								tempBoard := make(matrix, len(board))

								for i, _ := range tempBoard {
									tempBoard[i] = make([]string, len(board[0]))
								}

								for i := 0; i < len(board); i++ {
									for j := 0; j < len(board[1]); j++ {
										tempBoard[i][j] = board[i][j]
									}
								}

								lastWinnerIndex = boardIndex
								winnerNum = num
								boardMap[boardIndex] = tempBoard
							}
						}
					}
				}
			}
		}
	}

	winnerNumInt, _ := strconv.Atoi(winnerNum)
	result := sumBoard(boardMap[lastWinnerIndex]) * winnerNumInt

	fmt.Println(result)

}

func checkWin(m matrix) bool {
	for _, row := range m {
		if strings.Join(row, "") == "mmmmm" {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		count := 0

		for j := 0; j < 5; j++ {
			if string(m[j][i]) == "m" {
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
