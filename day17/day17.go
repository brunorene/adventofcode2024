package day17

import (
	"adventofcode2024/common"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

type instruction struct {
	opcode  int64
	operand int64
}

type computer struct {
	regA, regB, regC int64
	program          []instruction
}

func (c *computer) instructions() (list []string) {
	for _, instr := range c.program {
		list = append(list, fmt.Sprintf("%d", instr.opcode), fmt.Sprintf("%d", instr.operand))
	}

	return list
}

func (c *computer) lastOutput() int64 {
	if c.program[len(c.program)-1].operand < 0 {
		return c.program[len(c.program)-1].opcode
	}

	return c.program[len(c.program)-1].operand
}

func (c *computer) removeOne() []instruction {
	if c.program[len(c.program)-1].operand < 0 {
		return c.program[:len(c.program)-1]
	}

	c.program[len(c.program)-1].operand = -1

	return c.program
}

func (c *computer) realLen() int {
	if c.program[len(c.program)-1].operand < 0 {
		return len(c.program) - 1
	}

	return len(c.program)
}

func (c *computer) value(operand int64) int64 {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	}

	return -1
}

func (c *computer) valuePrint(operand int) string {
	switch operand {
	case 0, 1, 2, 3:
		return fmt.Sprintf("literal: %d", operand)
	case 4:
		return fmt.Sprintf("regA: %d", c.regA)
	case 5:
		return fmt.Sprintf("regB: %d", c.regB)
	case 6:
		return fmt.Sprintf("regC: %d", c.regC)
	}

	return ""
}

func (c *computer) execute() (output []string) {
	//fmt.Println(strconv.FormatInt(int64(c.regA), 2))
	//fmt.Println(c.regA, c.regB, c.regC)

	var pointer int64
	for {
		if pointer >= int64(len(c.program)) {
			break
		}

		instr := c.program[pointer]

		switch instr.opcode {
		case adv:
			//fmt.Printf("regA = regA: %d >> %s\n", c.regA, c.valuePrint(instr.operand))
			c.regA >>= c.value(instr.operand)
		case bxl:
			//fmt.Printf("regB = regB: %d XOR literal: %d\n", c.regB, instr.operand)
			c.regB ^= instr.operand
		case bst:
			//fmt.Printf("regB = %s AND 111\n", c.valuePrint(instr.operand))
			c.regB = c.value(instr.operand) & 7
		case jnz:
			//fmt.Println()
			if c.regA != 0 {
				pointer = instr.operand

				continue
			}
		case bxc:
			//fmt.Printf("regB = regB: %d XOR regC: %d\n", c.regB, c.regC)
			c.regB = c.regC ^ c.regB
		case out:
			//fmt.Printf("OUT %s AND 111\n", c.valuePrint(instr.operand))
			output = append(output, fmt.Sprintf("%d", c.value(instr.operand)&7))
		case bdv:
			//fmt.Printf("regB = regA: %d >> %s\n", c.regA, c.valuePrint(instr.operand))
			c.regB = c.regA >> c.value(instr.operand)
		case cdv:
			//fmt.Printf("regC = regA: %d >> %s\n", c.regA, c.valuePrint(instr.operand))
			c.regC = c.regA >> c.value(instr.operand)
		}

		pointer++
	}

	//fmt.Println(c.regA, c.regB, c.regC)

	return output
}

func parse(filename string) (result computer) {
	input, err := common.ReadInput("day17/" + filename)
	common.CheckError(err)

	numMatch := regexp.MustCompile(`\d+`)

	var lineNum int

	var regA, regB, regC int64
	var program []instruction

	for line := range input.ReadLines {
		switch lineNum {

		case 0:
			_, err = fmt.Sscanf(line, "Register A: %d", &regA)
			common.CheckError(err)
		case 1:
			_, err = fmt.Sscanf(line, "Register B: %d", &regB)
			common.CheckError(err)
		case 2:
			_, err = fmt.Sscanf(line, "Register C: %d", &regC)
			common.CheckError(err)
		case 3:
			instr := numMatch.FindAllString(line, -1)
			common.CheckError(err)

			for i := 0; i < len(instr); i += 2 {
				opcode, err := strconv.ParseInt(instr[i], 10, 64)
				common.CheckError(err)

				operand, err := strconv.ParseInt(instr[i+1], 10, 64)
				common.CheckError(err)

				program = append(program, instruction{opcode, operand})
			}
		}

		lineNum++
	}

	return computer{
		regA:    regA,
		regB:    regB,
		regC:    regC,
		program: program,
	}
}

func Solve1(filename string) {
	result := parse(filename)

	output := result.execute()

	fmt.Println("solution day 17 part 01:", strings.Join(output, ","))
}

func Solve2(filename string) {
	result := parse(filename)
	target := result.instructions()

	regA := int64(math.Pow(8, float64(len(target)-1)))
	exp := len(target) - 1

	for {
		step := int64(math.Pow(8, float64(exp)))
		current := computer{regA: regA, program: result.program}
		output := current.execute()

		if slices.Compare(target, output) == 0 {
			break
		}

		fmt.Printf("%d - %d - %v - %v\n", exp, regA, output, target)

		if output[exp] == target[exp] {
			exp--

			continue
		}

		regA += step
	}

	fmt.Println("solution day 17 part 02:", regA)
}
