package day7part1

import (
	"regexp"
	"strconv"
	"strings"
)

type innerBag struct {
	name   string
	amount int
}

type bag struct {
	name      string
	innerBags []innerBag
}

type bagMap map[string][]innerBag

func getBagNamesByContainedBag(bags bagMap, bagName string) []string {
	bagNames := []string{}

	for name, innerBags := range bags {
		containsBag := areInnerBagsContainingBag(bags, innerBags, bagName)

		if containsBag {
			bagNames = append(bagNames, name)
		}
	}

	return bagNames
}

func areInnerBagsContainingBag(bags bagMap, innerBags []innerBag, bagName string) bool {
	for _, innerBag := range innerBags {
		if innerBag.name == bagName || areInnerBagsContainingBag(bags, bags[innerBag.name], bagName) {
			return true
		}
	}

	return false
}

func convertStrArrToBags(strArr []string) bagMap {
	bags := make(bagMap)

	for _, str := range strArr {
		bag := convertStrToBag(str)
		bags[bag.name] = bag.innerBags
	}

	return bags
}

func convertStrToBag(str string) bag {
	splitted := strings.Split(str, " bags contain ")

	bag := bag{
		name:      splitted[0],
		innerBags: []innerBag{},
	}

	innerBagsRegex, _ := regexp.Compile(`(\d)\s([a-z]+\s[a-z]+)\s[a-z]+`)
	innerBagsMatches := innerBagsRegex.FindAllStringSubmatch(str, -1)

	for _, groups := range innerBagsMatches {
		name := groups[2]
		amount, _ := strconv.Atoi(groups[1])

		innerBag := innerBag{
			name,
			amount,
		}
		bag.innerBags = append(bag.innerBags, innerBag)
	}

	return bag
}
