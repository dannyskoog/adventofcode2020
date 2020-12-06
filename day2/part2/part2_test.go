package day2part2

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

func TestConvertStrToPasswordEntity(t *testing.T) {
	want := password{
		phrase: "xhgntttllqtpjzttltt",
		policy: passwordPolicy{
			character: "t",
			indices:   map[int]bool{4: true, 11: true},
		},
	}
	got := convertStrToPasswordEntity("5-12 t: xhgntttllqtpjzttltt")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Password entity was incorrect, got: %v, want: %v", got, want)
	}
}

func TestConvertStrArrToPasswordEntities(t *testing.T) {
	want := []password{
		password{
			phrase: "stvsnzssssssssml",
			policy: passwordPolicy{
				character: "s",
				indices:   map[int]bool{11: true, 15: true},
			},
		},
		password{
			phrase: "kkkkkvkkk",
			policy: passwordPolicy{
				character: "k",
				indices:   map[int]bool{5: true, 8: true},
			},
		},
	}

	got := convertStrArrToPasswordsEntities([]string{
		"12-16 s: stvsnzssssssssml",
		"6-9 k: kkkkkvkkk",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Password entities were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPasswords(t *testing.T) {
	want := []password{
		password{
			phrase: "jgjjljmmjjzkgjjjqx",
			policy: passwordPolicy{
				character: "j",
				indices:   map[int]bool{6: true, 13: true},
			},
		},
		password{
			phrase: "rjnvr",
			policy: passwordPolicy{
				character: "r",
				indices:   map[int]bool{3: true, 4: true},
			},
		},
	}

	got := getValidPasswords([]password{
		password{
			phrase: "jgjjljmmjjzkgjjjqx",
			policy: passwordPolicy{
				character: "j",
				indices:   map[int]bool{6: true, 13: true},
			},
		},
		password{
			phrase: "rjnvr",
			policy: passwordPolicy{
				character: "r",
				indices:   map[int]bool{3: true, 4: true},
			},
		},
		password{
			phrase: "xxtxxgrxlxzzwgxm",
			policy: passwordPolicy{
				character: "x",
				indices:   map[int]bool{11: true, 15: true},
			},
		},
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Valid passwords were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetValidPasswordsLength(t *testing.T) {
	passwords := getPasswordsFromTextFile("../", "input.txt")
	want := 699
	got := len(getValidPasswords(passwords))

	if got != want {
		t.Errorf("Length was incorrect, got: %d, want: %d", got, want)
	}
}
