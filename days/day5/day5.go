package day5

import (
	"github.com/tomasff/aoc-2024/days"
)

func countPageDependencies(update []int, rules map[int]set) map[int]int {
	numPagesBefore := make(map[int]int)

	for _, page := range update {
		numPagesBefore[page] = 0
	}

	for _, page := range update {
		for pageAfter := range rules[page] {
			if _, ok := numPagesBefore[pageAfter]; ok {
				numPagesBefore[pageAfter]++
			}
		}
	}

	return numPagesBefore
}

// Finds the correct middle page of an invalid update after it has been fixed.
func getMiddlePageOfFixedUpdate(update []int, rules map[int]set) int {
	queue := make([]int, 0, len(update))
	numPageDependencies := countPageDependencies(update, rules)

	// We only need to find the middle element.
	numRequiredProcessedPages := len(update)/2 + 1

	for page, count := range numPageDependencies {
		if count == 0 && len(queue) < numRequiredProcessedPages {
			queue = append(queue, page)
		}
	}

	numProcessedPages := 0

	for len(queue) > 0 {
		currentPage := queue[0]
		queue = queue[1:]

		numProcessedPages++

		if numProcessedPages == numRequiredProcessedPages {
			return currentPage
		}

		for nextPage := range rules[currentPage] {
			if _, nextPageInUpdate := numPageDependencies[nextPage]; !nextPageInUpdate {
				continue
			}

			numPageDependencies[nextPage]--
			if numPageDependencies[nextPage] == 0 {
				queue = append(queue, nextPage)
			}
		}
	}

	return -1
}

func sumMiddlePages(updates [][]int, rules map[int]set) (int, int) {
	validMiddlePageSum := 0
	invalidButFixedMiddlePageSum := 0

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			validMiddlePageSum += update[len(update)/2]
		} else {
			invalidButFixedMiddlePageSum += getMiddlePageOfFixedUpdate(update, rules)
		}
	}

	return validMiddlePageSum, invalidButFixedMiddlePageSum
}

func SolveDay(inputPath string) days.DaySolution {
	rules, updates := parseInput(inputPath)
	validMiddlePageSum, invalidMiddlePageSum := sumMiddlePages(updates, rules)

	return days.DaySolution{
		PartOne: validMiddlePageSum,
		PartTwo: invalidMiddlePageSum,
	}
}
