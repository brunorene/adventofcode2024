package day23

import (
	"adventofcode2024/common"
	"fmt"
	"sort"
	"strings"
)

type lan map[string]map[string]struct{}

func parse(filename string) (result lan) {
	input, err := common.ReadInput("day23/" + filename)
	common.CheckError(err)

	result = make(lan)

	for line := range input.ReadLines {
		pair := strings.Split(line, "-")

		if _, exists := result[pair[0]]; !exists {
			result[pair[0]] = make(map[string]struct{})
		}

		if _, exists := result[pair[1]]; !exists {
			result[pair[1]] = make(map[string]struct{})
		}

		result[pair[0]][pair[1]] = struct{}{}
		result[pair[1]][pair[0]] = struct{}{}
	}

	return result
}

func Solve1(filename string) {
	result := parse(filename)

	allConnected := make(map[[3]string]struct{})

	for node1, conns1 := range result {
		for node2 := range conns1 { // connected to node1
			for node3 := range result[node2] { // connected to node2
				if _, ok := result[node1][node3]; ok { // node3 connected to node 1
					if node1[0] == 't' || node2[0] == 't' || node3[0] == 't' {
						slice := []string{node1, node2, node3}
						sort.Strings(slice)

						allConnected[[3]string{slice[0], slice[1], slice[2]}] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Println("solution day 23 part 01:", len(allConnected))
}

type item struct {
	allConnected map[string]struct{}
	neighbours   map[string]struct{}
}

func contains(main map[string]struct{}, check map[string]struct{}) bool {
	for k := range check {
		if _, exists := main[k]; !exists {
			return false
		}
	}

	return true
}

func Solve2(filename string) {
	result := parse(filename)

	var stack []item

	for node1, neighbours := range result {
		stack = append(stack, item{map[string]struct{}{node1: {}}, neighbours})
	}

	largest := make(map[string]struct{})

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if len(current.allConnected) > len(largest) {
			largest = current.allConnected
		}

		for neighbour := range current.neighbours {
			if contains(result[neighbour], current.allConnected) {
				current.allConnected[neighbour] = struct{}{}

				for k := range current.allConnected {
					delete(current.neighbours, k)
				}

				stack = append(stack, current)
			}
		}
	}

	slice := make([]string, 0, len(largest))

	for k := range largest {
		slice = append(slice, k)
	}

	sort.Strings(slice)

	fmt.Println("solution day 23 part 02:", strings.Join(slice, ","))
}
