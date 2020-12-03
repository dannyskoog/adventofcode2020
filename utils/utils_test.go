package utils

import (
	"testing"
)

func TestReadTextFile(t *testing.T) {
	want := "1\r\n2\r\n3\r\n4\r\n5"
	got := readTextFile("./", "utils_test.txt")

	if want != got {
		t.Errorf("String was incorrect, got: %s, want: %s", got, want)
	}
}
