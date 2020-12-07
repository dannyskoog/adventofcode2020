package day3part2

import (
	day1part1 "adventofcode2020/day1/part1"
	day3part1 "adventofcode2020/day3/part1"
)

func calculateProductOfEncounteredTreesInTobogganMapBySlopes(tobogganMap day3part1.TobogganMap, slopes []day3part1.TobogganMapSlope) int {
	encounteredSlopeTrees := []int{}

	for _, slope := range slopes {
		encounteredTrees := day3part1.CalculateNumberOfTreesEncounteredInTobogganMapBySlope(tobogganMap, slope)
		encounteredSlopeTrees = append(encounteredSlopeTrees, encounteredTrees)
	}

	return day1part1.CalculateProductOfNumbers(encounteredSlopeTrees)
}
