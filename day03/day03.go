package day03

import (
	"adventofcode2024/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func numbers(mul string) (left, right int64, err error) {
	result := strings.ReplaceAll(mul, "mul(", "")
	result = strings.ReplaceAll(result, ")", "")

	parts := strings.Split(result, ",")

	leftInt, err := strconv.Atoi(parts[0])
	if err != nil {
		return -1, -1, fmt.Errorf("left convert: %w", err)
	}

	rightInt, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, -1, fmt.Errorf("right convert: %w", err)
	}

	return int64(leftInt), int64(rightInt), nil
}

func Solve1(filename string) {
	input, err := common.ReadInput("day03/" + filename)
	common.CheckError(err)

	line := input.Read()

	matcher := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	var result int64

	for _, mul := range matcher.FindAllString(line, -1) {
		left, right, err := numbers(mul)
		common.CheckError(err)

		result += left * right
	}

	fmt.Println("solution day 03 part 01:", result)
}

func Solve2(filename string) {
	input, err := common.ReadInput("day03/" + filename)
	common.CheckError(err)

	line := input.Read()

	matcher := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matcherDos := regexp.MustCompile(`don't\(\)(.*?)do\(\)`)

	var result int64

	line = matcherDos.ReplaceAllString(strings.ReplaceAll(line, "\n", ""), "")

	for _, mul := range matcher.FindAllString(line, -1) {
		left, right, err := numbers(mul)
		common.CheckError(err)

		result += left * right
	}

	fmt.Println("solution day 03 part 02:", result)
}
