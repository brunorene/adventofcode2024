package day11

import (
	"adventofcode2024/common"
	"fmt"
	"strconv"
	"strings"
)

func parse(filename string) (result []int64) {
	input, err := common.ReadInput("day11/" + filename)
	common.CheckError(err)

	for line := range input.ReadLines {
		parts := strings.Fields(line)

		for _, part := range parts {
			value, err := strconv.ParseInt(part, 10, 64)
			common.CheckError(err)

			result = append(result, value)
		}
	}

	return result
}

func Solve1(filename string) {
	result := parse(filename)

	for range 25 {
		var step []int64

		for _, num := range result {
			if num == 0 {
				step = append(step, 1)

				continue
			}

			strNum := fmt.Sprintf("%d", num)

			if len(strNum)%2 == 0 {
				left, err := strconv.ParseInt(strNum[:len(strNum)/2], 10, 64)
				common.CheckError(err)

				right, err := strconv.ParseInt(strNum[len(strNum)/2:], 10, 64)
				common.CheckError(err)

				step = append(step, []int64{left, right}...)

				continue
			}

			step = append(step, num*2024)
		}

		result = step
	}

	fmt.Println("solution day 11 part 01:", len(result))
}

func newStones(stone int64, cache map[[2]int64]int64, counter int64) int64 {
	if counter == 0 {
		return 1
	}

	if result, exists := cache[[2]int64{stone, counter}]; exists {
		return result
	}

	if stone == 0 {
		result := newStones(1, cache, counter-1)
		cache[[2]int64{1, counter - 1}] = result

		return result
	}

	strNum := fmt.Sprintf("%d", stone)

	if len(strNum)%2 == 0 {

		left, err := strconv.ParseInt(strNum[:len(strNum)/2], 10, 64)
		common.CheckError(err)

		right, err := strconv.ParseInt(strNum[len(strNum)/2:], 10, 64)
		common.CheckError(err)

		resultLeft := newStones(left, cache, counter-1)
		cache[[2]int64{left, counter - 1}] = resultLeft
		resultRight := newStones(right, cache, counter-1)
		cache[[2]int64{right, counter - 1}] = resultRight

		return resultLeft + resultRight
	}

	result := newStones(stone*2024, cache, counter-1)
	cache[[2]int64{stone * 2024, counter - 1}] = result

	return result
}

func Solve2(filename string) {
	result := parse(filename)

	var sum int64
	cache := make(map[[2]int64]int64)

	for _, num := range result {
		sum += newStones(num, cache, 75)
	}

	fmt.Println("solution day 11 part 02:", sum)
}
