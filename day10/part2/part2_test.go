package day10part2

import (
	"adventofcode2020/utils/typeconverter"
	"log"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getJoltRatingsFromTextFile(path string, fileName string) []int {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	joltRatings := typeconverter.StringArrayToIntArray(strArr)

	return joltRatings
}

func TestContainsNumber(t *testing.T) {
	tests := []struct {
		in struct {
			numbers []int
			number  int
		}
		want bool
	}{
		{
			in: struct {
				numbers []int
				number  int
			}{
				numbers: []int{1, 3, 9, 11},
				number:  5,
			},
			want: false,
		},
		{
			in: struct {
				numbers []int
				number  int
			}{
				numbers: []int{7, 3, 2, 10},
				number:  2,
			},
			want: true,
		},
	}

	for _, e := range tests {
		got := containsNumber(e.in.numbers, e.in.number)

		if got != e.want {
			t.Errorf("Value was incorrect, got: %v, want: %v", got, e.want)
		}
	}
}

func TestGetNumberOfJoltRatingArrangements(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			want: 8,
		},
		{
			in:   getJoltRatingsFromTextFile("../", "input.txt"),
			want: 113387824750592,
		},
	}

	for _, e := range tests {
		got := getNumberOfJoltRatingArrangements(e.in)

		if got != e.want {
			t.Errorf("Number was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
