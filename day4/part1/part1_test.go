package day4part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getPassportsFromTextFile(path string, fileName string) []Passport {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n\r\n")
	passports := ConvertStrArrToPassports(strArr)

	return passports
}

func TestConvertStrToPassport(t *testing.T) {
	want := Passport{
		PassportID:     "793008846",
		CountryID:      224,
		IssueYear:      2013,
		ExpirationYear: 2025,
		BirthYear:      1973,
		HairColor:      "#341e13",
		EyeColor:       "grn",
		Height:         "187cm",
	}
	got := convertStrToPassport("pid:793008846 eyr:2025 ecl:grn hcl:#341e13\r\nhgt:187cm\r\n\r\nbyr:1973 cid:224\r\niyr:2013")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Passport was incorrect, got: %v, want: %v", got, want)
	}
}

func TestConvertStrArrToPassports(t *testing.T) {
	want := []Passport{
		Passport{
			PassportID:     "57104308",
			CountryID:      295,
			IssueYear:      1980,
			ExpirationYear: 2040,
			BirthYear:      1974,
			HairColor:      "z",
			EyeColor:       "amb",
			Height:         "192in",
		},
		Passport{
			PassportID:     "918371363",
			IssueYear:      2012,
			ExpirationYear: 2029,
			BirthYear:      2012,
			EyeColor:       "xry",
			Height:         "65cm",
		},
	}

	got := ConvertStrArrToPassports([]string{
		"byr:1974\r\neyr:2040 pid:57104308 iyr:1980 hcl:z\r\nhgt:192in cid:295 ecl:amb",
		"pid:918371363\r\necl:xry\r\niyr:2012\r\nbyr:2012 hgt:65cm\r\neyr:2029",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Passports were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPassports(t *testing.T) {
	want := []Passport{
		Passport{
			PassportID:     "#5d8fcc",
			IssueYear:      2012,
			ExpirationYear: 1933,
			BirthYear:      2003,
			HairColor:      "#7d3b0c",
			EyeColor:       "gry",
			Height:         "175",
		},
	}
	got := getValidPassports([]Passport{
		Passport{
			PassportID:     "#5d8fcc",
			IssueYear:      2012,
			ExpirationYear: 1933,
			BirthYear:      2003,
			HairColor:      "#7d3b0c",
			EyeColor:       "gry",
			Height:         "175",
		},
		Passport{
			PassportID:     "921387245",
			CountryID:      82,
			IssueYear:      2015,
			ExpirationYear: 2023,
			HairColor:      "#c0946f",
			EyeColor:       "grn",
			Height:         "190cm",
		},
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Valid passports were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPassportsLength(t *testing.T) {
	passports := getPassportsFromTextFile("../", "input.txt")
	want := 226
	got := len(getValidPassports(passports))

	if got != want {
		t.Errorf("Valid passports length was incorrect, got: %d, want: %d", got, want)
	}
}
