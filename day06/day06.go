package day06

import (
	"adventofcode2024/common"
	"fmt"
)

type direction struct {
	sumX, sumY int
}

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	right = direction{1, 0}
)

func (d direction) turnRight() (next direction) {
	switch d {
	case up:
		return right
	case down:
		return left
	case right:
		return down
	default:
		return up
	}
}

func Solve1(filename string) {
	_, _, _, visited := visited(filename)

	fmt.Println("solution day 06 part 01:", len(visited))
}

func visited(filename string) (grid [][]rune, x, y int, visited map[[2]int]struct{}) {
	grid, origGuardX, origGuardY := parse(filename)

	currDir := up

	guardX := origGuardX
	guardY := origGuardY

	visited = make(map[[2]int]struct{})
	visited[[2]int{guardX, guardY}] = struct{}{}

	for {
		nextPosX := guardX + currDir.sumX
		nextPosY := guardY + currDir.sumY

		if nextPosX < 0 || nextPosY < 0 || nextPosX >= len(grid[0]) || nextPosY >= len(grid) {
			break
		}

		if grid[nextPosY][nextPosX] == '#' {
			currDir = currDir.turnRight()

			continue
		}

		guardX = nextPosX
		guardY = nextPosY

		visited[[2]int{guardX, guardY}] = struct{}{}
	}

	return grid, origGuardX, origGuardY, visited
}

func Solve2(filename string) {
	defer common.Timer("solve 2")()

	grid, origGuardX, origGuardY, visited := visited(filename)

	var hasLoop int

	for candidate := range visited {
		currDir := up
		guardY := origGuardY
		guardX := origGuardX

		grid[candidate[1]][candidate[0]] = '#'
		overlaps := make(map[[4]int]struct{})

		for {
			nextPosX := guardX + currDir.sumX
			nextPosY := guardY + currDir.sumY

			if nextPosX < 0 || nextPosY < 0 || nextPosX >= len(grid[0]) || nextPosY >= len(grid) {
				break
			}

			if grid[nextPosY][nextPosX] == '#' {
				currDir = currDir.turnRight()

				continue
			}

			guardX = nextPosX
			guardY = nextPosY

			if _, exists := overlaps[[4]int{guardX, guardY, currDir.sumX, currDir.sumY}]; exists {
				hasLoop++

				break
			}

			overlaps[[4]int{guardX, guardY, currDir.sumX, currDir.sumY}] = struct{}{}
		}

		grid[candidate[1]][candidate[0]] = '.'
	}

	fmt.Println("solution day 06 part 02:", hasLoop)
}

func parse(filename string) (result [][]rune, guardX, guardY int) {
	input, err := common.ReadInput("day06/" + filename)
	common.CheckError(err)

	var y int

	for line := range input.ReadLines {
		result = append(result, []rune{})

		for x, cell := range line {
			if cell == '^' {
				guardX = x
				guardY = y
			}

			result[len(result)-1] = append(result[len(result)-1], cell)
		}

		y++
	}

	return result, guardX, guardY
}
