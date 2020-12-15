package day9part1

import (
	"strconv"
)

func getFirstInvalidNumber(numbers []int, preambleCount int) int {
	for i := range numbers {
		previousNumbers, numberToCheck := numbers[i:i+preambleCount], numbers[i+preambleCount]
		isNumberValid := isNumberSumOfTwoOtherNumbers(previousNumbers, numberToCheck)

		if !isNumberValid {
			return numberToCheck
		}
	}

	panic("No invalid number was found")
}

func isNumberSumOfTwoOtherNumbers(numbers []int, number int) bool {
	numbersLength := len(numbers)

	for i := 0; i < numbersLength-1; i++ {
		for j := i + 1; j < numbersLength; j++ {
			if numbers[i]+numbers[j] == number {
				return true
			}
		}
	}

	return false
}

func convertStrArrToNumbers(strArr []string) []int {
	numbers := []int{}

	for _, numberStr := range strArr {
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
	}

	return numbers
}
