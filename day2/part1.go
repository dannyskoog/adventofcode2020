package day2

import "strings"

func getValidPasswords(passwords []password) []password {
	var validPasswords []password

PasswordLoop:
	for _, password := range passwords {
		if !strings.Contains(password.phrase, password.policy.character) {
			continue
		}

		characterOccurrences := make(map[string]int)

		for _, char := range password.phrase {
			charStr := string(char)

			if charStr == password.policy.character {
				characterOccurrences[charStr]++
			}
		}

		for _, val := range characterOccurrences {
			if val < password.policy.occurrences.min || val > password.policy.occurrences.max {
				continue PasswordLoop
			}
		}

		validPasswords = append(validPasswords, password)
	}

	return validPasswords
}
