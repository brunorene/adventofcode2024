package day21

import (
	"adventofcode2024/common"
	"fmt"
	"strconv"
	"strings"
)

func parse(filename string) (codes []string) {
	input, err := common.ReadInput("day21/" + filename)
	common.CheckError(err)

	for line := range input.ReadLines {
		codes = append(codes, line)
	}

	return codes
}

const (
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'
	a     = 'A'
	gap   = 'X'
)

type Coords struct {
	x, y int
}

type pad map[rune]Coords

var (
	numPad = pad{
		'7': {0, 0},
		'8': {1, 0},
		'9': {2, 0},
		'4': {0, 1},
		'5': {1, 1},
		'6': {2, 1},
		'1': {0, 2},
		'2': {1, 2},
		'3': {2, 2},
		gap: {0, 3},
		'0': {1, 3},
		a:   {2, 3},
	}
)

func (p pad) getCoords(r rune) Coords {
	return p[r]
}

func moveNumericPad(code string, pressButton map[string]int) {
	start := numPad.getCoords(a)
	for _, button := range code {
		path := toNumericButton(start, numPad.getCoords(button))
		start = numPad.getCoords(button)
		path += string(a)
		pressButton[path]++
	}
}

func toNumericButton(start Coords, target Coords) (path string) {
	gapCoords := numPad.getCoords(gap)

	for start != target {
		if start.x > target.x {
			if start.y == gapCoords.y && target.x == gapCoords.x {
				path += strings.Repeat(string(up), start.y-target.y)
				start.y = target.y

				continue
			}

			start.x--
			path += string(left)

			continue
		}

		if start.y > target.y {
			start.y--
			path += string(up)

			continue
		}

		if start.y < target.y {
			if start.x == gapCoords.x && target.y == gapCoords.y {
				path += strings.Repeat(string(right), target.x-start.x)
				start.x = target.x

				continue
			}

			start.y++
			path += string(down)

			continue
		}

		if start.x < target.x {
			start.x++
			path += string(right)

			continue
		}
	}

	return path
}

func moveDirPad(code string, count int, pressButton map[string]int) {
	dirMoves := map[[2]rune][]rune{
		{a, up}:       {left, a},
		{a, down}:     {left, down, a},
		{a, left}:     {down, left, left, a},
		{a, right}:    {down, a},
		{up, a}:       {right, a},
		{up, down}:    {down, a},
		{up, left}:    {down, left, a},
		{up, right}:   {down, right, a},
		{down, a}:     {up, right, a},
		{down, up}:    {up, a},
		{down, left}:  {left, a},
		{down, right}: {right, a},
		{left, a}:     {right, right, up, a},
		{left, up}:    {right, up, a},
		{left, down}:  {right, a},
		{left, right}: {right, right, a},
		{right, a}:    {up, a},
		{right, up}:   {left, up, a},
		{right, down}: {left, a},
		{right, left}: {left, left, a},
	}

	start := a
	for _, button := range code {
		path := string(a)

		if start != button {
			path = string(dirMoves[[2]rune{start, button}])
		}

		start = button

		pressButton[path] += count
	}
}

func Solve(filename string, dirPadCount int) {
	codes := parse(filename)

	var sum int64

	for _, code := range codes {
		pressA := make(map[string]int)

		moveNumericPad(code, pressA)

		for range dirPadCount {
			pressNext := make(map[string]int)

			for dirPath, count := range pressA {
				moveDirPad(dirPath, count, pressNext)
			}

			pressA = pressNext
		}

		var totalLength int64
		for dirPath, count := range pressA {
			totalLength += int64(count) * int64(len(dirPath))
		}

		numCode, err := strconv.ParseInt(code[:len(code)-1], 10, 64)
		common.CheckError(err)

		sum += totalLength * numCode
	}

	fmt.Println("solution day 21:", sum)
}
