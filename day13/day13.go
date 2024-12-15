package day13

import (
	"adventofcode2024/common"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

type variables struct {
	x, y *big.Int
}

type buttons struct {
	a, b, prize variables
}

func (b buttons) validate(buttonA, buttonB *big.Int) bool {
	return new(big.Int).Add(new(big.Int).Mul(buttonA, b.a.x), new(big.Int).Mul(buttonB, b.b.x)).Cmp(b.prize.x) == 0 &&
		new(big.Int).Add(new(big.Int).Mul(buttonA, b.a.y), new(big.Int).Mul(buttonB, b.b.y)).Cmp(b.prize.y) == 0
}

type state int

const (
	buttonA state = iota
	buttonB
	prize
)

func parse(filename string) (result []buttons) {
	input, err := common.ReadInput("day13/" + filename)
	common.CheckError(err)

	currentState := buttonA
	var current buttons

	buttonVal := regexp.MustCompile(`\d+`)

	for line := range input.ReadLines {
		vals := buttonVal.FindAllString(line, -1)

		var x, y int64

		if len(vals) == 2 {
			x, err = strconv.ParseInt(vals[0], 10, 64)
			common.CheckError(err)

			y, err = strconv.ParseInt(vals[1], 10, 64)
			common.CheckError(err)
		}

		switch currentState {
		case buttonA:
			current.a.x = big.NewInt(x)
			current.a.y = big.NewInt(y)
		case buttonB:
			current.b.x = big.NewInt(x)
			current.b.y = big.NewInt(y)
		case prize:
			current.prize.x = big.NewInt(x)
			current.prize.y = big.NewInt(y)

			result = append(result, current)

			current = buttons{}
		}

		currentState = (currentState + 1) % 3
	}

	return result
}

func Solve(filename string, add *big.Int) {
	result := parse(filename)

	tokens := new(big.Int)

	for _, machine := range result {
		machine.prize.x.Add(machine.prize.x, add)
		machine.prize.y.Add(machine.prize.y, add)

		// Cramer
		det := new(big.Int).Sub(new(big.Int).Mul(machine.a.x, machine.b.y), new(big.Int).Mul(machine.a.y, machine.b.x))

		buttonA := new(big.Int).Div(new(big.Int).Sub(new(big.Int).Mul(machine.prize.x, machine.b.y), new(big.Int).Mul(machine.prize.y, machine.b.x)), det)
		buttonB := new(big.Int).Div(new(big.Int).Sub(new(big.Int).Mul(machine.prize.y, machine.a.x), new(big.Int).Mul(machine.prize.x, machine.a.y)), det)

		if !machine.validate(buttonA, buttonB) {
			continue
		}

		tokens.Add(tokens, new(big.Int).Add(new(big.Int).Mul(buttonA, big.NewInt(3)), buttonB))
	}

	fmt.Println("solution day 13:", tokens)
}
