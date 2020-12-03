package day1

import (
	"reflect"
	"testing"
)

func TestConvertStrArrToIntArr(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := convertStrArrToIntArr([]string{"1", "2", "3", "4", "5"})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Array was incorrect, got: %d, want: %d", got, want)
	}
}

func TestGetNumbersFromTextFile(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := getNumbersFromTextFile("./", "utils_test.txt")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Numbers were incorrect, got: %d, want: %d", got, want)
	}
}
