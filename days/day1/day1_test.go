package day1

import (
	"testing"
)

var testList1 = []int{
	3, 4, 2, 1, 3, 3,
}
var testList2 = []int{
	4, 3, 5, 3, 9, 3,
}

const partOneSolution = 11
const partTwoSolution = 31

func TestPartOne(t *testing.T) {
	tentativeSolution := solvePartOne(testList1, testList2)

	if tentativeSolution != partOneSolution {
		t.Fatalf(
			`Expected part one solution %d, received %d`,
			tentativeSolution,
			partOneSolution,
		)
	}
}

func TestPartTwo(t *testing.T) {
	tentativeSolution := solvePartTwo(testList1, testList2)

	if tentativeSolution != partTwoSolution {
		t.Fatalf(
			`Expected part two solution %d, received %d`,
			tentativeSolution,
			partTwoSolution,
		)
	}
}
