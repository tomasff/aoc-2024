package day3

import (
	"os"
	"strconv"

	"github.com/tomasff/aoc-2024/days"
)

func parseArg(text string, delimiter byte) (int, int) {
	currentIndex := 0

	for ; text[currentIndex] != delimiter && currentIndex < len(text); currentIndex++ {
		_, err := strconv.Atoi(string(text[currentIndex]))

		if err != nil {
			break
		}
	}

	if text[currentIndex] != delimiter {
		return 0, currentIndex
	}

	val, _ := strconv.Atoi(text[:currentIndex])

	currentIndex++

	return val, currentIndex
}

func computeMultiplications(instructions string) int {
	total := 0

	for currentCharIndex := 0; currentCharIndex < len(instructions); {
		if currentCharIndex < len(instructions)-4 && instructions[currentCharIndex:currentCharIndex+4] == "mul(" {
			currentCharIndex += 4

			firstArg, newIndex := parseArg(instructions[currentCharIndex:], ',')
			currentCharIndex += newIndex

			secondArg, newIndex := parseArg(instructions[currentCharIndex:], ')')
			currentCharIndex += newIndex

			total += firstArg * secondArg
		} else {
			currentCharIndex++
		}
	}

	return total
}

func computeToggleableMultiplications(instructions string) int {
	total := 0

	mulEnabled := true

	for currentCharIndex := 0; currentCharIndex < len(instructions); {
		if mulEnabled && currentCharIndex < len(instructions)-4 && instructions[currentCharIndex:currentCharIndex+4] == "mul(" {
			currentCharIndex += 4

			firstArg, newIndex := parseArg(instructions[currentCharIndex:], ',')
			currentCharIndex += newIndex

			secondArg, newIndex := parseArg(instructions[currentCharIndex:], ')')
			currentCharIndex += newIndex

			total += firstArg * secondArg
		} else if currentCharIndex < len(instructions)-7 && instructions[currentCharIndex:currentCharIndex+7] == "don't()" {
			mulEnabled = false
			currentCharIndex += 7
		} else if currentCharIndex < len(instructions)-4 && instructions[currentCharIndex:currentCharIndex+4] == "do()" {
			mulEnabled = true
			currentCharIndex += 4
		} else {
			currentCharIndex++
		}
	}

	return total
}

func loadInstructions(inputPath string) string {
	inputBytes, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	return string(inputBytes)
}

// TODO(tomasff): clean up solution.
func SolveDay(inputPath string) days.DaySolution {
	instructions := loadInstructions(inputPath)

	return days.DaySolution{
		PartOne: computeMultiplications(instructions),
		PartTwo: computeToggleableMultiplications(instructions),
	}
}
