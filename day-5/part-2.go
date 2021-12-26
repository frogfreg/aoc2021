package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}
type matrix [][]int
type line []point

func main() {
	f, _ := os.Open("input-1.txt")
	defer f.Close()

	fs := bufio.NewScanner(f)

	lines := []line{}
	var max int

	for fs.Scan() {

		stringPoints := strings.Split(fs.Text(), "->")

		l := line{}

		for _, strPoint := range stringPoints {
			p := point{}

			xySlice := strings.Split(strPoint, ",")

			for i := range xySlice {
				xySlice[i] = strings.ReplaceAll(xySlice[i], " ", "")
			}

			p.x, _ = strconv.Atoi(xySlice[0])
			p.y, _ = strconv.Atoi(xySlice[1])

			if p.x > max {
				max = p.x
			}
			if p.y > max {
				max = p.y
			}

			l = append(l, p)
		}

		lines = append(lines, l)
	}

	oceanMap := make(matrix, max+1)

	for i := range oceanMap {
		oceanMap[i] = make([]int, max+1)
	}

	for _, line := range lines {
		a, b := line[0], line[1]
		if a.x == b.x || a.y == b.y {

			if a.x == b.x {
				minY, maxY := int(math.Min(float64(a.y), float64(b.y))), int(math.Max(float64(a.y), float64(b.y)))

				for i := minY; i <= maxY; i++ {
					oceanMap[i][a.x]++
				}
			}
			if a.y == b.y {
				minX, maxX := int(math.Min(float64(a.x), float64(b.x))), int(math.Max(float64(a.x), float64(b.x)))

				for i := minX; i <= maxX; i++ {
					oceanMap[a.y][i]++
				}
			}
		} else {
			i, j := a.y, a.x

			if a.x < b.x && a.y < b.y {

				for i <= b.y && j <= b.x {
					oceanMap[i][j]++
					i++
					j++
				}
			}
			if a.x > b.x && a.y < b.y {

				for i <= b.y && j >= b.x {
					oceanMap[i][j]++
					i++
					j--
				}
			}
			if a.x > b.x && a.y > b.y {

				for i >= b.y && j >= b.x {
					oceanMap[i][j]++
					i--
					j--
				}
			}
			if a.x < b.x && a.y > b.y {

				for i >= b.y && j <= b.x {
					oceanMap[i][j]++
					i--
					j++
				}
			}

		}
	}

	var count int

	for rowIndex := range oceanMap {
		for columnIndex := range oceanMap {
			if oceanMap[rowIndex][columnIndex] >= 2 {
				count++
			}
		}
	}

	fmt.Println(count)
}
