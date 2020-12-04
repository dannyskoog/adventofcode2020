package day2

import (
	"reflect"
	"testing"
)

func TestGetPasswordsFromTextFile(t *testing.T) {
	want := []password{
		password{
			phrase: "nnnnnnqnnntnnnnnnnnn",
			policy: passwordPolicy{
				character: "n",
				occurrences: passwordPolicyCharacterOccurrence{
					min: 16,
					max: 17,
				},
			},
		},
		password{
			phrase: "lfblltnlrllll",
			policy: passwordPolicy{
				character: "l",
				occurrences: passwordPolicyCharacterOccurrence{
					min: 2,
					max: 13,
				},
			},
		},
	}
	got := getPasswordsFromTextFile("./", "utils_test.txt")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Passwords were incorrect, got: %v, want %v", got, want)
	}
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

func TestConvertStrToPasswordEntity(t *testing.T) {
	want := password{
		phrase: "kjdfjjjjjjrpjfhjjrj",
		policy: passwordPolicy{
			character: "j",
			occurrences: passwordPolicyCharacterOccurrence{
				min: 8,
				max: 12,
			},
		},
	}
	got := convertStrToPasswordEntity("8-12 j: kjdfjjjjjjrpjfhjjrj")

	if got != want {
		t.Errorf("Password entity was incorrect, got: %v, want: %v", got, want)
	}
}
