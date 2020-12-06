package day2part1

import (
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

func getValidPasswords(passwords []password) []password {
	var validPasswords []password

	for _, password := range passwords {
		if !strings.Contains(password.phrase, password.policy.character) {
			continue
		}

		characterOccurrences := 0

		for _, char := range password.phrase {
			charStr := string(char)

			if charStr == password.policy.character {
				characterOccurrences++
			}
		}

		if characterOccurrences >= password.policy.occurrences.min && characterOccurrences <= password.policy.occurrences.max {
			validPasswords = append(validPasswords, password)
		}
	}

	return validPasswords
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
