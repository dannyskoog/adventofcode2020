package day1

// CalculateProductOfTwoNumbersBySum calculates the product of two numbers who's sum matches the criteria
func CalculateProductOfTwoNumbersBySum(numbers []int, sum int) int {
	foundNumbers := findTwoNumbersBySum(numbers, sum)
	product := calculateProductOfNumbers(foundNumbers)

	return product
}

func findTwoNumbersBySum(numbers []int, sum int) []int {
	var numberMap = make(map[int]bool)

	for i := 0; i < len(numbers); i++ {
		if numberMap[sum-numbers[i]] == true {
			return []int{sum - numbers[i], numbers[i]}
		}

		numberMap[numbers[i]] = true
	}

	return []int{}
}

func calculateProductOfNumbers(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		if sum == 0 {
			sum += num
		} else {
			sum *= num
		}
	}

	return sum
}
