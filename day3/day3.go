package day3

import (
	"fmt"
	"math"

	"github.com/Salve/AdventOfCode2025/inputs"
	"github.com/Salve/AdventOfCode2025/registry"
)

const day = 3

var input []byte

var example = []byte(`987654321111111
811111111111119
234234234234278
818181911112111
`)

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

func part1() {
	result := 0
	for _, line := range inputs.Lines(input) {
		result += maxJolt(line, 2)
	}
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	result := 0
	for _, line := range inputs.Lines(input) {
		result += maxJolt(line, 12)
	}
	fmt.Printf("Part 2: %v\n", result)
}

func maxJolt(line string, digits int) int {
	maxVals := make([]int, digits)
	for i, char := range line {
		v := int(char - '0')
		for j := 0; j < digits; j++ {
			if v > maxVals[j] && i <= len(line)-digits+j {
				maxVals[j] = v
				for k := j + 1; k < digits; k++ {
					maxVals[k] = 0
				}
				break
			}
		}
	}
	result := 0
	for p, val := range maxVals {
		exp := digits - p - 1
		inc := val * int(math.Pow10(exp))
		result += inc
		fmt.Printf("Position %d: %d * (10^%d) = %d\n", p, val, exp, inc)
	}
	fmt.Printf("Max jolts: %v\n\n", result)
	return result
}
