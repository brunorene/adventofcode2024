package day09

import (
	"adventofcode2024/common"
	"fmt"
	"os"
	"strconv"
)

type file struct {
	startAddress int
	length       int
}

func clean(memory []int) []int {
	memory = memory[:len(memory)-1]

	for memory[len(memory)-1] == -1 {
		memory = memory[:len(memory)-1]
	}

	return memory
}

func getMemory() (addresses []file, space []file, memory []int) {
	var id int
	isFile := true

	content, err := os.ReadFile("day09/input.txt")
	common.CheckError(err)

	for _, size := range content {
		if size == '\n' {
			break
		}

		number, err := strconv.Atoi(fmt.Sprintf("%c", size))
		common.CheckError(err)

		if isFile {
			addresses = append(addresses, file{
				startAddress: len(memory),
				length:       number,
			})
		} else {
			space = append(space, file{
				startAddress: len(memory),
				length:       number,
			})

		}

		for range number {
			if isFile {
				memory = append(memory, id)
			} else {
				memory = append(memory, -1)
			}
		}

		if isFile {
			id++
		}

		isFile = !isFile
	}

	return addresses, space, memory
}

func Solve1() {
	_, _, memory := getMemory()

	var address int

	for address < len(memory) {
		if memory[address] == -1 {
			memory[address] = memory[len(memory)-1]

			memory = clean(memory)
		}

		address++
	}

	var checksum int64

	for idx, value := range memory {
		checksum += int64(idx * value)
	}

	fmt.Println("solution day 09 part 01:", checksum)
}

func Solve2() {
	addresses, space, memory := getMemory()

	for idx := len(addresses) - 1; idx >= 0; idx-- {
		lastFile := addresses[idx]

		for idSlot, slot := range space {
			if slot.length >= lastFile.length && slot.startAddress < lastFile.startAddress {
				value := memory[lastFile.startAddress]

				for pos := range lastFile.length {
					memory[slot.startAddress+pos] = value
					memory[lastFile.startAddress+pos] = -1
				}

				space[idSlot].startAddress += lastFile.length
				space[idSlot].length -= lastFile.length
				break
			}
		}
	}

	var checksum int64

	for idx, value := range memory {
		if value < 0 {
			continue
		}

		checksum += int64(idx * value)
	}

	fmt.Println("solution day 09 part 02:", checksum)
}
