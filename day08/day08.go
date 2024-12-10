package day08

import (
	"adventofcode2024/common"
	"fmt"
)

type input struct {
	antennas   map[rune][][2]int
	grid       map[[2]int]struct{}
	rows, cols int
}

func parse(filename string) (result input) {
	input, err := common.ReadInput("day08/" + filename)
	common.CheckError(err)

	var y int

	result.antennas = make(map[rune][][2]int)
	result.grid = make(map[[2]int]struct{})

	for line := range input.ReadLines {
		result.cols = len(line)

		for x, cell := range line {
			if cell != '.' {
				result.antennas[cell] = append(result.antennas[cell], [2]int{x, y})
			}

			result.grid[[2]int{x, y}] = struct{}{}
		}

		y++
	}

	result.rows = y

	return result
}

func (i input) print(filename string, antinodes map[[2]int]struct{}) {
	input, err := common.ReadInput("day08/" + filename)
	common.CheckError(err)

	var y int

	for line := range input.ReadLines {
		for x, cell := range line {
			if _, exists := antinodes[[2]int{x, y}]; exists {
				fmt.Print("x")
			} else {
				fmt.Print(string(cell))
			}
		}

		fmt.Println()
		y++
	}

}

func Solve1(filename string) {
	result := parse(filename)

	antinodes := make(map[[2]int]struct{})

	for _, locations := range result.antennas {
		for idx1 := range locations {
			for idx2 := range locations {
				if idx1 == idx2 {
					continue
				}

				antenna1 := locations[idx1]
				antenna2 := locations[idx2]

				antinode1 := [2]int{2*antenna1[0] - antenna2[0], 2*antenna1[1] - antenna2[1]}
				if _, exists := result.grid[antinode1]; exists {
					antinodes[antinode1] = struct{}{}
				}

				antinode2 := [2]int{2*antenna2[0] - antenna1[0], 2*antenna2[1] - antenna1[1]}
				if _, exists := result.grid[antinode2]; exists {
					antinodes[antinode2] = struct{}{}
				}
			}
		}
	}

	fmt.Println("solution day 08 part 01:", len(antinodes))
}

func Solve2(filename string) {
	result := parse(filename)

	antinodes := make(map[[2]int]struct{})

	for _, locations := range result.antennas {
		for idx1 := range locations {
			for idx2 := range locations {
				if idx1 == idx2 {
					continue
				}

				antenna1 := locations[idx1]
				antenna2 := locations[idx2]

				antinodes[antenna1] = struct{}{}
				antinodes[antenna2] = struct{}{}

				for {
					antinode1 := [2]int{2*antenna1[0] - antenna2[0], 2*antenna1[1] - antenna2[1]}

					if _, exists := result.grid[antinode1]; exists {
						antinodes[antinode1] = struct{}{}
					} else {
						break
					}

					antenna2 = antenna1
					antenna1 = antinode1
				}

				antenna1 = locations[idx1]
				antenna2 = locations[idx2]

				for {
					antinode2 := [2]int{2*antenna2[0] - antenna1[0], 2*antenna2[1] - antenna1[1]}

					if _, exists := result.grid[antinode2]; exists {
						antinodes[antinode2] = struct{}{}
					} else {
						break
					}

					antenna1 = antenna2
					antenna2 = antinode2
				}
			}
		}
	}

	result.print(filename, antinodes)
	fmt.Println()
	result.print(filename, make(map[[2]int]struct{}))

	fmt.Println("solution day 08 part 02:", len(antinodes))
}
