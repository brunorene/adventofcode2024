package day01

import (
	"adventofcode2024/common"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Solve1(filename string) {
	var sum int64

	input, err := common.ReadInput("day01/" + filename)
	common.CheckError(err)

	var left, right []int

	for line := range input.ReadLines {
		parts := strings.Split(line, "   ")

		leftNum, err := strconv.Atoi(parts[0])
		common.CheckError(err)

		left = append(left, leftNum)

		rightNum, err := strconv.Atoi(parts[1])
		common.CheckError(err)

		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	for idx := range left {
		sum += int64(math.Abs(float64(left[idx] - right[idx])))
	}

	fmt.Println("solution day 01 part 01:", sum)
}

func Solve2(filename string) {
	var sum int64

	input, err := common.ReadInput("day01/" + filename)
	common.CheckError(err)

	left := make(map[int]int64)
	leftCount := make(map[int]int64)
	var right []int

	for line := range input.ReadLines {
		parts := strings.Split(line, "   ")

		leftNum, err := strconv.Atoi(parts[0])
		common.CheckError(err)

		leftCount[leftNum] += 1
		left[leftNum] = 0

		rightNum, err := strconv.Atoi(parts[1])
		common.CheckError(err)

		right = append(right, rightNum)
	}

	for _, num := range right {
		left[num] += int64(num)
	}

	for num, subSum := range left {
		sum += subSum * leftCount[num]
	}

	fmt.Println("solution day 01 part 02:", sum)
}
