package day9part2

import (
	"adventofcode2020/utils/typeconverter"
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
	numbers := typeconverter.StringArrayToIntArray(strArr)

	return numbers
}

func TestFindContiguousNumbersBySum(t *testing.T) {
	want := []int{15, 25, 47, 40}
	got := findContiguousNumbersBySum([]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 127)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Contiguous numbers were incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestGetMinAndMaxNumbers(t *testing.T) {
	want := struct {
		min int
		max int
	}{
		min: 15,
		max: 47,
	}
	gotMin, gotMax := getMinAndMaxNumbers([]int{15, 25, 47, 40})

	if gotMin != want.min {
		t.Errorf("Min number was incorrect, got: %d, want: %d", gotMin, want.min)
	}

	if gotMax != want.max {
		t.Errorf("Max number was incorrect, got: %d, want: %d", gotMax, want.max)
	}
}

func TestGetEncryptionWeakness(t *testing.T) {
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
			want: 62,
		},
		{
			in: struct {
				numbers  []int
				preamble int
			}{
				numbers:  getNumbersFromTextFile("../", "input.txt"),
				preamble: 25,
			},
			want: 3340942,
		},
	}

	for _, e := range tests {
		got := getEncryptionWeakness(e.in.numbers, e.in.preamble)

		if got != e.want {
			t.Errorf("Encryption weakness was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
