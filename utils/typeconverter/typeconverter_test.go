package typeconverter

import (
	"reflect"
	"testing"
)

func TestStringArrayToIntArray(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	got := StringArrayToIntArray([]string{"1", "2", "3", "4", "5"})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Int array was incorrect, got: %v, want: %v", got, want)
	}
}
