package day9part2

import (
	day9part1 "adventofcode2020/day9/part1"
)

func getEncryptionWeakness(numbers []int, preambleCount int) int {
	invalidNumber := day9part1.GetFirstInvalidNumber(numbers, preambleCount)
	contiguousNumbers := findContiguousNumbersBySum(numbers, invalidNumber)
	min, max := getMinAndMaxNumbers(contiguousNumbers)

	return min + max
}

func getMinAndMaxNumbers(numbers []int) (int, int) {
	min := numbers[0]
	max := min

	for _, number := range numbers {
		if number < min {
			min = number
		} else if number > max {
			max = number
		}
	}

	return min, max
}

func findContiguousNumbersBySum(numbers []int, sum int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		summedNumbers := []int{numbers[i]}
		accSum := numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			accSum += numbers[j]
			summedNumbers = append(summedNumbers, numbers[j])

			if accSum == sum {
				return summedNumbers
			} else if accSum > sum {
				break
			}
		}
	}

	return []int{}
}
