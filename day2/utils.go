package day2

import (
	"adventofcode2020/utils"
	"strconv"
	"strings"
)

type passwordPolicyCharacterOccurrence struct {
	min int
	max int
}

type passwordPolicy struct {
	character   string
	occurrences passwordPolicyCharacterOccurrence
}

type password struct {
	phrase string
	policy passwordPolicy
}

func getPasswordsFromTextFile(path string, fileName string) []password {
	strArr := utils.ReadLinesFromTextFile(path, fileName)
	passwords := convertStrArrToPasswordsEntities(strArr)

	return passwords
}

func convertStrArrToPasswordsEntities(strArr []string) []password {
	var passwords []password

	for _, str := range strArr {
		password := convertStrToPasswordEntity(str)
		passwords = append(passwords, password)
	}

	return passwords
}

func convertStrToPasswordEntity(str string) password {
	splitted := strings.Split(str, " ")

	occurrences := strings.Split(splitted[0], "-")
	minOccurrences, _ := strconv.Atoi(occurrences[0])
	maxOccurrences, _ := strconv.Atoi(occurrences[1])
	character := strings.Replace(splitted[1], ":", "", 1)
	phrase := splitted[2]

	return password{
		phrase: phrase,
		policy: passwordPolicy{
			character: character,
			occurrences: passwordPolicyCharacterOccurrence{
				min: minOccurrences,
				max: maxOccurrences,
			},
		},
	}
}
