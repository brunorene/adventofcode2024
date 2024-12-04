package day04

import (
	"adventofcode2024/common"
	"fmt"
)

type matrix [][]byte

func (m matrix) find1(row, col int, str string) (count int) {
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
					currentRow < 0 || currentCol < 0 {
					continue outerLoop
				}

				if m[currentRow][currentCol] != str[currentStr] {
					continue outerLoop
				}

				currentRow += vert
				currentCol += horz
				currentStr++

				if currentStr == len(str) {
					count++

					break
				}
			}
		}
	}

	return count
}

type format int

const (
	top    format = iota
	bottom format = iota
	left   format = iota
	right  format = iota
)

func (m matrix) find2(row, col int) (count int) {
	if m[row][col] != 'A' || row == 0 || row == len(m)-1 || col == 0 || col == len(m[row])-1 {
		return 0
	}

	for _, form := range []format{top, bottom, left, right} {
		switch form {
		case top:
			if m[row-1][col-1] == 'M' && m[row-1][col+1] == 'M' &&
				m[row+1][col-1] == 'S' && m[row+1][col+1] == 'S' {
				count++
			}
		case bottom:
			if m[row+1][col-1] == 'M' && m[row+1][col+1] == 'M' &&
				m[row-1][col-1] == 'S' && m[row-1][col+1] == 'S' {
				count++
			}
		case left:
			if m[row-1][col-1] == 'M' && m[row+1][col-1] == 'M' &&
				m[row-1][col+1] == 'S' && m[row+1][col+1] == 'S' {
				count++
			}
		case right:
			if m[row-1][col+1] == 'M' && m[row+1][col+1] == 'M' &&
				m[row-1][col-1] == 'S' && m[row+1][col-1] == 'S' {
				count++
			}
		}
	}

	return count
}

func Solve1(filename string) {
	input, err := common.ReadInput("day04/" + filename)
	if err != nil {
		panic(err.Error())
	}

	var count int

	matrix := make(matrix, 0, 140)

	for line := range input.ReadLines {
		matrix = append(matrix, []byte(line))
	}

	for row := range matrix {
		for col := range matrix[row] {
			count += matrix.find1(row, col, "XMAS")
		}
	}

	fmt.Println("solution day 04 part 01:", count)
}

func Solve2(filename string) {
	input, err := common.ReadInput("day04/" + filename)
	if err != nil {
		panic(err.Error())
	}

	var count int

	matrix := make(matrix, 0, 140)

	for line := range input.ReadLines {
		matrix = append(matrix, []byte(line))
	}

	for row := range matrix {
		for col := range matrix[row] {
			count += matrix.find2(row, col)
		}
	}

	fmt.Println("solution day 04 part 02:", count)
}
