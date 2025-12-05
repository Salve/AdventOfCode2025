package day4

import (
	"fmt"
	"image"

	"github.com/Salve/AdventOfCode2025/inputs"
	"github.com/Salve/AdventOfCode2025/registry"
)

const day = 4

var input []byte

var example = []byte(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`)

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

var dirs = []image.Point{
	image.Pt(-1, -1),
	image.Pt(-1, 0),
	image.Pt(-1, 1),
	image.Pt(0, -1),
	image.Pt(0, 1),
	image.Pt(1, -1),
	image.Pt(1, 0),
	image.Pt(1, 1),
}

func part1() {
	l := layout(inputs.Lines(input))
	result := remove(l)
	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	l := layout(inputs.Lines(input))
	result := 0
	for {
		removed := remove(l)
		if removed == 0 {
			break
		}
		result += removed
	}
	fmt.Printf("Part 2: %v\n", result)
}

func remove(l map[image.Point]struct{}) int {
	var rm []image.Point
	for roll := range l {
		adjacent := 0
		for _, dir := range dirs {
			if _, exists := l[roll.Add(dir)]; exists {
				adjacent++
			}
		}
		if adjacent < 4 {
			rm = append(rm, roll)
		}
	}
	for _, roll := range rm {
		delete(l, roll)
	}
	return len(rm)
}

func layout(i []string) map[image.Point]struct{} {
	o := make(map[image.Point]struct{})
	for x, line := range i {
		for y, char := range line {
			if char == '@' {
				o[image.Pt(x, y)] = struct{}{}
			}
		}
	}
	return o
}
