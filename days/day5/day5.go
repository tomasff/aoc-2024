package day5

import (
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/tomasff/aoc-2024/days"
)

type rule struct {
	firstPage int
	thenPage  int
}

func parseOrderingRules(unparsedRules string) map[rule]bool {
	rules := make(map[rule]bool)

	for _, unparsedRule := range strings.Split(unparsedRules, "\n") {
		ruleParts := strings.Split(unparsedRule, "|")

		firstPage, err := strconv.Atoi(ruleParts[0])
		if err != nil {
			panic(err)
		}

		thenPage, err := strconv.Atoi(ruleParts[1])
		if err != nil {
			panic(err)
		}

		rules[rule{
			firstPage: firstPage,
			thenPage:  thenPage,
		}] = true
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

func parseInput(inputPath string) (map[rule]bool, [][]int) {
	inputBytes, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	inputParts := strings.Split(string(inputBytes), "\n\n")

	return parseOrderingRules(inputParts[0]), parseSafetyManualUpdates(inputParts[1])
}

func isUpdateValid(update []int, rules map[rule]bool) bool {
	return sort.SliceIsSorted(update, func(i, j int) bool {
		return rules[rule{
			firstPage: update[i],
			thenPage:  update[j],
		}]
	})
}

func sumCorrectUpdatesMiddlePage(updates [][]int, rules map[rule]bool) int {
	middlePageSum := 0

	for _, update := range updates {
		if !isUpdateValid(update, rules) {
			continue
		}

		middlePageSum += update[len(update)/2]
	}

	return middlePageSum
}

func fixUpdate(update []int, rules map[rule]bool) {
	sort.Slice(update, func(i, j int) bool {
		return rules[rule{
			firstPage: update[i],
			thenPage:  update[j],
		}]
	})
}

func sumFixedUpdatesMiddlePage(updates [][]int, rules map[rule]bool) int {
	middlePageSum := 0

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			continue
		}

		// In-place is OK for this exercise since this is the second part.
		// Aternatively, should fix the copied update.
		fixUpdate(update, rules)
		middlePageSum += update[len(update)/2]
	}

	return middlePageSum
}

// This solution only works because the AoC puzzle input is nice.
// For part 1, we could instead keep track of "previous pages" and intersect
// with the pages that should be after according to the rules to determine
// if a rule is invalid.
// For part 2, we can do a topological sort (assuming no cycles).
func SolveDay(inputPath string) days.DaySolution {
	rules, updates := parseInput(inputPath)

	return days.DaySolution{
		PartOne: sumCorrectUpdatesMiddlePage(updates, rules),
		PartTwo: sumFixedUpdatesMiddlePage(updates, rules),
	}
}
