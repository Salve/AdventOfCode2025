package main

import (
	"fmt"
	"time"

	_ "github.com/Salve/AdventOfCode2025/day1"
	_ "github.com/Salve/AdventOfCode2025/day2"
	_ "github.com/Salve/AdventOfCode2025/day3"
	_ "github.com/Salve/AdventOfCode2025/day4"
	_ "github.com/Salve/AdventOfCode2025/day5"
	_ "github.com/Salve/AdventOfCode2025/day6"
	"github.com/Salve/AdventOfCode2025/registry"
)

func main() {
	name, f := registry.Last()
	fmt.Printf("--- Running last day (%d) ---\n", name)
	timeFunc(f)
}

func timeFunc(f func()) time.Duration {
	start := time.Now()
	f()
	d := time.Now().Sub(start)
	fmt.Printf("--- Execution time: %s ---\n", d)
	return d
}
