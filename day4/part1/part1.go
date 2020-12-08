package day4part1

import (
	"regexp"
	"strconv"
)

type passport struct {
	passportID     string
	countryID      int
	issueYear      int
	expirationYear int
	birthYear      int
	hairColor      string
	eyeColor       string
	height         string
}

func getValidPassports(passports []passport) []passport {
	validPassports := []passport{}

	for _, passport := range passports {
		if passport.passportID != "" &&
			passport.issueYear != 0 &&
			passport.expirationYear != 0 &&
			passport.birthYear != 0 &&
			passport.hairColor != "" &&
			passport.eyeColor != "" &&
			passport.height != "" {
			validPassports = append(validPassports, passport)
		}
	}

	return validPassports
}

func convertStrArrToPassports(strArr []string) []passport {
	passports := []passport{}

	for _, str := range strArr {
		passport := convertStrToPassport(str)
		passports = append(passports, passport)
	}

	return passports
}

func convertStrToPassport(str string) passport {
	passport := passport{}

	r, _ := regexp.Compile(`([a-z]{3}):(#*\w+)`)

	groups := r.FindAllStringSubmatch(str, -1)

	for _, group := range groups {
		fieldName := group[1]
		fieldVal := group[2]

		switch fieldName {
		case "pid":
			passport.passportID = fieldVal
			break
		case "cid":
			passport.countryID, _ = strconv.Atoi(fieldVal)
			break
		case "iyr":
			passport.issueYear, _ = strconv.Atoi(fieldVal)
			break
		case "eyr":
			passport.expirationYear, _ = strconv.Atoi(fieldVal)
			break
		case "byr":
			passport.birthYear, _ = strconv.Atoi(fieldVal)
			break
		case "hcl":
			passport.hairColor = fieldVal
			break
		case "ecl":
			passport.eyeColor = fieldVal
			break
		case "hgt":
			passport.height = fieldVal
			break
		}
	}

	return passport
}
