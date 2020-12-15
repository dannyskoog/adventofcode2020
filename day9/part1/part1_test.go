package day9part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getNumbersFromTextFile(path string, fileName string) []int {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	numbers := convertStrArrToNumbers(strArr)

	return numbers
}

func TestConvertStrArrToNumbers(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := convertStrArrToNumbers([]string{"1", "2", "3", "4", "5"})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Numbers were incorrect, got: %v, want: %v", got, want)
	}
}

func TestIsNumberSumOfTwoOtherNumbers(t *testing.T) {
	tests := []struct {
		in struct {
			first  []int
			second int
		}
		want bool
	}{
		{
			in: struct {
				first  []int
				second int
			}{
				first:  []int{35, 20, 15, 25, 47},
				second: 40,
			},
			want: true,
		},
		{
			in: struct {
				first  []int
				second int
			}{
				first:  []int{95, 102, 117, 150, 182},
				second: 127,
			},
			want: false,
		},
	}

	for _, e := range tests {
		got := isNumberSumOfTwoOtherNumbers(e.in.first, e.in.second)

		if got != e.want {
			t.Errorf("Value was incorrect, got: %v, want: %v", got, e.want)
		}
	}
}

func TestGetFirstInvalidNumber(t *testing.T) {
	tests := []struct {
		in struct {
			first  []int
			second int
		}
		want int
	}{
		{
			in: struct {
				first  []int
				second int
			}{
				first:  []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576},
				second: 5,
			},
			want: 127,
		},
		{
			in: struct {
				first  []int
				second int
			}{
				first:  getNumbersFromTextFile("../", "input.txt"),
				second: 25,
			},
			want: 25918798,
		},
	}

	for _, e := range tests {
		got := getFirstInvalidNumber(e.in.first, e.in.second)

		if got != e.want {
			t.Errorf("First invalid number was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
