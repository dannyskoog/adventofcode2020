package day4part2

import (
	day4part1 "adventofcode2020/day4/part1"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getPassportsFromTextFile(path string, fileName string) []day4part1.Passport {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n\r\n")
	passports := day4part1.ConvertStrArrToPassports(strArr)

	return passports
}

func TestIsHeightValid(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"", false}, {"cm", false}, {"in", false}, {"190", false}, {"55cm", false}, {"149cm", false}, {"194cm", false}, {"7in", false}, {"58in", false}, {"77in", false}, {"190in", false},
		{"150cm", true}, {"190cm", true}, {"193cm", true}, {"59in", true}, {"60in", true}, {"76in", true},
	}

	for _, e := range tests {
		got := isHeightValid(e.in)

		if got != e.want {
			t.Errorf("Valid value for %s was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestIsEyeColorValid(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"", false}, {"wat", false}, {"famb", false}, {"gryl", false},
		{"amb", true}, {"blu", true}, {"brn", true}, {"gry", true}, {"grn", true}, {"hzl", true}, {"oth", true},
	}

	for _, e := range tests {
		got := isEyeColorValid(e.in)

		if got != e.want {
			t.Errorf("Valid value for %s was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestIsHairColorValid(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"", false}, {"#123abz", false}, {"#123ab", false}, {"123abc", false}, {"##123abc", false}, {"#123abcc", false},
		{"#123abc", true}, {"#abc123", true}, {"#aaaaaa", true}, {"#111111", true},
	}

	for _, e := range tests {
		got := isHairColorValid(e.in)

		if got != e.want {
			t.Errorf("Valid value for %s was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestIsBirthYearValid(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{0, false}, {10, false}, {100, false}, {1919, false}, {2003, false}, {99999, false},
		{1920, true}, {1975, true}, {2002, true},
	}

	for _, e := range tests {
		got := isBirthYearValid(e.in)

		if got != e.want {
			t.Errorf("Valid value for %d was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestIsExpirationYearValid(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{0, false}, {10, false}, {100, false}, {2019, false}, {2031, false}, {99999, false},
		{2020, true}, {2025, true}, {2030, true},
	}

	for _, e := range tests {
		got := isExpirationYearValid(e.in)

		if got != e.want {
			t.Errorf("Valid value for %d was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestIsIssueYearValid(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{0, false}, {10, false}, {100, false}, {2009, false}, {2021, false}, {99999, false},
		{2010, true}, {2015, true}, {2020, true},
	}

	for _, e := range tests {
		got := isIssueYearValid(e.in)

		if got != e.want {
			t.Errorf("Valid value for %d was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestIsPassportIDValid(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"", false}, {"10", false}, {"0123456789", false}, {"a12345678", false}, {"1234 6789", false},
		{"000000001", true}, {"123456789", true},
	}

	for _, e := range tests {
		got := isPassportIDValid(e.in)

		if got != e.want {
			t.Errorf("Valid value for %s was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestGetValidPassports(t *testing.T) {
	want := []day4part1.Passport{
		day4part1.Passport{
			PassportID:     "896056539",
			CountryID:      129,
			IssueYear:      2014,
			ExpirationYear: 2029,
			BirthYear:      1989,
			HairColor:      "#a97842",
			EyeColor:       "blu",
			Height:         "165cm",
		},
	}
	got := getValidPassports([]day4part1.Passport{
		day4part1.Passport{
			PassportID:     "896056539",
			CountryID:      129,
			IssueYear:      2014,
			ExpirationYear: 2029,
			BirthYear:      1989,
			HairColor:      "#a97842",
			EyeColor:       "blu",
			Height:         "165cm",
		},
		day4part1.Passport{
			PassportID:     "021572410",
			CountryID:      277,
			IssueYear:      2012,
			ExpirationYear: 2020,
			BirthYear:      1992,
			HairColor:      "dab227",
			EyeColor:       "brn",
			Height:         "182cm",
		},
		day4part1.Passport{
			PassportID:     "3556412378",
			IssueYear:      2023,
			ExpirationYear: 2038,
			BirthYear:      2007,
			HairColor:      "74454a",
			EyeColor:       "zzz",
			Height:         "59cm",
		},
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Valid passports were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPassportsLength(t *testing.T) {
	passports := getPassportsFromTextFile("../", "input.txt")
	want := 160
	got := len(getValidPassports(passports))

	if got != want {
		t.Errorf("Valid passports length was incorrect, got: %d, want: %d", got, want)
	}
}
