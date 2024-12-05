package day5

import (
	"sort"

	"github.com/tomasff/aoc-2024/days"
)

func isUpdateValidSort(update []int, rules map[int]set) bool {
	return sort.SliceIsSorted(update, func(i, j int) bool {
		return rules[update[i]][update[j]]
	})
}

func fixUpdate(update []int, rules map[int]set) {
	sort.Slice(update, func(i, j int) bool {
		return rules[update[i]][update[j]]
	})
}

func sumMiddlePagesSort(updates [][]int, rules map[int]set) (int, int) {
	validMiddlePageSum := 0
	invalidButFixedMiddlePageSum := 0

	for _, update := range updates {
		middlePageIndex := len(update) / 2

		if isUpdateValidSort(update, rules) {
			validMiddlePageSum += update[middlePageIndex]
		} else {
			fixUpdate(update, rules)
			invalidButFixedMiddlePageSum += update[middlePageIndex]
		}
	}

	return validMiddlePageSum, invalidButFixedMiddlePageSum
}

// This solution only works because the AoC puzzle input is nice.
// For part 1, we could instead keep track of "previous pages" and intersect
// with the pages that should be after according to the rules to determine
// if a rule is invalid.
// For part 2, we can do a topological sort (assuming no cycles).
//
// See day5.solveDay
func SolveDaySort(inputPath string) days.DaySolution {
	rules, updates := parseInput(inputPath)
	validMiddlePageSum, invalidMiddlePageSum := sumMiddlePagesSort(updates, rules)

	return days.DaySolution{
		PartOne: validMiddlePageSum,
		PartTwo: invalidMiddlePageSum,
	}
}
