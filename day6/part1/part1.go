package day6part1

import (
	"strings"
)

type uniqueQuestions map[string]bool

type person struct {
	questions []string
}

type group struct {
	people []person
}

func calculateSumOfUniqueQuestionsForGroups(groups []group) int {
	sum := 0

	for _, group := range groups {
		sum += len(getUniqueQuestionsByGroup(group))
	}

	return sum
}

func getUniqueQuestionsByGroup(group group) uniqueQuestions {
	uniqueQuestions := make(uniqueQuestions)

	for _, person := range group.people {
		for _, question := range person.questions {
			uniqueQuestions[question] = true
		}
	}

	return uniqueQuestions
}

func convertStrArrToGroups(strArr []string) []group {
	groups := []group{}

	for _, str := range strArr {
		group := convertStrToGroup(str)
		groups = append(groups, group)
	}

	return groups
}

func convertStrToGroup(strGroup string) group {
	group := group{
		people: []person{},
	}
	splitted := strings.Split(strGroup, "\r\n")

	for _, strPerson := range splitted {
		person := convertStrToPerson(strPerson)
		group.people = append(group.people, person)
	}

	return group
}

func convertStrToPerson(str string) person {
	person := person{
		questions: []string{},
	}

	for _, char := range str {
		person.questions = append(person.questions, string(char))
	}

	return person
}
