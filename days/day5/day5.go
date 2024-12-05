package day5

import "github.com/tomasff/aoc-2024/days"

func SolveDayTopologicalSort(inputPath string) days.DaySolution {
	rules, updates := parseInput(inputPath)
	validMiddlePageSum, invalidMiddlePageSum := sumMiddlePagesTopologicalSort(updates, rules)

	return days.DaySolution{
		PartOne: validMiddlePageSum,
		PartTwo: invalidMiddlePageSum,
	}
}

func SolveDaySort(inputPath string) days.DaySolution {
	rules, updates := parseInput(inputPath)
	validMiddlePageSum, invalidMiddlePageSum := sumMiddlePagesSort(updates, rules)

	return days.DaySolution{
		PartOne: validMiddlePageSum,
		PartTwo: invalidMiddlePageSum,
	}
}
