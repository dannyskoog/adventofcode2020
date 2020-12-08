package day4part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getPassportsFromTextFile(path string, fileName string) []passport {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n\r\n")
	passports := convertStrArrToPassports(strArr)

	return passports
}

func TestConvertStrToPassport(t *testing.T) {
	want := passport{
		passportID:     "793008846",
		countryID:      224,
		issueYear:      2013,
		expirationYear: 2025,
		birthYear:      1973,
		hairColor:      "#341e13",
		eyeColor:       "grn",
		height:         "187cm",
	}
	got := convertStrToPassport("pid:793008846 eyr:2025 ecl:grn hcl:#341e13\r\nhgt:187cm\r\n\r\nbyr:1973 cid:224\r\niyr:2013")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Passport was incorrect, got: %v, want: %v", got, want)
	}
}

func TestConvertStrArrToPassports(t *testing.T) {
	want := []passport{
		passport{
			passportID:     "57104308",
			countryID:      295,
			issueYear:      1980,
			expirationYear: 2040,
			birthYear:      1974,
			hairColor:      "z",
			eyeColor:       "amb",
			height:         "192in",
		},
		passport{
			passportID:     "918371363",
			issueYear:      2012,
			expirationYear: 2029,
			birthYear:      2012,
			eyeColor:       "xry",
			height:         "65cm",
		},
	}

	got := convertStrArrToPassports([]string{
		"byr:1974\r\neyr:2040 pid:57104308 iyr:1980 hcl:z\r\nhgt:192in cid:295 ecl:amb",
		"pid:918371363\r\necl:xry\r\niyr:2012\r\nbyr:2012 hgt:65cm\r\neyr:2029",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Passports were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPassports(t *testing.T) {
	want := []passport{
		passport{
			passportID:     "#5d8fcc",
			issueYear:      2012,
			expirationYear: 1933,
			birthYear:      2003,
			hairColor:      "#7d3b0c",
			eyeColor:       "gry",
			height:         "175",
		},
	}
	got := getValidPassports([]passport{
		passport{
			passportID:     "#5d8fcc",
			issueYear:      2012,
			expirationYear: 1933,
			birthYear:      2003,
			hairColor:      "#7d3b0c",
			eyeColor:       "gry",
			height:         "175",
		},
		passport{
			passportID:     "921387245",
			countryID:      82,
			issueYear:      2015,
			expirationYear: 2023,
			hairColor:      "#c0946f",
			eyeColor:       "grn",
			height:         "190cm",
		},
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Valid passwords were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPasswordsLength(t *testing.T) {
	passports := getPassportsFromTextFile("../", "input.txt")
	want := 226
	got := len(getValidPassports(passports))

	if got != want {
		t.Errorf("Valid passwords length was incorrect, got: %d, want: %d", got, want)
	}
}
