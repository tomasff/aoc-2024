package day6

import (
	"bytes"
	"os"

	"github.com/tomasff/aoc-2024/days"
)

const (
	guard      = '^'
	obstacle   = '#'
	outsideMap = '!'
)

type move struct {
	position    vector
	orientation vector
}

type vector struct {
	x, y int
}

func (v *vector) add(otherVector vector) vector {
	return vector{
		v.x + otherVector.x,
		v.y + otherVector.y,
	}
}

func (v *vector) rotate90() vector {
	return vector{
		v.y, -v.x,
	}
}

func (v *vector) rotate180() vector {
	return vector{
		-v.x, -v.y,
	}
}

func loadInput(inputPath string) [][]byte {
	inputBytes, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	return bytes.Split(inputBytes, []byte{'\n'})
}

func findGuardPosition(guardMap [][]byte) vector {
	for row, line := range guardMap {
		for column, entity := range line {
			if entity == guard {
				return vector{row, column}
			}
		}
	}

	return vector{-1, -1}
}

func withinMap(position vector, guardMap [][]byte) bool {
	numRows := len(guardMap)
	numColums := len(guardMap[0])

	return position.x >= 0 && position.x < numRows && position.y >= 0 && position.y < numColums
}

// Simulate guard movements and return true if the guard exits the map.
func simulateGuard(
	guardMap [][]byte, startPosition, orientation vector, onVisit func(vector, vector) bool,
) bool {
	currentPosition := startPosition

	for withinMap(currentPosition, guardMap) {
		if shouldContinue := onVisit(currentPosition, orientation); !shouldContinue {
			return false
		}

		positionForward := currentPosition.add(orientation)

		for getEntity(guardMap, positionForward) == obstacle {
			orientation = orientation.rotate90()
			positionForward = currentPosition.add(orientation)
		}

		currentPosition = positionForward
	}

	return true
}

func getEntity(guardMap [][]byte, position vector) byte {
	if !withinMap(position, guardMap) {
		return outsideMap
	}

	return guardMap[position.x][position.y]
}

func countLoopingObstacles(guardMap [][]byte, obstacles []move, guardPosition vector) int {
	loopingCount := 0

	for _, move := range obstacles {
		if move.position == guardPosition {
			continue
		}

		previousValue := guardMap[move.position.x][move.position.y]
		startPosition := move.position.add(move.orientation.rotate180())

		guardMap[move.position.x][move.position.y] = obstacle
		if guardLoops(guardMap, startPosition, move.orientation) {
			loopingCount++
		}

		guardMap[move.position.x][move.position.y] = previousValue
	}

	return loopingCount
}

func simulateGuardPath(guardMap [][]byte, guardPosition vector) []move {
	hasVisited := make(map[vector]bool)
	visited := make([]move, 0)

	simulateGuard(guardMap, guardPosition, vector{-1, 0}, func(position, orientation vector) bool {
		if !hasVisited[position] {
			visited = append(visited, move{position, orientation})
		}
		hasVisited[position] = true

		return true
	})

	return visited
}

func guardLoops(guardMap [][]byte, startPosition vector, orientation vector) bool {
	seen := make(map[move]bool)

	return !simulateGuard(guardMap, startPosition, orientation, func(position, orientation vector) bool {
		positionForward := position.add(orientation)

		if getEntity(guardMap, positionForward) == obstacle {
			currentMove := move{positionForward, orientation}
			if seen[currentMove] {
				return false
			}

			seen[currentMove] = true
		}

		return true
	})
}

// TODO(tomasff): A bit overengineered, refactor later.
func SolveDay(inputPath string) days.DaySolution {
	guardMap := loadInput(inputPath)

	guardPosition := findGuardPosition(guardMap)
	visited := simulateGuardPath(guardMap, guardPosition)

	return days.DaySolution{
		PartOne: len(visited),
		PartTwo: countLoopingObstacles(guardMap, visited, guardPosition),
	}
}
