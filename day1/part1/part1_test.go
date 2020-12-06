package day1part1

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
	intArr := typeconverter.StringArrayToIntArray(strArr)

	return intArr
}

func TestFindTwoNumbersBySum(t *testing.T) {
	sum := 16
	want := []int{7, 9}
	got := findTwoNumbersBySum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, sum)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Numbers by sum (%d) were incorrect, got: %d, want: %d", sum, got, want)
	}
}

func TestCalculateProductOfNumbers(t *testing.T) {
	want := 120
	got := CalculateProductOfNumbers([]int{1, 2, 3, 4, 5})

	if got != want {
		t.Errorf("Product was incorrect, got: %d, want: %d", got, want)
	}
}

func TestCalulcateProductOfTwoNumbersBySum(t *testing.T) {
	numbers := getNumbersFromTextFile("../", "input.txt")
	want := 32064
	got := CalculateProductOfTwoNumbersBySum(numbers, 2020)

	if got != want {
		t.Errorf("Product was incorrect, got: %d; want %d", got, want)
	}
}
