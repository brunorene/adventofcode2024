package day12

import (
	"adventofcode2024/common"
	"fmt"
	"sort"
)

type farm struct {
	lines [][]rune
}

func parse(filename string) (result farm) {
	input, err := common.ReadInput("day12/" + filename)
	common.CheckError(err)

	for line := range input.ReadLines {
		result.lines = append(result.lines, []rune(line))
	}

	return result
}

func (f *farm) onePlot() (name rune, area int64, fences map[[4]int]struct{}) {
	var start [2]int

upperLoop:
	for y, line := range f.lines {
		for x, plant := range line {
			if plant != '.' {
				start = [2]int{x, y}
				name = plant

				break upperLoop
			}
		}
	}

	if name == 0 {
		return
	}

	fences = make(map[[4]int]struct{})
	plotCells := make(map[[4]int]struct{})

	f.lines[start[1]][start[0]] = '.'
	area++

	queue := [][2]int{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			neighbour := [2]int{current[0] + next[0], current[1] + next[1]}

			fences[[4]int{neighbour[0], neighbour[1], next[0], next[1]}] = struct{}{}
			plotCells[[4]int{current[0], current[1], next[0], next[1]}] = struct{}{}

			if neighbour[0] < 0 || neighbour[1] < 0 || neighbour[0] >= len(f.lines[0]) || neighbour[1] >= len(f.lines) ||
				f.lines[neighbour[1]][neighbour[0]] != name || f.lines[neighbour[1]][neighbour[0]] == '.' {

				continue
			}

			queue = append(queue, neighbour)
			f.lines[neighbour[1]][neighbour[0]] = '.'
			area++
		}
	}

	for plotCell := range plotCells {
		delete(fences, plotCell)
	}

	return name, area, fences
}

func (f *farm) print() {
	for _, line := range f.lines {
		fmt.Println(string(line))
	}
}

func Solve1(filename string) {
	result := parse(filename)

	var sum int64

	for {
		_, area, fences := result.onePlot()
		sum += area * int64(len(fences))

		if area == 0 {
			break
		}
	}

	fmt.Println("solution day 12 part 01:", sum)
}

func sides(fences map[[4]int]struct{}) (sides [][][4]int) {
	accumX := make(map[[3]int][][4]int)
	accumY := make(map[[3]int][][4]int)

	for fence := range fences {
		if fence[2] == 0 {
			accumY[[3]int{fence[1], fence[2], fence[3]}] = append(accumY[[3]int{fence[1], fence[2], fence[3]}], fence)
		}

		if fence[3] == 0 {
			accumX[[3]int{fence[0], fence[2], fence[3]}] = append(accumX[[3]int{fence[0], fence[2], fence[3]}], fence)
		}
	}

	sides = count(accumX, 1)
	sides = append(sides, count(accumY, 0)...)

	return sides
}

func count(accum map[[3]int][][4]int, index int) (sides [][][4]int) {
	for _, side := range accum {
		sort.Slice(side, func(i, j int) bool {
			return side[i][index] < side[j][index]
		})

		sides = append(sides, [][4]int{})

		for _, fence := range side {
			last := sides[len(sides)-1]

			if len(last) > 0 && fence[index]-last[len(last)-1][index] > 1 {
				sides = append(sides, [][4]int{})
			}

			sides[len(sides)-1] = append(sides[len(sides)-1], fence)
		}
	}

	return sides
}

func Solve2(filename string) {
	result := parse(filename)

	var sum int64

	for {
		_, area, fences := result.onePlot()

		if area == 0 {
			break
		}

		sides := sides(fences)

		sum += area * int64(len(sides))
	}

	fmt.Println("solution day 12 part 02:", sum)
}
