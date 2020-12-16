package day10part1

import (
	"adventofcode2020/utils/typeconverter"
	"log"
	"reflect"
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

func TestGetJoltRatingDifferences(t *testing.T) {
	tests := []struct {
		in   []int
		want map[int]int
	}{
		{
			in: []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			want: map[int]int{
				1: 7,
				3: 5,
			},
		},
		{
			in: []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			want: map[int]int{
				1: 22,
				3: 10,
			},
		},
	}

	for _, e := range tests {
		got := getJoltRatingDifferences(e.in)

		if !reflect.DeepEqual(got, e.want) {
			t.Errorf("Jolt rating differences were incorrect, got: %+v, want: %+v", got, e.want)
		}
	}
}

func TestGetProductOfOneAndThreeJoltRatingDifferenceCounts(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			want: 35,
		},
		{
			in:   []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			want: 220,
		},
		{
			in:   getJoltRatingsFromTextFile("../", "input.txt"),
			want: 1917,
		},
	}

	for _, e := range tests {
		got := getProductOfOneAndThreeJoltRatingDifferenceCounts(e.in)

		if got != e.want {
			t.Errorf("Product was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
