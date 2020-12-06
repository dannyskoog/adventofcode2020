package day2part2

import (
	"strconv"
	"strings"
)

type passwordPolicy struct {
	character string
	indices   map[int]bool
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

		numberOfOccurences := 0

		for i, char := range password.phrase {
			charStr := string(char)

			if password.policy.indices[i] && charStr == password.policy.character {
				numberOfOccurences++
			}
		}

		if numberOfOccurences == 1 {
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
	positionOne, _ := strconv.Atoi(occurrences[0])
	positionTwo, _ := strconv.Atoi(occurrences[1])
	character := strings.Replace(splitted[1], ":", "", 1)
	phrase := splitted[2]

	return password{
		phrase: phrase,
		policy: passwordPolicy{
			character: character,
			indices: map[int]bool{
				positionOne - 1: true,
				positionTwo - 1: true,
			},
		},
	}
}
