package day15

import (
	"adventofcode2024/common"
	"fmt"
	"slices"
	"strings"
)

type position struct {
	x, y int
}
type space [][]item
type direction rune
type item rune

const (
	up       direction = '^'
	right    direction = '>'
	down     direction = 'v'
	left     direction = '<'
	robot    item      = '@'
	box      item      = 'O'
	wall     item      = '#'
	empty    item      = '.'
	boxLeft  item      = '['
	boxRight item      = ']'
)

func (d direction) opposite() direction {
	switch d {
	case up:
		return down
	case down:
		return up
	case left:
		return right
	case right:
		return left
	}

	return d

}

func (d direction) diff() position {
	switch d {
	case up:
		return position{
			x: 0,
			y: -1,
		}
	case down:
		return position{
			x: 0,
			y: 1,
		}
	case left:
		return position{
			x: -1,
			y: 0,
		}
	case right:
		return position{
			x: 1,
			y: 0,
		}
	}

	return position{}
}

func (s space) withinBounds(p position) bool {
	return p.x >= 0 && p.x < len(s[0]) && p.y >= 0 && p.y < len(s)
}

func (p position) move(d direction) position {
	switch d {
	case up:
		return position{
			x: p.x,
			y: p.y - 1,
		}
	case down:
		return position{
			x: p.x,
			y: p.y + 1,
		}
	case left:
		return position{
			x: p.x - 1,
			y: p.y,
		}
	case right:
		return position{
			x: p.x + 1,
			y: p.y,
		}
	}

	return p
}

func (s space) print() {
	for _, line := range s {
		fmt.Println(string(line))
	}
}

func (s space) item(p position) item {
	return s[p.y][p.x]
}

func parse(filename string, double bool) (result space, directions []direction, robPos position) {
	input, err := common.ReadInput("day15/" + filename)
	common.CheckError(err)

	var y int

	for line := range input.ReadLines {
		if strings.ContainsAny(line, fmt.Sprintf("%c%c%c%c", up, down, left, right)) {
			directions = append(directions, []direction(line)...)

			continue
		}

		if double {
			line = strings.ReplaceAll(line, string(wall), string(wall)+string(wall))
			line = strings.ReplaceAll(line, string(box), "[]")
			line = strings.ReplaceAll(line, string(empty), string(empty)+string(empty))
			line = strings.ReplaceAll(line, string(robot), string(robot)+string(empty))
		}

		if x := strings.Index(line, "@"); x >= 0 {
			robPos.x = x
			robPos.y = y
		}

		result = append(result, []item(line))

		y++
	}

	return result, directions, robPos
}

func Solve1(filename string) {
	warehouse, directions, robotPos := parse(filename, false)

moveLoop:
	for _, dir := range directions {
		switch warehouse[robotPos.move(dir).y][robotPos.move(dir).x] {
		case box:
			for p := robotPos.move(dir); warehouse.withinBounds(p); p = p.move(dir) {
				switch warehouse.item(p) {
				case box:
					continue
				case wall:
					continue moveLoop
				case empty:
					warehouse[robotPos.y][robotPos.x] = empty
					warehouse[robotPos.move(dir).y][robotPos.move(dir).x] = robot
					warehouse[p.y][p.x] = box
					robotPos = robotPos.move(dir)
					continue moveLoop
				}
			}
		case empty:
			warehouse[robotPos.y][robotPos.x] = empty
			warehouse[robotPos.move(dir).y][robotPos.move(dir).x] = robot
			robotPos = robotPos.move(dir)
		case wall:
			continue
		}
	}

	var sum int

	for y, line := range warehouse {
		for x, cell := range line {
			switch item(cell) {
			case box:
				sum += 100*y + x
			}
		}
	}

	fmt.Println("solution day 15 part 01:", sum)
}

func reverseSlice[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Solve2(filename string) {
	warehouse, directions, robotPos := parse(filename, true)

dirLoop:
	for _, dir := range directions {
		switch warehouse[robotPos.move(dir).y][robotPos.move(dir).x] {
		case boxLeft, boxRight:
			if dir == left || dir == right {
				for p := robotPos.move(dir); warehouse.withinBounds(p); p = p.move(dir) {
					switch warehouse[p.y][p.x] {
					case boxLeft, boxRight:
						continue
					case wall:
						continue dirLoop
					case empty:
						for current := p; current != robotPos; current = current.move(dir.opposite()) {
							warehouse[current.y][current.x] = warehouse[current.move(dir.opposite()).y][current.move(dir.opposite()).x]
						}

						warehouse[robotPos.y][robotPos.x] = empty
						warehouse[robotPos.move(dir).y][robotPos.move(dir).x] = robot

						robotPos = robotPos.move(dir)
						continue dirLoop
					}
				}
			}
			if dir == up || dir == down {
				stack := [][]position{{robotPos}}

				// find pyramid of boxes
				for {
					var newEntries []position
					for _, pos := range stack[len(stack)-1] {
						switch warehouse.item(pos.move(dir)) {
						case wall:
							continue dirLoop
						case empty:
							continue
						case boxLeft:
							if !slices.Contains(newEntries, pos.move(dir)) {
								newEntries = append(newEntries, pos.move(dir), pos.move(dir).move(right))
							}
						case boxRight:
							if !slices.Contains(newEntries, pos.move(dir)) {
								newEntries = append(newEntries, pos.move(dir).move(left), pos.move(dir))
							}
						}
					}

					// only found empties
					if newEntries == nil {
						break
					}

					stack = append(stack, newEntries)
				}

				// move boxes! start from end
				reverseSlice(stack)

				for _, boxes := range stack[:len(stack)-1] {
					for _, item := range boxes {
						warehouse[item.move(dir).y][item.move(dir).x] = warehouse[item.y][item.x]
						warehouse[item.y][item.x] = empty
					}
				}

				warehouse[robotPos.y][robotPos.x] = empty
				warehouse[robotPos.move(dir).y][robotPos.move(dir).x] = robot
				robotPos = robotPos.move(dir)
			}
		case empty:
			warehouse[robotPos.y][robotPos.x] = empty
			warehouse[robotPos.move(dir).y][robotPos.move(dir).x] = robot
			robotPos = robotPos.move(dir)
		case wall:
			continue
		}
	}

	var sum int

	for y, line := range warehouse {
		for x, cell := range line {
			switch cell {
			case boxLeft:
				sum += 100*y + x
			}
		}
	}

	fmt.Println("solution day 15 part 02:", sum)
}
