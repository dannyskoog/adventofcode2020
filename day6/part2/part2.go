package day6part2

import (
	day6part1 "adventofcode2020/day6/part1"
)

func getSumOfUnanimousQuestionsByGroups(groups []day6part1.Group) int {
	sum := 0

	for _, group := range groups {
		sum += len(getUnanimousQuestionsByGroup(group))
	}

	return sum
}

func getUnanimousQuestionsByGroup(group day6part1.Group) []string {
	allQuestions := day6part1.GetUniqueQuestionsByGroup(group)
	numberOfPeople := len(group.People)
	unanimousQuestions := []string{}

	for question, count := range allQuestions {
		if count == numberOfPeople {
			unanimousQuestions = append(unanimousQuestions, question)
		}
	}

	return unanimousQuestions
}
