package day4

import (
	"os"
	"strings"

	"github.com/tomasff/aoc-2024/days"
)

const xmas = "XMAS"

func countXmasOccurences(grid grid) int {
	count := 0

	for row := range grid.numRows() {
		for col := range grid.numColumns() {
			// No need to check neigbours if the centre is not correct.
			if grid.get(row, col) != xmas[0] {
				continue
			}

			for _, orientation := range allOrientations {
				slice := grid.orientedSlice(
					row, col, orientation[0], orientation[1], 4,
				)

				if slice == xmas {
					count++
				}
			}
		}
	}

	return count
}

func countXmasCrosses(grid grid) int {
	count := 0

	for row := range grid.numRows() {
		for col := range grid.numColumns() {
			if !isXmasCross(grid, row, col) {
				continue
			}

			count++
		}
	}

	return count
}

func isXmasCross(grid grid, row int, col int) bool {
	if !isValidCrossLocation(grid, row, col) {
		return false
	}

	diagonalOpposite := ((grid.get(row+1, col-1) == 'S' && grid.get(row-1, col+1) == 'M') ||
		(grid.get(row+1, col-1) == 'M' && grid.get(row-1, col+1) == 'S'))

	// X-MAS cross with first diagonal MS,
	//
	// M M      M S
	//  A   or   A
	// S S      M S
	msCross := grid.get(row-1, col-1) == 'M' && grid.get(row+1, col+1) == 'S' && diagonalOpposite

	// X-MAS cross with first diagonal SM,
	//
	// S M      S S
	//  A   or   A
	// S M      M M
	smCross := grid.get(row-1, col-1) == 'S' && grid.get(row+1, col+1) == 'M' && diagonalOpposite

	return msCross || smCross
}

func isValidCrossLocation(grid grid, row, col int) bool {
	return (grid.get(row, col) == 'A' &&
		grid.coordinatesAreValid(row-1, col-1) &&
		grid.coordinatesAreValid(row+1, col+1))
}

func loadGrid(inputPath string) grid {
	inputBytes, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	return grid{
		source: strings.Split(string(inputBytes), "\n"),
	}
}

func SolveDay(inputPath string) days.DaySolution {
	input := loadGrid(inputPath)

	return days.DaySolution{
		PartOne: countXmasOccurences(input),
		PartTwo: countXmasCrosses(input),
	}
}
