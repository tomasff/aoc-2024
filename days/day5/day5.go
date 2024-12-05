package day5

import (
	"os"
	"strconv"
	"strings"

	"github.com/tomasff/aoc-2024/days"
)

type set = map[int]bool

func parseRule(unparsedRule string) (int, int) {
	ruleParts := strings.Split(unparsedRule, "|")

	firstPage, err := strconv.Atoi(ruleParts[0])
	if err != nil {
		panic(err)
	}

	thenPage, err := strconv.Atoi(ruleParts[1])
	if err != nil {
		panic(err)
	}

	return firstPage, thenPage
}

func parseOrderingRules(unparsedRules string) map[int]set {
	rules := make(map[int]set)

	for _, unparsedRule := range strings.Split(unparsedRules, "\n") {
		firstPage, thenPage := parseRule(unparsedRule)

		_, ok := rules[firstPage]
		if !ok {
			rules[firstPage] = make(map[int]bool)
		}

		rules[firstPage][thenPage] = true
	}

	return rules
}

func parseSafetyManualUpdates(rawUpdates string) [][]int {
	unparsedUpdates := strings.Split(rawUpdates, "\n")
	updates := make([][]int, 0, len(unparsedUpdates))

	for _, unparsedUpdate := range unparsedUpdates {
		unparsedPages := strings.Split(unparsedUpdate, ",")
		pages := make([]int, 0, len(unparsedPages))

		for _, unparsedPage := range unparsedPages {
			page, err := strconv.Atoi(unparsedPage)
			if err != nil {
				panic(err)
			}

			pages = append(pages, page)
		}

		updates = append(updates, pages)
	}

	return updates
}

func parseInput(inputPath string) (map[int]set, [][]int) {
	inputBytes, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	inputParts := strings.Split(string(inputBytes), "\n\n")

	return parseOrderingRules(inputParts[0]), parseSafetyManualUpdates(inputParts[1])
}

func isUpdateValid(update []int, rules map[int]set) bool {
	for i, firstPage := range update {
		for _, otherPage := range update[i+1:] {
			if rules[otherPage][firstPage] {
				return false
			}
		}
	}

	return true
}

func countPagesBefore(update []int, rules map[int]set) map[int]int {
	numPagesBefore := make(map[int]int)

	for _, page := range update {
		numPagesBefore[page] = 0
	}

	for i, page := range update {
		for j, otherPage := range update {
			if i != j && rules[page][otherPage] {
				numPagesBefore[otherPage]++
			}
		}
	}

	return numPagesBefore
}

// Finds the correct middle page of an invalid update after it has been fixed.
func getMiddlePageOfFixedUpdate(update []int, rules map[int]set) int {
	queue := make([]int, 0, len(update))
	numPagesBefore := countPagesBefore(update, rules)

	for page, count := range numPagesBefore {
		if count == 0 {
			queue = append(queue, page)
		}
	}

	middlePageIndex := len(update) / 2
	numProcessedPages := 0

	for len(queue) > 0 {
		currentPage := queue[0]
		queue = queue[1:]

		numProcessedPages++

		if numProcessedPages == middlePageIndex+1 {
			return currentPage
		}

		for nextPage := range rules[currentPage] {
			if _, nextPageInUpdate := numPagesBefore[nextPage]; !nextPageInUpdate {
				continue
			}

			numPagesBefore[nextPage]--
			if numPagesBefore[nextPage] == 0 {
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
