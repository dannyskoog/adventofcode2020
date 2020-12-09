package day4part2

import (
	day4part1 "adventofcode2020/day4/part1"
	"regexp"
)

func getValidPassports(passports []day4part1.Passport) []day4part1.Passport {
	validPassports := []day4part1.Passport{}

	for _, passport := range passports {
		if isPassportIDValid(passport.PassportID) &&
			isIssueYearValid(passport.IssueYear) &&
			isExpirationYearValid(passport.ExpirationYear) &&
			isBirthYearValid(passport.BirthYear) &&
			isHairColorValid(passport.HairColor) &&
			isEyeColorValid(passport.EyeColor) &&
			isHeightValid(passport.Height) {
			validPassports = append(validPassports, passport)
		}
	}

	return validPassports
}

func isPassportIDValid(passportID string) bool {
	matched, _ := regexp.MatchString(`^\d{9}$`, passportID)

	return matched
}

func isIssueYearValid(issueYear int) bool {
	return issueYear >= 2010 && issueYear <= 2020
}

func isExpirationYearValid(expirationYear int) bool {
	return expirationYear >= 2020 && expirationYear <= 2030
}

func isBirthYearValid(birthYear int) bool {
	return birthYear >= 1920 && birthYear <= 2002
}

func isHairColorValid(hairColor string) bool {
	matched, _ := regexp.MatchString("^#[0-9a-f]{6}$", hairColor)

	return matched
}

func isEyeColorValid(eyeColor string) bool {
	matched, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", eyeColor)

	return matched
}

func isHeightValid(height string) bool {
	matched, _ := regexp.MatchString("^((1[5-8][0-9]|19[0-3])cm)|((59|6[0-9]|7[0-6])in)$", height)

	return matched
}
