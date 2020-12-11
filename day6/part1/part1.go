package day6part1

import (
	"strings"
)

// Questions represents a collection of unique questions
type Questions map[string]int

// Person represents a person
type Person struct {
	Questions []string
}

// Group represents a group of people
type Group struct {
	People []Person
}

func calculateSumOfUniqueQuestionsForGroups(groups []Group) int {
	sum := 0

	for _, group := range groups {
		sum += len(GetUniqueQuestionsByGroup(group))
	}

	return sum
}

// GetUniqueQuestionsByGroup returns unique questions by group
func GetUniqueQuestionsByGroup(group Group) Questions {
	uniqueQuestions := make(Questions)

	for _, person := range group.People {
		for _, question := range person.Questions {
			uniqueQuestions[question]++
		}
	}

	return uniqueQuestions
}

// ConvertStrArrToGroups converts []string to []Group
func ConvertStrArrToGroups(strArr []string) []Group {
	groups := []Group{}

	for _, str := range strArr {
		group := convertStrToGroup(str)
		groups = append(groups, group)
	}

	return groups
}

func convertStrToGroup(strGroup string) Group {
	group := Group{
		People: []Person{},
	}
	splitted := strings.Split(strGroup, "\r\n")

	for _, strPerson := range splitted {
		person := convertStrToPerson(strPerson)
		group.People = append(group.People, person)
	}

	return group
}

func convertStrToPerson(str string) Person {
	person := Person{
		Questions: []string{},
	}

	for _, char := range str {
		person.Questions = append(person.Questions, string(char))
	}

	return person
}
