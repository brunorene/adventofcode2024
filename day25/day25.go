package day25

import (
	"adventofcode2024/common"
	"fmt"
)

type columns [5]int

type installation struct {
	locks []columns
	keys  []columns
}

func parse(filename string) (result installation) {
	content, err := common.ReadInput("day25/" + filename)
	common.CheckError(err)

	var row int

	var current columns
	var isLock bool

	for line := range content.ReadLines {
		if row == 0 {
			isLock = line == "#####"

			row = (row + 1) % 7

			continue
		}

		if isLock || row < 6 {
			for i, char := range line {
				if char == '#' {
					current[i]++
				}
			}
		}

		if row == 6 {
			if isLock {
				result.locks = append(result.locks, current)
			} else {
				result.keys = append(result.keys, current)
			}

			current = columns{}
		}

		row = (row + 1) % 7
	}

	return result
}

func Solve1(filename string) {
	result := parse(filename)

	var fitCount int

	for _, key := range result.keys {
	lockLoop:
		for _, lock := range result.locks {
			for idx := range key {
				if key[idx]+lock[idx] > 5 {
					continue lockLoop
				}
			}

			fitCount++
		}
	}

	fmt.Println("solution day 25 part 01:", fitCount)
}
