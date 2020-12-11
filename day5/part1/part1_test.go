package day5part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getBoardingPassesFromTextFile(path string, fileName string) []BoardingPass {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	boardingPasses := ConvertStrArrToBoardingPasses(strArr)

	return boardingPasses
}

func TestConvertCharToColumnDirection(t *testing.T) {
	tests := []struct {
		in   string
		want ColumnDirection
	}{
		{"L", ColumnDirectionLeft},
		{"R", ColumnDirectionRight},
	}

	for _, e := range tests {
		got := convertCharToColumnDirection(e.in)

		if got != e.want {
			t.Errorf("Column direction for %s was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestConvertCharToRowDirection(t *testing.T) {
	tests := []struct {
		in   string
		want RowDirection
	}{
		{"F", RowDirectionFront},
		{"B", RowDirectionBack},
	}

	for _, e := range tests {
		got := convertCharToRowDirection(e.in)

		if got != e.want {
			t.Errorf("Column direction for %s was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestConvertStrToBoardingPass(t *testing.T) {
	want := BoardingPass{
		RowDirections: []RowDirection{
			RowDirectionFront,
			RowDirectionBack,
			RowDirectionFront,
			RowDirectionBack,
			RowDirectionBack,
			RowDirectionFront,
			RowDirectionFront,
		},
		ColumnDirections: []ColumnDirection{
			ColumnDirectionRight,
			ColumnDirectionLeft,
			ColumnDirectionRight,
		},
	}

	got := convertStrToBoardingPass("FBFBBFFRLR")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Boarding pass was incorrect, got: %v, want: %v", got, want)
	}
}

func TestConvertStrArrToBoardingPasses(t *testing.T) {
	want := []BoardingPass{
		BoardingPass{
			RowDirections: []RowDirection{
				RowDirectionFront,
				RowDirectionFront,
				RowDirectionBack,
				RowDirectionBack,
				RowDirectionFront,
				RowDirectionBack,
				RowDirectionBack,
			},
			ColumnDirections: []ColumnDirection{
				ColumnDirectionRight,
				ColumnDirectionRight,
				ColumnDirectionRight,
			},
		},
		BoardingPass{
			RowDirections: []RowDirection{
				RowDirectionBack,
				RowDirectionFront,
				RowDirectionFront,
				RowDirectionBack,
				RowDirectionBack,
				RowDirectionBack,
				RowDirectionBack,
			},
			ColumnDirections: []ColumnDirection{
				ColumnDirectionLeft,
				ColumnDirectionLeft,
				ColumnDirectionLeft,
			},
		},
		BoardingPass{
			RowDirections: []RowDirection{
				RowDirectionFront,
				RowDirectionBack,
				RowDirectionBack,
				RowDirectionFront,
				RowDirectionBack,
				RowDirectionFront,
				RowDirectionBack,
			},
			ColumnDirections: []ColumnDirection{
				ColumnDirectionRight,
				ColumnDirectionRight,
				ColumnDirectionLeft,
			},
		},
	}
	got := ConvertStrArrToBoardingPasses([]string{
		"FFBBFBBRRR",
		"BFFBBBBLLL",
		"FBBFBFBRRL",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Boarding passes were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetColumnNumberFromBoardingPass(t *testing.T) {
	want := 5
	got := GetColumnNumberFromBoardingPass(BoardingPass{
		ColumnDirections: []ColumnDirection{
			ColumnDirectionRight,
			ColumnDirectionLeft,
			ColumnDirectionRight,
		},
	})

	if got != want {
		t.Errorf("Column number was incorrect, got: %d, want: %d", got, want)
	}
}

func TestGetRowNumberFromBoardingPass(t *testing.T) {
	want := 44
	got := GetRowNumberFromBoardingPass(BoardingPass{
		RowDirections: []RowDirection{
			RowDirectionFront,
			RowDirectionBack,
			RowDirectionFront,
			RowDirectionBack,
			RowDirectionBack,
			RowDirectionFront,
			RowDirectionFront,
		},
	})

	if got != want {
		t.Errorf("Row number was incorrect, got: %d, want: %d", got, want)
	}
}

func TestGetSeatIDFromBoardingPass(t *testing.T) {
	want := 567
	got := GetSeatIDFromRowAndColumnNumbers(70, 7)

	if got != want {
		t.Errorf("Seat ID was incorrect, got: %d, want: %d", got, want)
	}
}

func TestGetHighestSeatIDFromBoardingPasses(t *testing.T) {
	boardingPasses := getBoardingPassesFromTextFile("../", "input.txt")
	want := 861
	got := getHighestSeatIDFromBoardingPasses(boardingPasses, 128, 8)

	if got != want {
		t.Errorf("Highest seat ID was incorrect, got: %d, want: %d", got, want)
	}
}
