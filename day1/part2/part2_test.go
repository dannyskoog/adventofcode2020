package day1part2

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

func TestFindThreeNumbersBySum(t *testing.T) {
	sum := 19
	want := []int{2, 8, 9}
	got := findThreeNumbersBySum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, sum)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Numbers by sum (%d) were incorrect, got: %d, want: %d", sum, got, want)
	}
}

func TestCalculateProductOfThreeNumbersBySum(t *testing.T) {
	numbers := getNumbersFromTextFile("../", "input.txt")
	want := 193598720
	got := CalculateProductOfThreeNumbersBySum(numbers, 2020)

	if got != want {
		t.Errorf("Product was incorrect, got: %d, want: %d", got, want)
	}
}
