package day10part1

import (
	"sort"
)

func getProductOfOneAndThreeJoltRatingDifferenceCounts(joltRatings []int) int {
	joltRatingDifferences := getJoltRatingDifferences(joltRatings)

	return joltRatingDifferences[1] * joltRatingDifferences[3]
}

func getJoltRatingDifferences(joltRatings []int) map[int]int {
	joltRatingDifferences := make(map[int]int)
	previousJoltRating := 0 // Default value represents charging outlet rating

	sort.Ints(joltRatings)

	for _, joltRating := range joltRatings {
		joltRatingDifference := joltRating - previousJoltRating
		joltRatingDifferences[joltRatingDifference]++
		previousJoltRating = joltRating
	}

	// Add device jolt rating difference
	joltRatingDifferences[3]++

	return joltRatingDifferences
}
