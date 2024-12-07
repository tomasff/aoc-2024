package day7

import (
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/tomasff/aoc-2024/days"
)

type equation struct {
	target int
	values []int
}

type candidate struct {
	currentValueIndex int
	currentTotal      int
}

func parseEquation(unparsedEquation string) equation {
	unparsedEquationParts := strings.Split(unparsedEquation, ": ")

	target, err := strconv.Atoi(unparsedEquationParts[0])
	if err != nil {
		panic(err)
	}

	unparsedValues := strings.Split(unparsedEquationParts[1], " ")
	values := make([]int, 0, len(unparsedValues))

	for _, unparsedValue := range unparsedValues {
		value, err := strconv.Atoi(unparsedValue)
		if err != nil {
			panic(err)
		}

		values = append(values, value)
	}

	return equation{target, values}
}

func loadEquations(inputPath string) []equation {
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	unparsedEquations := strings.Split(string(bytes), "\n")
	equations := make([]equation, 0, len(unparsedEquations))

	for _, unparsedEquation := range unparsedEquations {
		equations = append(equations, parseEquation(unparsedEquation))
	}

	return equations
}

func concat(a, b int) int {
	shift := int(math.Pow10(len(strconv.Itoa(b))))

	return a*shift + b
}

func canSolveEquation(equation equation, withConcat bool) bool {
	candidates := make([]candidate, 0)
	candidates = append(candidates, candidate{0, equation.values[0]})

	for len(candidates) > 0 {
		currentCandidate := candidates[0]
		candidates = candidates[1:]

		if currentCandidate.currentTotal > equation.target {
			continue
		}

		if currentCandidate.currentValueIndex == len(equation.values)-1 {
			if currentCandidate.currentTotal == equation.target {
				return true
			}

			continue
		}

		nextIndex := currentCandidate.currentValueIndex + 1
		nextValue := equation.values[nextIndex]

		candidates = append(candidates, candidate{
			nextIndex,
			currentCandidate.currentTotal * nextValue,
		})
		candidates = append(candidates, candidate{
			nextIndex,
			currentCandidate.currentTotal + nextValue,
		})

		if withConcat {
			candidates = append(candidates, candidate{
				nextIndex,
				concat(currentCandidate.currentTotal, nextValue),
			})
		}
	}

	return false
}

func computeCalibrationScore(equations []equation, withConcat bool) int {
	calibrationScore := 0

	for _, equation := range equations {
		if canSolveEquation(equation, withConcat) {
			calibrationScore += equation.target
		}
	}

	return calibrationScore
}

// TODO(tomasff): Improve complexity by starting from the equation solution.
func SolveDay(inputPath string) days.DaySolution {
	equations := loadEquations(inputPath)

	return days.DaySolution{
		PartOne: computeCalibrationScore(equations, false),
		PartTwo: computeCalibrationScore(equations, true),
	}
}
