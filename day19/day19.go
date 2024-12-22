package day19

import (
	"adventofcode2024/common"
	"fmt"
	"strings"
)

type Onsen struct {
	towels   map[string]struct{}
	patterns []string
}

func parse(filename string) (result Onsen) {
	input, err := common.ReadInput("day19/" + filename)
	common.CheckError(err)

	firstLine := true

	result.towels = make(map[string]struct{})

	for line := range input.ReadLines {
		if firstLine {
			towels := strings.Split(line, ", ")

			for _, towel := range towels {
				result.towels[towel] = struct{}{}
			}

			firstLine = false

			continue
		}

		result.patterns = append(result.patterns, line)
	}

	return result
}

func combinations(pattern string, towels map[string]struct{}, memoize map[string]int64) int64 {
	if pattern == "" {
		return 1
	}

	if count, ok := memoize[pattern]; ok {
		return count
	}

	var count int64

	for idx := 1; idx <= len(pattern); idx++ {
		if _, ok := towels[pattern[:idx]]; ok {
			count += combinations(pattern[idx:], towels, memoize)
		}
	}

	memoize[pattern] = count

	return count
}

func Solve1(filename string) {
	result := parse(filename)

	var count int64

	memoize := make(map[string]int64)

	for _, pattern := range result.patterns {
		if combinations(pattern, result.towels, memoize) != 0 {
			count++
		}
	}

	fmt.Println("solution day 19 part 01:", count)
}

func Solve2(filename string) {
	result := parse(filename)

	var sum int64

	memoize := make(map[string]int64)

	for _, pattern := range result.patterns {
		sum += combinations(pattern, result.towels, memoize)
	}

	fmt.Println("solution day 19 part 02:", sum)
}
