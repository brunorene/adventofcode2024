package day22

import (
	"adventofcode2024/common"
	"fmt"
	"math"
	"strconv"
)

func parse(filename string) (result []int64) {
	input, err := common.ReadInput("day22/" + filename)
	common.CheckError(err)

	for line := range input.ReadLines {
		number, err := strconv.ParseInt(line, 10, 64)
		common.CheckError(err)

		result = append(result, number)
	}

	return result
}

func transform(secret int64) int64 {
	mod := int64(math.Pow(2, 24)) - 1

	secret = ((secret << 6) ^ secret) & mod
	secret = ((secret >> 5) ^ secret) & mod
	return ((secret << 11) ^ secret) & mod
}

func nextSecret(count int, current int64) int64 {

	for range count {
		current = transform(current)
	}

	return current
}

func Solve1(filename string) {
	result := parse(filename)

	var sum int64

	for _, secret := range result {
		sum += nextSecret(2000, secret)
	}

	fmt.Println("solution day 22 part 01:", sum)
}

func Solve2(filename string) {
	result := parse(filename)

	bananas := make(map[[4]int64]int64)

	for _, buyer := range result {
		var diffs []int64
		var prices []int64

		current := buyer

		for range 2000 {
			next := transform(current)
			currentDigit := current % 10
			nextDigit := next % 10

			prices = append(prices, nextDigit)
			diffs = append(diffs, currentDigit-nextDigit)
			current = next
		}

		var seq []int64
		visited := make(map[[4]int64]struct{})

		for idx, change := range diffs {
			seq = append(seq, change)
			if len(seq) == 4 {
				array := [4]int64{seq[0], seq[1], seq[2], seq[3]}
				if _, ok := visited[array]; !ok {
					bananaCount := prices[idx]
					bananas[array] += bananaCount
					visited[array] = struct{}{}
				}

				seq = seq[1:]
			}
		}
	}

	var maxBananas int64

	for _, count := range bananas {
		if count > maxBananas {
			maxBananas = count
		}
	}

	fmt.Println("solution day 22 part 02:", maxBananas)
}
