package day5

import (
	"sort"

	"github.com/tomasff/aoc-2024/days"
)

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

		if isUpdateValid(update, rules) {
			validMiddlePageSum += update[middlePageIndex]
		} else {
			fixUpdate(update, rules)
			invalidButFixedMiddlePageSum += update[middlePageIndex]
		}
	}

	return validMiddlePageSum, invalidButFixedMiddlePageSum
}

func SolveDaySort(inputPath string) days.DaySolution {
	rules, updates := parseInput(inputPath)
	validMiddlePageSum, invalidMiddlePageSum := sumMiddlePagesSort(updates, rules)

	return days.DaySolution{
		PartOne: validMiddlePageSum,
		PartTwo: invalidMiddlePageSum,
	}
}
