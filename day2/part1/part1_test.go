package day2part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getPasswordsFromTextFile(path string, fileName string) []password {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	passwords := convertStrArrToPasswordsEntities(strArr)

	return passwords
}

func TestConvertStrArrToPasswordsEntities(t *testing.T) {
	want := []password{
		password{
			phrase: "rrrrrrsrrrrrrjr",
			policy: passwordPolicy{
				character: "r",
				occurrences: passwordPolicyCharacterOccurrence{
					min: 13,
					max: 14,
				},
			},
		},
		password{
			phrase: "xxxmxtxwxvxxpxx",
			policy: passwordPolicy{
				character: "x",
				occurrences: passwordPolicyCharacterOccurrence{
					min: 10,
					max: 12,
				},
			},
		},
	}
	got := convertStrArrToPasswordsEntities([]string{
		"13-14 r: rrrrrrsrrrrrrjr",
		"10-12 x: xxxmxtxwxvxxpxx",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Password entities were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPasswords(t *testing.T) {
	passwords := []password{
		password{
			phrase: "nvbedbvfwbewbkh",
			policy: passwordPolicy{
				character: "e",
				occurrences: passwordPolicyCharacterOccurrence{
					min: 1,
					max: 2,
				},
			},
		},
		password{
			phrase: "aaiiuuiippooiii",
			policy: passwordPolicy{
				character: "i",
				occurrences: passwordPolicyCharacterOccurrence{
					min: 4,
					max: 6,
				},
			},
		},
	}
	want := []password{
		password{
			phrase: "nvbedbvfwbewbkh",
			policy: passwordPolicy{
				character: "e",
				occurrences: passwordPolicyCharacterOccurrence{
					min: 1,
					max: 2,
				},
			},
		},
	}
	got := getValidPasswords(passwords)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Valid passwords were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPasswordsLength(t *testing.T) {
	passwords := getPasswordsFromTextFile("../", "input.txt")
	want := 638
	got := len(getValidPasswords(passwords))

	if got != want {
		t.Errorf("Length was incorrect, got: %d, want: %d", got, want)
	}
}
