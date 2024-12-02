package day2

import (
	"os"
	"strconv"
	"strings"

	"github.com/tomasff/aoc-2024/days"
	"github.com/tomasff/aoc-2024/util"
)

func isReportSafe(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	isIncreasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		levelDelta := util.Abs(report[i] - report[i-1])

		if isIncreasing && report[i] <= report[i-1] {
			return false
		}

		if !isIncreasing && report[i] >= report[i-1] {
			return false
		}

		if levelDelta > 3 || levelDelta < 1 {
			return false
		}
	}

	return true
}

func parseReport(report string) []int {
	unparsedLevels := strings.Split(report, " ")
	levels := make([]int, 0, len(unparsedLevels))

	for _, unparsedLevel := range unparsedLevels {
		level, err := strconv.Atoi(unparsedLevel)

		if err != nil {
			panic(err)
		}

		levels = append(levels, level)
	}

	return levels
}

func loadReports(inputPath string) [][]int {
	inputBytes, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	inputLines := strings.Split(string(inputBytes), "\n")
	reports := make([][]int, 0, len(inputLines))

	for _, unparsedReport := range inputLines {
		reports = append(reports, parseReport(unparsedReport))
	}

	return reports
}

func countSafeReports(reports [][]int) int {
	numSafeReports := 0

	for _, report := range reports {
		if isReportSafe(report) {
			numSafeReports += 1
		}
	}

	return numSafeReports
}

func countSafeReportsWithDampener(reports [][]int) int {
	numSafeReports := 0

	// TODO(tomasff): Do it in better than O(n^2).
	for _, report := range reports {
		dampenedReport := make([]int, len(report)-1)

		for skipIndex := range report {
			copy(dampenedReport, report[:skipIndex])
			copy(dampenedReport[skipIndex:], report[skipIndex+1:])

			if isReportSafe(dampenedReport) {
				numSafeReports += 1
				break
			}
		}
	}

	return numSafeReports
}

func SolveDay(inputPath string) days.DaySolution {
	reports := loadReports(inputPath)

	return days.DaySolution{
		PartOne: countSafeReports(reports),
		PartTwo: countSafeReportsWithDampener(reports),
	}
}
