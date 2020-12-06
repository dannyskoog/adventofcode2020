package day1part2

import (
	day1part1 "adventofcode2020/day1/part1"
	"sort"
)

// CalculateProductOfThreeNumbersBySum calculates the product of two numbers who's sum matches the criteria
func CalculateProductOfThreeNumbersBySum(numbers []int, sum int) int {
	foundNumbers := findThreeNumbersBySum(numbers, sum)
	product := day1part1.CalculateProductOfNumbers(foundNumbers)

	return product
}

func findThreeNumbersBySum(numbers []int, sum int) []int {
	sort.Ints(numbers)
	rightIndex := len(numbers) - 1

	for i := 0; i < len(numbers)-2; i++ {
		leftIndex := i + 1
		for leftIndex < rightIndex {
			if numbers[i]+numbers[leftIndex]+numbers[rightIndex] == sum {
				return []int{numbers[i], numbers[leftIndex], numbers[rightIndex]}
			} else if numbers[i]+numbers[leftIndex]+numbers[rightIndex] < sum {
				leftIndex++
			} else {
				rightIndex--
			}
		}

	}

	return []int{}
}
