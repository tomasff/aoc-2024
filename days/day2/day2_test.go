package day2

import "testing"

var reports = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

const numStrictSafeReports = 2
const numDampenedSafeReports = 4

func TestCountSafeReports(t *testing.T) {
	tentativeSolution := countSafeReports(reports)

	if tentativeSolution != numStrictSafeReports {
		t.Fatalf(
			`Expected %d strictly safe reports, counted %d.`,
			numStrictSafeReports,
			tentativeSolution,
		)
	}
}

func TestCountDampenedSafeReports(t *testing.T) {
	tentativeSolution := countSafeReportsWithDampener(reports)

	if tentativeSolution != numDampenedSafeReports {
		t.Fatalf(
			`Expected %d dampened safe reports, counted %d.`,
			numDampenedSafeReports,
			tentativeSolution,
		)
	}
}
