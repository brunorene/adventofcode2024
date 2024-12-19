package day16

import (
	"adventofcode2024/common"
	"fmt"
	"math"
	"sort"
)

type position struct {
	x, y int
}

type maze struct {
	lines [][]rune
}

type direction position

type movement struct {
	pos position
	dir direction
}

type score struct {
	move  movement
	value int
	path  string
}

var (
	north = direction{x: 0, y: -1}
	south = direction{x: 0, y: 1}
	east  = direction{x: 1, y: 0}
	west  = direction{x: -1, y: 0}
)

func parse(filename string) (result maze) {
	input, err := common.ReadInput("day16/" + filename)
	common.CheckError(err)

	for line := range input.ReadLines {
		result.lines = append(result.lines, []rune(line))
	}

	return result
}

func (m maze) position(name rune) (p position) {
	for y, line := range m.lines {
		for x, cell := range line {
			if cell == name {
				return position{x: x, y: y}
			}
		}
	}

	return position{x: -1, y: -1}
}

func (d direction) turnLeft() (next direction) {
	switch d {
	case north:
		return west
	case south:
		return east
	case east:
		return north
	case west:
		return south
	}

	return d
}

func (d direction) turnRight() (next direction) {
	switch d {
	case north:
		return east
	case south:
		return west
	case east:
		return south
	case west:
		return north
	}

	return d
}

type queue []score

func (q queue) Push(s score) queue {
	q = append(q, s)

	sort.Slice(q, func(i, j int) bool {
		return q[i].value < q[j].value
	})

	return q
}

func (q queue) Pop() (score, queue) {
	return q[0], q[1:]
}

func (m maze) neighbours(move movement, path string) []score {
	return []score{
		{movement{position{x: move.pos.x + move.dir.x, y: move.pos.y + move.dir.y}, move.dir}, 1, path + "F"},
		{movement{position{x: move.pos.x, y: move.pos.y}, move.dir.turnLeft()}, 1000, path + "L"},
		{movement{position{x: move.pos.x, y: move.pos.y}, move.dir.turnRight()}, 1000, path + "R"},
	}
}

func (m maze) dijkstra(start, end position, currentDirection direction) (result int, bestCount int) {
	visited := make(map[movement]struct{})
	lowest := queue{}.Push(score{movement{start, currentDirection}, 0, ""})
	var current score
	minScore := math.MaxInt64
	var bestPaths []string

	for {
		current, lowest = lowest.Pop()

		if current.value > minScore {
			break
		}

		if current.move.pos == end {
			minScore = current.value
			bestPaths = append(bestPaths, current.path)
			continue
		}

		visited[current.move] = struct{}{}

		for _, neighbour := range m.neighbours(current.move, current.path) {
			if _, ok := visited[neighbour.move]; !ok && neighbour.move.pos.x >= 0 &&
				neighbour.move.pos.x < len(m.lines[0]) &&
				neighbour.move.pos.y >= 0 && neighbour.move.pos.y < len(m.lines) &&
				m.lines[neighbour.move.pos.y][neighbour.move.pos.x] != '#' {

				newScore := score{
					move:  neighbour.move,
					value: current.value + neighbour.value,
					path:  neighbour.path,
				}

				lowest = lowest.Push(newScore)
			}
		}
	}

	return minScore, m.bestPatchCount(start, currentDirection, bestPaths)
}

func (m maze) bestPatchCount(start position, dir direction, paths []string) int {
	uniques := make(map[position]struct{})

	for _, path := range paths {
		current := start
		uniques[current] = struct{}{}
		m.lines[current.y][current.x] = 'O'
		currentDirection := dir
		for _, cell := range path {
			switch cell {
			case 'F':
				current = position{x: current.x + currentDirection.x, y: current.y + currentDirection.y}
				uniques[current] = struct{}{}
				m.lines[current.y][current.x] = 'O'
			case 'L':
				currentDirection = currentDirection.turnLeft()
			case 'R':
				currentDirection = currentDirection.turnRight()
			}
		}
	}

	return len(uniques)
}

func (m maze) print() {
	for _, line := range m.lines {
		fmt.Println(string(line))
	}
}

func Solve1(filename string) {
	maze := parse(filename)

	start := maze.position('S')
	end := maze.position('E')
	currentDirection := east

	score, _ := maze.dijkstra(start, end, currentDirection)

	fmt.Println("solution day 16 part 01:", score)
}

func Solve2(filename string) {
	maze := parse(filename)

	start := maze.position('S')
	end := maze.position('E')
	currentDirection := east

	_, best := maze.dijkstra(start, end, currentDirection)

	maze.print()

	fmt.Println("solution day 16 part 02:", best)
}
