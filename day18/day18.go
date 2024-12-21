package day18

import (
	"adventofcode2024/common"
	"fmt"
	"math"
	"sort"
)

type Coords struct {
	x, y int
}

type Item struct {
	Coords
	distance int
}
type RAM struct {
	grid       map[Coords]struct{}
	list       []Coords
	maxX, maxY int
}

type direction Coords

var (
	north = direction{x: 0, y: -1}
	south = direction{x: 0, y: 1}
	east  = direction{x: 1, y: 0}
	west  = direction{x: -1, y: 0}
)

func parse(filename string, count int) (result RAM) {
	input, err := common.ReadInput("day18/" + filename)
	common.CheckError(err)

	var currentCount int

	result.grid = make(map[Coords]struct{})

	for line := range input.ReadLines {
		var x, y int
		_, err = fmt.Sscanf(line, "%d,%d", &x, &y)
		common.CheckError(err)

		if currentCount < count {
			result.grid[Coords{x: x, y: y}] = struct{}{}
			result.list = append(result.list, Coords{x: x, y: y})
		}

		if x > result.maxX {
			result.maxX = x
		}

		if y > result.maxY {
			result.maxY = y
		}

		currentCount++
	}

	return result
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

func bfs(result RAM) int {
	visited := make(map[Coords]struct{})
	queue := queue{{Coords: Coords{x: 0, y: 0}, distance: 0}}

	var (
		finalDistance int
		current       Item
	)

	for len(queue) > 0 {
		current, queue = queue.Pop()

		if current.Coords.x == result.maxX && current.Coords.y == result.maxY {
			finalDistance = current.distance

			break
		}

		if _, ok := visited[current.Coords]; ok {
			continue
		}

		visited[current.Coords] = struct{}{}

		for _, next := range []Coords{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			neighbour := Coords{x: current.x + next.x, y: current.y + next.y}

			if _, corrupt := result.grid[neighbour]; !corrupt && neighbour.x >= 0 && neighbour.y >= 0 &&
				neighbour.x <= result.maxX && neighbour.y <= result.maxY {
				queue = queue.Push(Item{Coords: neighbour, distance: current.distance + 1})
			}
		}
	}
	return finalDistance
}

func Solve1(filename string) {
	result := parse(filename, 1024)

	finalDistance := bfs(result)

	fmt.Println("solution day 18 part 01:", finalDistance)
}

func Solve2(filename string) {
	result := parse(filename, math.MaxInt64)

	var blocker Coords

	for end := 1024; end <= len(result.grid); end++ {
		current := RAM{
			grid: make(map[Coords]struct{}),
			list: result.list,
			maxX: result.maxX,
			maxY: result.maxY,
		}

		for i := 0; i < end; i++ {
			current.grid[result.list[i]] = struct{}{}
		}

		distance := bfs(current)

		fmt.Println(distance)

		if distance == 0 {
			blocker = result.list[end-1]

			break
		}
	}

	fmt.Printf("solution day 18 part 02: %d,%d\n", blocker.x, blocker.y)
}
