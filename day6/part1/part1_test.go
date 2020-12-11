package day6part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getGroupsFromTextFile(path string, fileName string) []Group {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n\r\n")
	groups := ConvertStrArrToGroups(strArr)

	return groups
}

func TestConvertStrToPerson(t *testing.T) {
	want := Person{
		Questions: []string{
			"a",
			"b",
			"c",
		},
	}

	got := convertStrToPerson("abc")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Person was incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestConvertStrToGroup(t *testing.T) {
	want := Group{
		People: []Person{
			Person{
				Questions: []string{
					"a",
					"b",
					"c",
					"x",
				},
			},
			Person{
				Questions: []string{
					"a",
					"b",
					"c",
					"y",
				},
			},
			Person{
				Questions: []string{
					"a",
					"b",
					"c",
					"z",
				},
			},
		},
	}

	got := convertStrToGroup("abcx\r\nabcy\r\nabcz")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Group was incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestConvertStrArrToGroups(t *testing.T) {
	want := []Group{
		Group{
			People: []Person{
				Person{
					Questions: []string{
						"a",
						"b",
						"c",
					},
				},
			},
		},
		Group{
			People: []Person{
				Person{
					Questions: []string{
						"a",
					},
				},
				Person{
					Questions: []string{
						"b",
					},
				},
				Person{
					Questions: []string{
						"c",
					},
				},
			},
		},
		Group{
			People: []Person{
				Person{
					Questions: []string{
						"a",
						"b",
					},
				},
				Person{
					Questions: []string{
						"a",
						"c",
					},
				},
			},
		},
		Group{
			People: []Person{
				Person{
					Questions: []string{
						"a",
					},
				},
				Person{
					Questions: []string{
						"a",
					},
				},
				Person{
					Questions: []string{
						"a",
					},
				},
				Person{
					Questions: []string{
						"a",
					},
				},
			},
		},
		Group{
			People: []Person{
				Person{
					Questions: []string{
						"b",
					},
				},
			},
		},
	}
	got := ConvertStrArrToGroups([]string{
		"abc",
		"a\r\nb\r\nc",
		"ab\r\nac",
		"a\r\na\r\na\r\na",
		"b",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Groups were incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestGetUniqueQuestionsByGroup(t *testing.T) {
	want := Questions{
		"a": 2,
		"b": 1,
		"c": 1,
	}
	got := GetUniqueQuestionsByGroup(Group{
		People: []Person{
			Person{
				Questions: []string{
					"a",
					"b",
				},
			},
			Person{
				Questions: []string{
					"a",
					"c",
				},
			},
		},
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unique questions were incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestCalculateSumOfUniqueQuestionsForGroups(t *testing.T) {
	groups := getGroupsFromTextFile("../", "input.txt")
	want := 6382
	got := calculateSumOfUniqueQuestionsForGroups(groups)

	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d", got, want)
	}
}
