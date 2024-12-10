package day02

import (
	"adventofcode2024/common"
	"fmt"
	"strconv"
	"strings"
)

func toInts(in []string) (out []int) {
	for _, level := range in {
		val, err := strconv.Atoi(level)
		common.CheckError(err)

		out = append(out, val)
	}

	return out
}

func goodLevel(isIncreasing bool, first, second int) bool {
	if first == second {
		return false
	}

	if isIncreasing {
		return second-first > 0 && second-first <= 3
	}

	return second-first >= -3 && second-first < 0
}

func Solve1(filename string) {
	input, err := common.ReadInput("day02/" + filename)
	common.CheckError(err)

	var safeCount int

reportLoop:
	for report := range input.ReadLines {
		levels := toInts(strings.Split(report, " "))

		if len(levels) <= 1 {
			safeCount++

			continue
		}

		isIncreasing := levels[1]-levels[0] > 0

		for i, val := range levels[:len(levels)-1] {
			if goodLevel(isIncreasing, val, levels[i+1]) {
				continue
			}

			continue reportLoop
		}

		safeCount++
	}

	fmt.Println("solution day 02 part 01:", safeCount)
}

func reportsToCheck(report []int) (result [][]int) {
	result = append(result, report)

	for i := range report {
		var current []int

		current = append(current, report[0:i]...)
		current = append(current, report[i+1:]...)

		result = append(result, current)
	}

	return result
}

func Solve2(filename string) {
	input, err := common.ReadInput("day02/" + filename)
	common.CheckError(err)

	var safeCount int

	for report := range input.ReadLines {
		levels := toInts(strings.Split(report, " "))

		if len(levels) <= 1 {
			safeCount++

			continue
		}

		allLevels := reportsToCheck(levels)

		var goodLevels int

	innerLoop:
		for _, innerLevels := range allLevels {
			isIncreasing := innerLevels[1]-innerLevels[0] > 0

			for i, val := range innerLevels[:len(innerLevels)-1] {
				if goodLevel(isIncreasing, val, innerLevels[i+1]) {
					continue
				}

				continue innerLoop
			}

			goodLevels++
		}

		if goodLevels > 0 {
			safeCount++
		}
	}

	fmt.Println("solution day 02 part 02:", safeCount)
}
