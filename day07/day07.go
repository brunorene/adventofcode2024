package day07

import (
	"adventofcode2024/common"
	"fmt"
	"strconv"
	"strings"
)

type equation struct {
	result int64
	terms  []int64
}

func parse(filename string) (equations []equation) {
	input, err := common.ReadInput("day07/" + filename)
	common.CheckError(err)

	for line := range input.ReadLines {
		parts := strings.Split(line, ":")

		result, err := strconv.ParseInt(parts[0], 10, 64)
		common.CheckError(err)

		parts = strings.Fields(parts[1])

		item := equation{result: result}

		for _, part := range parts {
			term, err := strconv.ParseInt(part, 10, 64)
			common.CheckError(err)

			item.terms = append(item.terms, term)
		}

		equations = append(equations, item)
	}

	return equations
}

func isValid(eq equation, operator string, part2 bool) bool {
	if len(eq.terms) == 1 {
		return eq.terms[0] == eq.result
	}

	// +
	newEq := equation{terms: append([]int64{eq.terms[0] + eq.terms[1]}, eq.terms[2:]...), result: eq.result}

	if operator == "*" {
		newEq = equation{terms: append([]int64{eq.terms[0] * eq.terms[1]}, eq.terms[2:]...), result: eq.result}
	}

	if operator == "||" {
		start, err := strconv.ParseInt(fmt.Sprintf("%d%d", eq.terms[0], eq.terms[1]), 10, 64)
		common.CheckError(err)

		newEq = equation{terms: append([]int64{start}, eq.terms[2:]...), result: eq.result}
	}

	sumValid := isValid(newEq, "+", part2)
	mulValid := isValid(newEq, "*", part2)
	concValid := false

	if part2 {
		concValid = isValid(newEq, "||", part2)
	}

	return sumValid || mulValid || concValid
}

func Solve1(filename string) {
	equations := parse(filename)

	var sum int64

	for _, eq := range equations {
		sumValid := isValid(eq, "+", false)
		mulValid := isValid(eq, "*", false)

		if sumValid || mulValid {
			sum += eq.result
		}
	}

	fmt.Println("solution day 07 part 01:", sum)
}

func Solve2(filename string) {
	equations := parse(filename)

	var sum int64

	for _, eq := range equations {
		sumValid := isValid(eq, "+", true)
		mulValid := isValid(eq, "*", true)
		concValid := isValid(eq, "||", true)

		if sumValid || mulValid || concValid {
			sum += eq.result
		}
	}

	fmt.Println("solution day 07 part 02:", sum)
}
