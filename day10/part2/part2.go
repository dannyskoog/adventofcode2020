package day10part2

import (
	"sort"
)

func getNumberOfJoltRatingArrangements(joltRatings []int) int {
	sort.Ints(joltRatings)
	maxJoltRating := joltRatings[len(joltRatings)-1]
	joltRatingsWithChargingOutletAndDevice := append([]int{0}, append(joltRatings, maxJoltRating+3)...)

	arrangementsCount := 1

	for i, currentOffset := 0, 1; i < len(joltRatingsWithChargingOutletAndDevice); i++ {
		if containsNumber(joltRatings, joltRatingsWithChargingOutletAndDevice[i]+1) {
			currentOffset++
		} else {
			arrangementsCount *= getTribonacci(currentOffset)
			currentOffset = 1
		}
	}

	return arrangementsCount
}

func getTribonacci(num int) int {
	tribonacciSequence := []int{1, 1, 2, 4, 7, 13, 24, 44, 81, 149}
	return tribonacciSequence[num-1]
}

func containsNumber(numbers []int, number int) bool {
	for i := range numbers {
		if numbers[i] == number {
			return true
		}
	}

	return false
}
