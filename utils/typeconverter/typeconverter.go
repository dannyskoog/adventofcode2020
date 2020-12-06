package typeconverter

import (
	"log"
	"strconv"
)

// StringArrayToIntArray converts []string into []int
func StringArrayToIntArray(strArr []string) []int {
	var intArr = []int{}

	for _, i := range strArr {
		integer, err := strconv.Atoi(i)

		if err != nil {
			log.Fatal(err)
		}

		intArr = append(intArr, integer)
	}

	return intArr
}
