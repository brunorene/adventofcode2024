package day14

import (
	"adventofcode2024/common"
	"fmt"
	"regexp"
	"strconv"
)

type point struct {
	x, y int
}

type robot struct {
	position, velocity point
	quadrant           int
}

type grid [103][101]int

func (r *robot) move(time, rows, cols int) {
	r.position.x = (r.position.x + time*r.velocity.x) % cols
	if r.position.x < 0 {
		r.position.x += cols
	}

	r.position.y = (r.position.y + time*r.velocity.y) % rows
	if r.position.y < 0 {
		r.position.y += rows
	}
}

func parse(filename string) (result []robot) {
	input, err := common.ReadInput("day14/" + filename)
	common.CheckError(err)

	buttonVal := regexp.MustCompile(`(-|)\d+`)

	for line := range input.ReadLines {
		vals := buttonVal.FindAllString(line, -1)

		var current robot

		current.position.x, err = strconv.Atoi(vals[0])
		common.CheckError(err)

		current.position.y, err = strconv.Atoi(vals[1])
		common.CheckError(err)

		current.velocity.x, err = strconv.Atoi(vals[2])
		common.CheckError(err)

		current.velocity.y, err = strconv.Atoi(vals[3])
		common.CheckError(err)

		result = append(result, current)
	}

	return result
}

func (g grid) print(middle bool) {
	for y, line := range g {
		for x, count := range line {
			if !middle && (x == len(g[0])/2 || y == len(g)/2) {
				fmt.Print(" ")

				continue
			}

			if count == 0 {
				fmt.Print(".")

				continue
			}

			fmt.Printf("%d", count)
		}

		fmt.Println()
	}
}

func Solve1(filename string) {
	robots := parse(filename)

	var tiles grid

	for idx := range robots {
		robots[idx].move(100, len(tiles), len(tiles[0]))
		tiles[robots[idx].position.y][robots[idx].position.x]++
	}

	quadrants := [4]point{
		{len(tiles[0]), len(tiles)},
		{len(tiles[0]), len(tiles) / 2},
		{len(tiles[0]) / 2, len(tiles)},
		{len(tiles[0]) / 2, len(tiles) / 2},
	}

	var quadrCount [4]int64

rbtLoop:
	for ridx, rbt := range robots {
		for idx, quadr := range quadrants {
			if rbt.position.x == len(tiles[0])/2 || rbt.position.y == len(tiles)/2 {
				continue rbtLoop
			}

			if rbt.position.x < quadr.x && rbt.position.y < quadr.y {
				robots[ridx].quadrant = idx
			}
		}

		quadrCount[robots[ridx].quadrant]++
	}

	mul := int64(1)

	for _, count := range quadrCount {
		mul *= count
	}

	tiles.print(false)

	fmt.Println("solution day 14 part 01:", mul)
}

func (g grid) maybeATree() bool {
	for y, line := range g {
	xLoop:
		for x, val := range line {
			if x > 2 && x < len(g[0])-3 && y > 2 && y < len(g)-3 && val == 1 {
				for diffY := -3; diffY <= 3; diffY++ {
					for diffX := -3; diffX <= 3; diffX++ {
						if g[y+diffY][x+diffX] != 1 {
							continue xLoop
						}
					}
				}

				return true
			}
		}
	}

	return false
}

func Solve2(filename string) {
	robots := parse(filename)

	step := 1

mainLoop:
	for {
		var tiles grid

		for idx := range robots {
			robots[idx].move(1, len(tiles), len(tiles[0]))
			tiles[robots[idx].position.y][robots[idx].position.x] = 1
		}

		if tiles.maybeATree() {
			fmt.Println(step)
			tiles.print(true)

			break mainLoop
		}

		step++
	}

	fmt.Println("solution day 14 part 02:", step)
}
