package day1

import (
	"adventofcode2020/utils"
	"log"
	"strconv"
)

func getNumbersFromTextFile(path string, fileName string) []int {
	strArr := utils.ReadLinesFromTextFile(path, fileName)
	intArr := convertStrArrToIntArr(strArr)

	return intArr
}

func convertStrArrToIntArr(strArr []string) []int {
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
