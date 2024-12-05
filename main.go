package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomasff/aoc-2024/days"
	"github.com/tomasff/aoc-2024/days/day1"
	"github.com/tomasff/aoc-2024/days/day2"
	"github.com/tomasff/aoc-2024/days/day3"
	"github.com/tomasff/aoc-2024/days/day4"
	"github.com/tomasff/aoc-2024/days/day5"
)

type DaySolver func(string) days.DaySolution

var solvers = map[int]DaySolver{
	1: day1.SolveDay,
	2: day2.SolveDay,
	3: day3.SolveDay,
	4: day4.SolveDay,
	5: day5.SolveDaySort,
}

func main() {
	day := flag.Int("day", 1, "Day to be solved.")
	dayInputPath := flag.String("input", "", "Path to input for the specified day.")

	flag.Parse()

	if *dayInputPath == "" {
		*dayInputPath = fmt.Sprintf("input/day%d.txt", *day)
	}

	daySolver, ok := solvers[*day]
	if !ok {
		fmt.Printf("No solver found for day %d\n", *day)
		os.Exit(1)
	}

	solution := daySolver(*dayInputPath)

	fmt.Printf(
		"Solution for day %d: Part one %d, Part two %d\n",
		*day,
		solution.PartOne,
		solution.PartTwo,
	)
}
