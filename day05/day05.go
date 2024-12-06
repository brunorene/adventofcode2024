package day05

import (
	"adventofcode2024/common"
	"fmt"
	"strconv"
	"strings"
)

func Solve1(filename string) {
	before, updates := parse(filename)

	var sum int

	for _, update := range updates {
		if isOrderCorrect(update, before) {
			sum += update[len(update)/2]
		}
	}

	fmt.Println("solution day 05 part 01:", sum)
}

func Solve2(filename string) {
	before, updates := parse(filename)

	var sum int

	for _, update := range updates {
		if !isOrderCorrect(update, before) {
			update = fix(update, before)

			sum += update[len(update)/2]
		}
	}

	fmt.Println("solution day 05 part 02:", sum)
}

func fix(update []int, before map[int]map[int]struct{}) []int {
	for !isOrderCorrect(update, before) {
		firstIdx, secondIdx := wrongPair(update, before)

		first := update[firstIdx]
		update[firstIdx] = update[secondIdx]
		update[secondIdx] = first
	}

	return update
}

func wrongPair(update []int, before map[int]map[int]struct{}) (int, int) {
	for idx, page := range update {
		for i := idx + 1; i < len(update); i++ {
			if _, exists := before[page][update[i]]; exists {
				return idx, i
			}
		}
	}

	return -1, -1
}

func isOrderCorrect(update []int, before map[int]map[int]struct{}) bool {
	for idx, page := range update {
		for i := idx + 1; i < len(update); i++ {
			if _, exists := before[page][update[i]]; exists {
				return false
			}
		}
	}

	return true
}

func parse(filename string) (map[int]map[int]struct{}, [][]int) {
	input, err := common.ReadInput("day05/" + filename)
	common.CheckError(err)

	before := make(map[int]map[int]struct{})
	var updates [][]int
	var parseUpdates bool

	for line := range input.ReadLines {
		if strings.Index(line, ",") > 0 {
			parseUpdates = true
		}

		if parseUpdates {
			pages := strings.Split(line, ",")

			var update []int

			for _, page := range pages {
				value, err := strconv.Atoi(page)
				common.CheckError(err)

				update = append(update, value)
			}

			updates = append(updates, update)

			continue
		}

		pages := strings.Split(line, "|")

		p0, err := strconv.Atoi(pages[0])
		common.CheckError(err)

		p1, err := strconv.Atoi(pages[1])
		common.CheckError(err)

		if before[p1] == nil {
			before[p1] = make(map[int]struct{})
		}

		before[p1][p0] = struct{}{}
	}

	return before, updates
}
