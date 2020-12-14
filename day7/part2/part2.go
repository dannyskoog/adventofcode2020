package day7part2

import (
	day7part1 "adventofcode2020/day7/part1"
)

func getInnerBagsCountByContainingBag(bags day7part1.BagMap, bagName string) int {
	count := getInnerBagsCount(bags, bags[bagName])

	return count
}

func getInnerBagsCount(bags day7part1.BagMap, innerBags []day7part1.InnerBag) int {
	count := 0

	for _, innerBag := range innerBags {
		count += innerBag.Amount + innerBag.Amount*getInnerBagsCount(bags, bags[innerBag.Name])
	}

	return count
}
