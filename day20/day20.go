package day20

import (
	"adventofcode2024/common"
	"fmt"
	"math"
	"sort"
)

type Coords struct {
	x, y int
}

type racetrack struct {
	walls map[Coords]struct{}
	lines [][]rune
	start Coords
}

func parse(filename string) (result racetrack) {
	input, err := common.ReadInput("day20/" + filename)
	common.CheckError(err)

	var y int

	result.walls = make(map[Coords]struct{})

	for line := range input.ReadLines {
		result.lines = append(result.lines, []rune(line))

		for x, cell := range line {
			if cell == '#' {
				result.walls[Coords{x, y}] = struct{}{}
			}

			if cell == 'S' {
				result.start = Coords{x, y}
			}
		}

		y++
	}

	return result
}

type Item struct {
	Coords
	distance int
}
type queue []Item

func (q queue) Push(s Item) queue {
	q = append(q, s)

	sort.Slice(q, func(i, j int) bool {
		return q[i].distance < q[j].distance
	})

	return q
}

func (q queue) Pop() (Item, queue) {
	return q[0], q[1:]
}

func bfs(track racetrack) map[Coords]int {
	distances := make(map[Coords]int)

	queue := queue{{Coords: track.start, distance: 0}}
	distances[track.start] = 0

	var current Item

	for len(queue) > 0 {
		current, queue = queue.Pop()

		for _, next := range []Coords{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			neighbour := Coords{x: current.x + next.x, y: current.y + next.y}

			if _, wall := track.walls[neighbour]; wall || !track.inside(neighbour) {
				continue
			}

			if _, visited := distances[neighbour]; !visited {
				queue = queue.Push(Item{Coords: neighbour, distance: current.distance + 1})
				distances[neighbour] = current.distance + 1
			}
		}
	}

	return distances
}

func (r racetrack) inside(c Coords) bool {
	return c.x >= 0 && c.y >= 0 && c.x < len(r.lines[0]) && c.y <= len(r.lines)
}

func Solve(filename string, limitCheat int, cutoff int) {
	result := parse(filename)

	distances := bfs(result)

	cheatCount := make(map[int]int)

	for coords, distance := range distances {
		for diffY := -limitCheat; diffY <= limitCheat; diffY++ {
			for diffX := -limitCheat; diffX <= limitCheat; diffX++ {
				cheatCoords := Coords{coords.x + diffX, coords.y + diffY}

				if !result.inside(cheatCoords) {
					continue
				}

				cheatDist := int(math.Abs(float64(diffX)) + math.Abs(float64(diffY)))
				if cheatDist > limitCheat || cheatDist == 0 {
					continue
				}

				if endCheat, ok := distances[cheatCoords]; ok {
					cheated := endCheat - distance - cheatDist
					if cheated > 0 {
						cheatCount[cheated]++
					}
				}
			}
		}
	}

	var sum int

	for cheat, count := range cheatCount {
		if cheat >= cutoff {
			sum += count
		}
	}

	fmt.Println("solution day 20:", sum)
}
