package day5

import (
	"os"
	"strconv"
	"strings"
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
