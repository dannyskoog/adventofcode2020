package day5part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getBoardingPassesFromTextFile(path string, fileName string) []boardingPass {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	boardingPasses := convertStrArrToBoardingPasses(strArr)

	return boardingPasses
}

func TestConvertCharToColumnDirection(t *testing.T) {
	tests := []struct {
		in   string
		want columnDirection
	}{
		{"L", columnDirectionLeft},
		{"R", columnDirectionRight},
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
		want rowDirection
	}{
		{"F", rowDirectionFront},
		{"B", rowDirectionBack},
	}

	for _, e := range tests {
		got := convertCharToRowDirection(e.in)

		if got != e.want {
			t.Errorf("Column direction for %s was incorrect, got: %v, want: %v", e.in, got, e.want)
		}
	}
}

func TestConvertStrToBoardingPass(t *testing.T) {
	want := boardingPass{
		rowDirections: []rowDirection{
			rowDirectionFront,
			rowDirectionBack,
			rowDirectionFront,
			rowDirectionBack,
			rowDirectionBack,
			rowDirectionFront,
			rowDirectionFront,
		},
		columnDirections: []columnDirection{
			columnDirectionRight,
			columnDirectionLeft,
			columnDirectionRight,
		},
	}

	got := convertStrToBoardingPass("FBFBBFFRLR")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Boarding pass was incorrect, got: %v, want: %v", got, want)
	}
}

func TestConvertStrArrToBoardingPasses(t *testing.T) {
	want := []boardingPass{
		boardingPass{
			rowDirections: []rowDirection{
				rowDirectionFront,
				rowDirectionFront,
				rowDirectionBack,
				rowDirectionBack,
				rowDirectionFront,
				rowDirectionBack,
				rowDirectionBack,
			},
			columnDirections: []columnDirection{
				columnDirectionRight,
				columnDirectionRight,
				columnDirectionRight,
			},
		},
		boardingPass{
			rowDirections: []rowDirection{
				rowDirectionBack,
				rowDirectionFront,
				rowDirectionFront,
				rowDirectionBack,
				rowDirectionBack,
				rowDirectionBack,
				rowDirectionBack,
			},
			columnDirections: []columnDirection{
				columnDirectionLeft,
				columnDirectionLeft,
				columnDirectionLeft,
			},
		},
		boardingPass{
			rowDirections: []rowDirection{
				rowDirectionFront,
				rowDirectionBack,
				rowDirectionBack,
				rowDirectionFront,
				rowDirectionBack,
				rowDirectionFront,
				rowDirectionBack,
			},
			columnDirections: []columnDirection{
				columnDirectionRight,
				columnDirectionRight,
				columnDirectionLeft,
			},
		},
	}
	got := convertStrArrToBoardingPasses([]string{
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
	got := getColumnNumberFromBoardingPass(boardingPass{
		columnDirections: []columnDirection{
			columnDirectionRight,
			columnDirectionLeft,
			columnDirectionRight,
		},
	})

	if got != want {
		t.Errorf("Column number was incorrect, got: %d, want: %d", got, want)
	}
}

func TestGetRowNumberFromBoardingPass(t *testing.T) {
	want := 44
	got := getRowNumberFromBoardingPass(boardingPass{
		rowDirections: []rowDirection{
			rowDirectionFront,
			rowDirectionBack,
			rowDirectionFront,
			rowDirectionBack,
			rowDirectionBack,
			rowDirectionFront,
			rowDirectionFront,
		},
	})

	if got != want {
		t.Errorf("Row number was incorrect, got: %d, want: %d", got, want)
	}
}

func TestGetSeatIDFromBoardingPass(t *testing.T) {
	want := 820
	got := getSeatIDFromBoardingPass(boardingPass{
		rowDirections: []rowDirection{
			rowDirectionBack,
			rowDirectionBack,
			rowDirectionFront,
			rowDirectionFront,
			rowDirectionBack,
			rowDirectionBack,
			rowDirectionFront,
		},
		columnDirections: []columnDirection{
			columnDirectionRight,
			columnDirectionLeft,
			columnDirectionLeft,
		},
	})

	if got != want {
		t.Errorf("Seat ID was incorrect, got: %d, want: %d", got, want)
	}
}

func TestGetHighestSeatIDFromBoardingPasses(t *testing.T) {
	boardingPasses := getBoardingPassesFromTextFile("../", "input.txt")
	want := 861
	got := getHighestSeatIDFromBoardingPasses(boardingPasses)

	if got != want {
		t.Errorf("Highest seat ID was incorrect, got: %d, want: %d", got, want)
	}
}
