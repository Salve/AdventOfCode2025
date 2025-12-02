package day1

import (
	"fmt"
	"strconv"

	"github.com/Salve/AdventOfCode2025/inputs"
	"github.com/Salve/AdventOfCode2025/registry"
)

const day = 1

var input []byte

var example = []byte(`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
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
	d := dial{pos: 50}
	for _, l := range inputs.Lines(input) {
		d.turn(l)
	}
	result := d.landedZero
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	d := dial{pos: 50}
	for _, l := range inputs.Lines(input) {
		d.turn(l)
	}
	result := d.landedZero + d.passedZero
	fmt.Printf("Part 2: %v\n", result)
}

type dial struct {
	pos        int
	landedZero int
	passedZero int
}

func (d *dial) turn(instruction string) {
	v, _ := strconv.Atoi(instruction[1:])
	d.passedZero += v / 100
	v = v % 100
	passed := false
	wasZero := false
	if d.pos == 0 {
		wasZero = true
	}
	switch instruction[:1] {
	case "L":
		d.pos -= v
		if d.pos < 0 {
			d.pos += 100
			if !wasZero {
				passed = true
			}
		}
	case "R":
		d.pos += v
		if d.pos > 99 {
			d.pos -= 100
			if !wasZero {
				passed = true
			}
		}
	}
	if d.pos == 0 {
		d.landedZero++
	} else {
		if passed {
			d.passedZero++
		}
	}
	//fmt.Printf("%s: %d (%d, %d)\n", instruction, d.pos, d.landedZero, d.passedZero)
}
