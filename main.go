package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomasff/aoc-2024/days"
	"github.com/tomasff/aoc-2024/days/day1"
)

type DaySolver func(string) days.DaySolution

var solvers = map[int]DaySolver{
	1: day1.SolveDay,
}

func main() {
	day := flag.Int("day", 1, "Day to be solved.")
	dayInputPath := flag.String("input", "", "Path to input for the specified day.")

	flag.Parse()

	if *dayInputPath == "" {
		fmt.Println("Invalid day input file path.")
		os.Exit(1)
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
