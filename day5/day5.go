package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Salve/AdventOfCode2025/inputs"
	"github.com/Salve/AdventOfCode2025/registry"
)

const day = 5

var input []byte

var example = []byte(`3-5
10-14
16-20
12-18

1
5
8
11
17
32
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
	d := parseInput(string(input))
	result := 0

	for _, ingredient := range d.available {
		for _, r := range d.freshRanges {
			if ingredient >= r[0] && ingredient <= r[1] {
				result++
				break
			}
		}
	}

	fmt.Printf("Part 1: %v\n", result)
}

func part2() {
	d := parseInput(string(input))

	sort.Slice(d.freshRanges, func(i, j int) bool {
		return d.freshRanges[i][0] < d.freshRanges[j][0]
	})

	processed := [][2]int{}
	for _, rng := range d.freshRanges {
		overlapped := false
		for i, prng := range processed {
			if merged, overlap := merge(rng, prng); overlap {
				processed[i] = merged
				overlapped = true
				break
			}
		}
		if !overlapped {
			processed = append(processed, rng)
		}
	}
	result := 0
	for _, rng := range processed {
		result += rng[1] - rng[0] + 1
	}
	fmt.Printf("Part 2: %v\n", result)
}

func merge(a, b [2]int) (merged [2]int, overlap bool) {
	if a[1] < b[0] || b[1] < a[0] {
		return [2]int{}, false
	}
	return [2]int{min(a[0], b[0]), max(a[1], b[1])}, true
}

func parseInput(input string) (d data) {
	s := strings.Split(input, "\n\n")
	for _, line := range strings.Split(s[0], "\n") {
		splitrng := strings.Split(line, "-")
		start, _ := strconv.Atoi(splitrng[0])
		end, _ := strconv.Atoi(splitrng[1])
		d.freshRanges = append(d.freshRanges, [2]int{start, end})
	}
	for _, line := range strings.Split(strings.TrimRight(s[1], "\n"), "\n") {
		v, _ := strconv.Atoi(line)
		d.available = append(d.available, v)
	}
	return d
}

type data struct {
	freshRanges [][2]int
	available   []int
}
