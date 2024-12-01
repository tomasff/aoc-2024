package day1

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/tomasff/aoc-2024/days"
)

func computeListDistance(list1 []int, list2 []int) int {
	// Ideally, this should not mutate the original input.
	// For this problem this is ok.
	slices.Sort(list1)
	slices.Sort(list2)

	total := 0

	for i, val1 := range list1 {
		difference := val1 - list2[i]

		// Note that this isn't always correct because the negative and positive
		// range represented by an int are not symmetric.
		difference = max(difference, -difference)

		total += difference
	}

	return total
}

func computeListSimilarity(list1 []int, list2 []int) int {
	frequencies := make(map[int]int)

	for _, num := range list2 {
		frequencies[num] += 1
	}

	score := 0
	for _, num := range list1 {
		score += num * frequencies[num]
	}

	return score
}

func loadDayInput(inputPath string) ([]int, []int) {
	input_bytes, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	input_lines := strings.Split(string(input_bytes), "\n")

	list1 := make([]int, 0, len(input_lines))
	list2 := make([]int, 0, len(input_lines))

	for _, line := range input_lines {
		parts := strings.Split(line, "   ")

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}

func SolveDay(inputPath string) days.DaySolution {
	list1, list2 := loadDayInput(inputPath)

	return days.DaySolution{
		PartOne: computeListDistance(list1, list2),
		PartTwo: computeListSimilarity(list1, list2),
	}
}
