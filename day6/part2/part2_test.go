package day6part2

import (
	day6part1 "adventofcode2020/day6/part1"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getGroupsFromTextFile(path string, fileName string) []day6part1.Group {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n\r\n")
	groups := day6part1.ConvertStrArrToGroups(strArr)

	return groups
}

func TestGetUnanimousQuestionsByGroup(t *testing.T) {
	want := []string{"a"}
	got := getUnanimousQuestionsByGroup(day6part1.Group{
		People: []day6part1.Person{
			day6part1.Person{
				Questions: []string{
					"a",
					"b",
				},
			},
			day6part1.Person{
				Questions: []string{
					"a",
					"c",
				},
			},
		},
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unanimous questions were incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestGetSumOfUnanimousQuestionsByGroups(t *testing.T) {
	groups := getGroupsFromTextFile("../", "input.txt")
	want := 3197
	got := getSumOfUnanimousQuestionsByGroups(groups)

	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d", got, want)
	}
}
