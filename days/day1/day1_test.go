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

const trueListDistance = 11
const trueListSimilarity = 31

func TestListDistanceExample(t *testing.T) {
	tentativeSolution := computeListDistance(testList1, testList2)

	if tentativeSolution != trueListDistance {
		t.Fatalf(
			`Expected list distance %d, calculated %d.`,
			tentativeSolution,
			trueListDistance,
		)
	}
}

func TestListSimilarityExample(t *testing.T) {
	listSimilarity := computeListSimilarity(testList1, testList2)

	if listSimilarity != trueListSimilarity {
		t.Fatalf(
			`Expected similarity score %d, calculated %d.`,
			listSimilarity,
			trueListSimilarity,
		)
	}
}
