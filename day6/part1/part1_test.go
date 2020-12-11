package day6part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getGroupsFromTextFile(path string, fileName string) []group {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n\r\n")
	groups := convertStrArrToGroups(strArr)

	return groups
}

func TestConvertStrToPerson(t *testing.T) {
	want := person{
		questions: []string{
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
	want := group{
		people: []person{
			person{
				questions: []string{
					"a",
					"b",
					"c",
					"x",
				},
			},
			person{
				questions: []string{
					"a",
					"b",
					"c",
					"y",
				},
			},
			person{
				questions: []string{
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
	want := []group{
		group{
			people: []person{
				person{
					questions: []string{
						"a",
						"b",
						"c",
					},
				},
			},
		},
		group{
			people: []person{
				person{
					questions: []string{
						"a",
					},
				},
				person{
					questions: []string{
						"b",
					},
				},
				person{
					questions: []string{
						"c",
					},
				},
			},
		},
		group{
			people: []person{
				person{
					questions: []string{
						"a",
						"b",
					},
				},
				person{
					questions: []string{
						"a",
						"c",
					},
				},
			},
		},
		group{
			people: []person{
				person{
					questions: []string{
						"a",
					},
				},
				person{
					questions: []string{
						"a",
					},
				},
				person{
					questions: []string{
						"a",
					},
				},
				person{
					questions: []string{
						"a",
					},
				},
			},
		},
		group{
			people: []person{
				person{
					questions: []string{
						"b",
					},
				},
			},
		},
	}
	got := convertStrArrToGroups([]string{
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
	want := uniqueQuestions{
		"a": true,
		"b": true,
		"c": true,
	}
	got := getUniqueQuestionsByGroup(group{
		people: []person{
			person{
				questions: []string{
					"a",
					"b",
				},
			},
			person{
				questions: []string{
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
