package day04

import (
	"adventofcode2024/common"
	"fmt"
)

type matrix [][]byte

func find1(m matrix, row, col int) (count int) {
	xmas := "XMAS"

	for _, horz := range []int{-1, 0, 1} {
	outerLoop:
		for _, vert := range []int{-1, 0, 1} {
			if horz == 0 && vert == 0 {
				continue
			}

			currentRow := row
			currentCol := col
			currentStr := 0

			for {
				if currentRow >= len(m) || currentCol >= len(m[row]) ||
					currentRow < 0 || currentCol < 0 || m[currentRow][currentCol] != xmas[currentStr] {
					continue outerLoop
				}

				currentRow += vert
				currentCol += horz
				currentStr++

				if currentStr == len(xmas) {
					count++

					break
				}
			}
		}
	}

	return count
}

type format struct {
	topLeft     byte
	topRight    byte
	bottomLeft  byte
	bottomRight byte
}

func find2(m matrix, row, col int) (count int) {
	if m[row][col] != 'A' || row == 0 || row == len(m)-1 || col == 0 || col == len(m[row])-1 {
		return 0
	}

	top := format{'M', 'M', 'S', 'S'}
	bottom := format{'S', 'S', 'M', 'M'}
	left := format{'M', 'S', 'M', 'S'}
	right := format{'S', 'M', 'S', 'M'}

	for _, form := range []format{top, bottom, left, right} {
		if m[row-1][col-1] == form.topLeft && m[row-1][col+1] == form.topRight &&
			m[row+1][col-1] == form.bottomLeft && m[row+1][col+1] == form.bottomRight {
			return 1
		}
	}

	return 0
}

func Solve(filename, part string, find func(matrix, int, int) int) {
	input, err := common.ReadInput("day04/" + filename)
	common.CheckError(err)

	var count int

	matrix := make(matrix, 0, 140)

	for line := range input.ReadLines {
		matrix = append(matrix, []byte(line))
	}

	for row := range matrix {
		for col := range matrix[row] {
			count += find(matrix, row, col)
		}
	}

	fmt.Printf("solution day 04 part 0%s: %d\n", part, count)
}

func Solve1(filename string) {
	Solve(filename, "1", find1)
}

func Solve2(filename string) {
	Solve(filename, "1", find2)
}
