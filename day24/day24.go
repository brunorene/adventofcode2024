package day24

import (
	"adventofcode2024/common"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type and struct {
	inputs [2]string
	output string
}

type or struct {
	inputs [2]string
	output string
}

type xor struct {
	inputs [2]string
	output string
}

type computer struct {
	inputs     map[string]*bool
	operations []operation
}

type operation interface {
	apply(*computer)
	outputName() string
}

func parse(filename string) (result computer) {
	content, err := common.ReadInput("day24/" + filename)
	common.CheckError(err)

	result.inputs = make(map[string]*bool)

	for line := range content.ReadLines {
		if !strings.Contains(line, "->") {
			parts := strings.Split(line, ": ")

			result.inputs[parts[0]] = &[]bool{parts[1] == "1"}[0]
		} else {
			parts := strings.Split(line, " ")

			switch parts[1] {
			case "AND":
				result.operations = append(result.operations, and{
					inputs: [2]string{parts[0], parts[2]},
					output: parts[4],
				})
			case "OR":
				result.operations = append(result.operations, or{
					inputs: [2]string{parts[0], parts[2]},
					output: parts[4],
				})
			case "XOR":
				result.operations = append(result.operations, xor{
					inputs: [2]string{parts[0], parts[2]},
					output: parts[4],
				})

			}
		}
	}

	return result
}

func (a and) apply(comp *computer) {
	val1 := comp.inputs[a.inputs[0]]
	val2 := comp.inputs[a.inputs[1]]

	if val1 != nil && val2 != nil {
		comp.inputs[a.output] = &[]bool{*val1 && *val2}[0]
	}
}

func (o or) apply(comp *computer) {
	val1 := comp.inputs[o.inputs[0]]
	val2 := comp.inputs[o.inputs[1]]

	if val1 != nil && val2 != nil {
		comp.inputs[o.output] = &[]bool{*val1 || *val2}[0]
	}
}

func (x xor) apply(comp *computer) {
	val1 := comp.inputs[x.inputs[0]]
	val2 := comp.inputs[x.inputs[1]]

	if val1 != nil && val2 != nil {
		comp.inputs[x.output] = &[]bool{*val1 != *val2}[0]
	}
}

func (a and) outputName() string {
	return a.output
}

func (o or) outputName() string {
	return o.output
}

func (x xor) outputName() string {
	return x.output
}

func (c *computer) done() bool {
	for _, op := range c.operations {
		if c.inputs[op.outputName()] == nil {
			return false
		}
	}

	return true
}

func (c *computer) zNumber() (num int64) {
	for name, input := range c.inputs {
		if strings.HasPrefix(name, "z") && input != nil && *input {
			exp, err := strconv.Atoi(name[1:])
			common.CheckError(err)

			num += 1 << uint(exp)
		}
	}

	return num
}

func (c *computer) calculate() int64 {

	for !c.done() {
		for _, op := range c.operations {
			op.apply(c)
		}
	}

	return c.zNumber()
}

func (c *computer) setInputs(x, y int64) {

	binX := strconv.FormatInt(x, 2)
	binY := strconv.FormatInt(y, 2)

	for name := range c.inputs {
		if !strings.HasPrefix(name, "x") && !strings.HasPrefix(name, "y") {
			delete(c.inputs, name)

			continue
		}

		idx, err := strconv.Atoi(name[1:])
		common.CheckError(err)

		if strings.HasPrefix(name, "x") {
			if idx < len(binX) {
				c.inputs[name] = &[]bool{binX[len(binX)-idx-1] == '1'}[0]
			} else {
				c.inputs[name] = &[]bool{false}[0]
			}
		} else if strings.HasPrefix(name, "y") {
			if idx < len(binY) {
				c.inputs[name] = &[]bool{binY[len(binY)-idx-1] == '1'}[0]
			} else {
				c.inputs[name] = &[]bool{false}[0]
			}
		}
	}
}

func Solve1(filename string) {
	result := parse(filename)

	fmt.Println("solution day 24 part 01:", result.calculate())
}

func Solve2(filename string) {
	result := parse(filename)

	// switch z05, jst, gdf, mcm, z15, dnt, z30, gwc

	for idx, oper := range result.operations {
		switch oper.outputName() {
		case "z05":
			result.operations[idx] = or{
				inputs: [2]string{"sgt", "bhb"},
				output: "jst",
			}
		case "jst":
			result.operations[idx] = xor{
				inputs: [2]string{"ggh", "tvp"},
				output: "z05",
			}
		case "gdf":
			result.operations[idx] = xor{
				inputs: [2]string{"x10", "y10"},
				output: "mcm",
			}
		case "mcm":
			result.operations[idx] = and{
				inputs: [2]string{"x10", "y10"},
				output: "gdf",
			}
		case "dnt":
			result.operations[idx] = xor{
				inputs: [2]string{"vhr", "dvj"},
				output: "z15",
			}
		case "z15":
			result.operations[idx] = and{
				inputs: [2]string{"x15", "y15"},
				output: "dnt",
			}
		case "gwc":
			result.operations[idx] = xor{
				inputs: [2]string{"kgr", "vrg"},
				output: "z30",
			}
		case "z30":
			result.operations[idx] = and{
				inputs: [2]string{"kgr", "vrg"},
				output: "gwc",
			}
		}
	}

	for exp := range 45 {
		result.setInputs(0, int64(math.Pow(2, float64(exp))))
		if result.calculate() != int64(math.Pow(2, float64(exp))) {
			fmt.Println(exp, result.calculate(), int64(math.Pow(2, float64(exp))), strconv.FormatInt(int64(math.Pow(2, float64(exp))), 2))
		}
	}

	faultyPorts := []string{"z05", "jst", "gdf", "mcm", "z15", "dnt", "z30", "gwc"}

	sort.Strings(faultyPorts)

	fmt.Println("solution day 24 part 0:", strings.Join(faultyPorts, ","))
}
