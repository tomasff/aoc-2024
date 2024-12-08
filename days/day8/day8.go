package day8

import (
	"os"
	"strings"

	"github.com/tomasff/aoc-2024/days"
)

type vector struct {
	x, y int
}

func (v *vector) subtract(other vector) vector {
	return vector{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v *vector) add(other vector) vector {
	return vector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

type antenna struct {
	location  vector
	frequency rune
}

const empty = '.'

func parseAntennas(inputPath string) ([]antenna, int, int) {
	bytes, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	antennas := make([]antenna, 0)
	unparsedAntennaMap := strings.Split(string(bytes), "\n")

	for x, line := range unparsedAntennaMap {
		for y, entity := range line {
			if entity == empty {
				continue
			}

			antennas = append(antennas, antenna{
				location: vector{x, y}, frequency: entity,
			})
		}
	}

	return antennas, len(unparsedAntennaMap), len(unparsedAntennaMap[0])
}

func isValidLocation(location vector, maximumX, maximumY int) bool {
	return location.x >= 0 && location.x < maximumX && location.y >= 0 && location.y < maximumY
}

func countAntinodes(
	antennasByFrequency map[rune][]antenna, maximumX, maximumY int, allAntiNodes bool,
) int {
	isAntinode := make(map[vector]bool)

	for _, antennaGroup := range antennasByFrequency {
		for i, a := range antennaGroup {
			for j, b := range antennaGroup {
				if i == j {
					continue
				}

				aToB := b.location.subtract(a.location)
				antinode := b.location

				if !allAntiNodes {
					antinode = antinode.add(aToB)
				}

				for isValidLocation(antinode, maximumX, maximumY) {
					isAntinode[antinode] = true
					if !allAntiNodes {
						break
					}

					antinode = antinode.add(aToB)
				}
			}
		}
	}

	return len(isAntinode)
}

func groupAntennasByFrequency(antennas []antenna) map[rune][]antenna {
	antennasIndex := make(map[rune][]antenna)

	for _, a := range antennas {
		if _, ok := antennasIndex[a.frequency]; !ok {
			antennasIndex[a.frequency] = make([]antenna, 0)
		}

		antennasIndex[a.frequency] = append(antennasIndex[a.frequency], a)
	}

	return antennasIndex
}

// TODO(tomasff): Simplify vector abstraction?
func SolveDay(inputPath string) days.DaySolution {
	antennas, maximumX, maximumY := parseAntennas(inputPath)
	antennasByFrequency := groupAntennasByFrequency(antennas)

	return days.DaySolution{
		PartOne: countAntinodes(antennasByFrequency, maximumX, maximumY, false),
		PartTwo: countAntinodes(antennasByFrequency, maximumX, maximumY, true),
	}
}
