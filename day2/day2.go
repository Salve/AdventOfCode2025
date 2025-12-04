package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Salve/AdventOfCode2025/inputs"
	"github.com/Salve/AdventOfCode2025/registry"
)

const day = 2

var input []byte

var example = []byte(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`)

func init() {
	registry.Register(day, Run)
}

func Run() {
	input = inputs.Input(day)
	part1()
	part2()
}

func part1() {
	ranges := strings.Split(strings.TrimRight(string(input), "\n"), ",")
	sum := 0
	for _, pair := range ranges {
		splitpair := strings.Split(pair, "-")
		start, _ := strconv.Atoi(splitpair[0])
		end, _ := strconv.Atoi(splitpair[1])
		//fmt.Printf("Range: %d-%d\n", start, end)
		for i := start; i <= end; i++ {
			if invalid(i) {
				//fmt.Printf("invalid: %d\n", i)
				sum += i
			}
		}
	}
	result := sum
	fmt.Printf("Part 1: %v\n", result)
}

func invalid(num int) bool {
	s := strconv.Itoa(num)
	if len(s)%2 != 0 {
		return false
	}
	left := s[:len(s)/2]
	right := s[len(s)/2:]
	return left == right
}

func invalid2(num int) bool {
	s := strconv.Itoa(num)
	for i := 1; i <= len(s)/2; i++ {
		if repeats(s, i) {
			return true
		}
	}
	return false
}

func repeats(s string, l int) bool {
	if len(s)%l != 0 || len(s) == 1 {
		return false
	}
	for i := 0; i < len(s); i += l {
		if s[i:i+l] != s[:l] {
			return false
		}
	}
	return true
}

func part2() {
	ranges := strings.Split(strings.TrimRight(string(input), "\n"), ",")
	sum := 0
	for _, pair := range ranges {
		splitpair := strings.Split(pair, "-")
		start, _ := strconv.Atoi(splitpair[0])
		end, _ := strconv.Atoi(splitpair[1])
		//fmt.Printf("Range: %d-%d\n", start, end)
		for i := start; i <= end; i++ {
			if invalid2(i) {
				//fmt.Printf("invalid: %d\n", i)
				sum += i
			}
		}
	}
	result := sum
	fmt.Printf("Part 2: %v\n", result)
}
