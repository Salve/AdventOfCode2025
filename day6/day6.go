package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Salve/AdventOfCode2025/inputs"
	"github.com/Salve/AdventOfCode2025/registry"
)

const day = 6

var input []byte

var example = []byte(`123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
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
	lines := inputs.Lines(input)
	var splitlines [][]string
	for _, line := range lines {
		splitlines = append(splitlines, strings.Fields(line))
	}
	result := 0
	for col := 0; col < len(splitlines[0]); col++ {
		var f func(a, b int) int
		switch splitlines[len(splitlines)-1][col] {
		case "*":
			f = func(a, b int) int { return max(a, 1) * b }
		case "+":
			f = func(a, b int) int { return a + b }
		}
		r := 0
		for ln := 0; ln < len(splitlines)-1; ln++ {
			v, _ := strconv.Atoi(splitlines[ln][col])
			r = f(r, v)
		}
		result += r
	}

	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	lines := inputs.Lines(input)
	var fu func(a, b int) int
	var vs []int
	result := 0
	for col := 0; col < len(lines[0]); col++ {
		v, f, isSpace := readCol(lines, col)
		if f != nil {
			fu = f
		}
		if v != 0 {
			vs = append(vs, v)
		}
		if isSpace {
			st := 0
			for _, u := range vs {
				st = fu(st, u)
				fmt.Printf("Value: %d, subtotal: %d\n", u, st)
			}
			result += st
			fmt.Printf("Result inc by: %v, now %v\n\n", st, result)
			vs = []int{}
			fu = nil
			continue
		}
	}

	fmt.Printf("Part 2: %v\n", result)
}

func readCol(lines []string, col int) (v int, f func(int, int) int, isSpace bool) {
	switch lines[len(lines)-1][col] {
	case '*':
		f = func(a, b int) int { return max(a, 1) * b }
	case '+':
		f = func(a, b int) int { return a + b }
	}
	s := strings.Builder{}
	for line := range lines {
		c := lines[line][col]
		if c >= '0' && c <= '9' {
			s.WriteByte(c)
		}
	}
	sv := s.String()
	if sv == "" {
		isSpace = true
		return
	}
	v, _ = strconv.Atoi(sv)
	if col == len(lines[0])-1 {
		isSpace = true
	}
	return
}
