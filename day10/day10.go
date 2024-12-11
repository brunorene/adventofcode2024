package day10

import (
	"adventofcode2024/common"
	"fmt"
	"strconv"
)

type grid [][]int

type cell struct {
	x, y, value int
}

func parse(filename string) (result grid) {
	input, err := common.ReadInput("day10/" + filename)
	common.CheckError(err)

	for line := range input.ReadLines {
		intLine := make([]int, 0, len(line))
		for _, c := range line {
			val, err := strconv.Atoi(string(c))
			common.CheckError(err)

			intLine = append(intLine, val)
		}

		result = append(result, intLine)
	}

	return result
}

func (g grid) readCells(yield func(c cell) bool) {
	for y, line := range g {
		for x, v := range line {
			yield(cell{x, y, v})
		}
	}
}

func (g grid) score1(c cell, visited map[cell]struct{}) (result map[cell]struct{}) {
	result = make(map[cell]struct{})

	if c.value == 9 {
		result[c] = struct{}{}

		return result
	}

	visited[c] = struct{}{}

	for _, nextDiff := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		nextX := c.x + nextDiff[0]
		nextY := c.y + nextDiff[1]

		if nextX < 0 || nextY < 0 || nextX >= len(g[0]) || nextY >= len(g) || g[c.y][c.x]+1 != g[nextY][nextX] {
			continue
		}

		next := cell{nextX, nextY, g[nextY][nextX]}

		if _, exists := visited[next]; exists {
			continue
		}

		endings := g.score1(next, visited)

		for k := range endings {
			result[k] = struct{}{}
		}
	}

	return result
}

func (g grid) score2(c cell) (result map[[10]cell]struct{}) {
	result = make(map[[10]cell]struct{})

	if c.value == 9 {
		key := [10]cell{}
		key[9] = c

		result[key] = struct{}{}

		return result
	}

	for _, nextDiff := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		nextX := c.x + nextDiff[0]
		nextY := c.y + nextDiff[1]

		if nextX < 0 || nextY < 0 || nextX >= len(g[0]) || nextY >= len(g) || g[c.y][c.x]+1 != g[nextY][nextX] {
			continue
		}

		next := cell{nextX, nextY, g[nextY][nextX]}

		tails := g.score2(next)

		for k := range tails {
			k[c.value] = c

			result[k] = struct{}{}
		}
	}

	return result
}

func Solve1(filename string) {
	result := parse(filename)

	var sum int

	for c := range result.readCells {
		if c.value == 0 {
			sum += len(result.score1(c, make(map[cell]struct{})))
		}
	}

	fmt.Println("solution day 10 part 01:", sum)
}

func Solve2(filename string) {
	result := parse(filename)

	var sum int

	for c := range result.readCells {
		if c.value == 0 {
			sum += len(result.score2(c))
		}
	}

	fmt.Println("solution day 10 part 02:", sum)
}
