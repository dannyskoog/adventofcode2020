package day7part1

import (
	"regexp"
	"strconv"
	"strings"
)

type InnerBag struct {
	Name   string
	Amount int
}

type bag struct {
	name      string
	innerBags []InnerBag
}

type BagMap map[string][]InnerBag

func getBagNamesByContainedBag(bags BagMap, bagName string) []string {
	bagNames := []string{}

	for name, innerBags := range bags {
		containsBag := areInnerBagsContainingBag(bags, innerBags, bagName)

		if containsBag {
			bagNames = append(bagNames, name)
		}
	}

	return bagNames
}

func areInnerBagsContainingBag(bags BagMap, innerBags []InnerBag, bagName string) bool {
	for _, innerBag := range innerBags {
		if innerBag.Name == bagName || areInnerBagsContainingBag(bags, bags[innerBag.Name], bagName) {
			return true
		}
	}

	return false
}

func ConvertStrArrToBags(strArr []string) BagMap {
	bags := make(BagMap)

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
		innerBags: []InnerBag{},
	}

	innerBagsRegex, _ := regexp.Compile(`(\d)\s([a-z]+\s[a-z]+)\s[a-z]+`)
	innerBagsMatches := innerBagsRegex.FindAllStringSubmatch(str, -1)

	for _, groups := range innerBagsMatches {
		name := groups[2]
		amount, _ := strconv.Atoi(groups[1])

		innerBag := InnerBag{
			name,
			amount,
		}
		bag.innerBags = append(bag.innerBags, innerBag)
	}

	return bag
}
