package day2

import (
	"reflect"
	"testing"
)

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
	passwords := getPasswordsFromTextFile("./", "input.txt")
	want := 638
	got := len(getValidPasswords(passwords))

	if got != want {
		t.Errorf("Length was incorrect, got: %d, want: %d", got, want)
	}
}
