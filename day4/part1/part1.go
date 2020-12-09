package day4part1

import (
	"regexp"
	"strconv"
)

// Passport represents a passport
type Passport struct {
	PassportID     string
	CountryID      int
	IssueYear      int
	ExpirationYear int
	BirthYear      int
	HairColor      string
	EyeColor       string
	Height         string
}

func getValidPassports(passports []Passport) []Passport {
	validPassports := []Passport{}

	for _, passport := range passports {
		if passport.PassportID != "" &&
			passport.IssueYear != 0 &&
			passport.ExpirationYear != 0 &&
			passport.BirthYear != 0 &&
			passport.HairColor != "" &&
			passport.EyeColor != "" &&
			passport.Height != "" {
			validPassports = append(validPassports, passport)
		}
	}

	return validPassports
}

// ConvertStrArrToPassports converts []string to []Passport
func ConvertStrArrToPassports(strArr []string) []Passport {
	passports := []Passport{}

	for _, str := range strArr {
		passport := convertStrToPassport(str)
		passports = append(passports, passport)
	}

	return passports
}

func convertStrToPassport(str string) Passport {
	passport := Passport{}

	r, _ := regexp.Compile(`([a-z]{3}):(#*\w+)`)

	groups := r.FindAllStringSubmatch(str, -1)

	for _, group := range groups {
		fieldName := group[1]
		fieldVal := group[2]

		switch fieldName {
		case "pid":
			passport.PassportID = fieldVal
			break
		case "cid":
			passport.CountryID, _ = strconv.Atoi(fieldVal)
			break
		case "iyr":
			passport.IssueYear, _ = strconv.Atoi(fieldVal)
			break
		case "eyr":
			passport.ExpirationYear, _ = strconv.Atoi(fieldVal)
			break
		case "byr":
			passport.BirthYear, _ = strconv.Atoi(fieldVal)
			break
		case "hcl":
			passport.HairColor = fieldVal
			break
		case "ecl":
			passport.EyeColor = fieldVal
			break
		case "hgt":
			passport.Height = fieldVal
			break
		}
	}

	return passport
}
