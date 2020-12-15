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
	numbers := ConvertStrArrToNumbers(strArr)

	return numbers
}

func TestConvertStrArrToNumbers(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := ConvertStrArrToNumbers([]string{"1", "2", "3", "4", "5"})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Numbers were incorrect, got: %v, want: %v", got, want)
	}
}

func TestIsNumberSumOfTwoOtherNumbers(t *testing.T) {
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
				numbers: []int{35, 20, 15, 25, 47},
				number:  40,
			},
			want: true,
		},
		{
			in: struct {
				numbers []int
				number  int
			}{
				numbers: []int{95, 102, 117, 150, 182},
				number:  127,
			},
			want: false,
		},
	}

	for _, e := range tests {
		got := isNumberSumOfTwoOtherNumbers(e.in.numbers, e.in.number)

		if got != e.want {
			t.Errorf("Value was incorrect, got: %v, want: %v", got, e.want)
		}
	}
}

func TestGetFirstInvalidNumber(t *testing.T) {
	tests := []struct {
		in struct {
			numbers  []int
			preamble int
		}
		want int
	}{
		{
			in: struct {
				numbers  []int
				preamble int
			}{
				numbers:  []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576},
				preamble: 5,
			},
			want: 127,
		},
		{
			in: struct {
				numbers  []int
				preamble int
			}{
				numbers:  getNumbersFromTextFile("../", "input.txt"),
				preamble: 25,
			},
			want: 25918798,
		},
	}

	for _, e := range tests {
		got := GetFirstInvalidNumber(e.in.numbers, e.in.preamble)

		if got != e.want {
			t.Errorf("First invalid number was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
