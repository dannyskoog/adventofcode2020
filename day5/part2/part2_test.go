package day5part2

import (
	day5part1 "adventofcode2020/day5/part1"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getBoardingPassesFromTextFile(path string, fileName string) []day5part1.BoardingPass {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	boardingPasses := day5part1.ConvertStrArrToBoardingPasses(strArr)

	return boardingPasses
}

func TestGetSeatIDsFromBoardingPasses(t *testing.T) {
	want := seatIDs{
		567: true,
		119: true,
		820: true,
	}

	got := getSeatIDsFromBoardingPasses([]day5part1.BoardingPass{
		day5part1.BoardingPass{
			RowDirections: []day5part1.RowDirection{
				day5part1.RowDirectionBack,
				day5part1.RowDirectionFront,
				day5part1.RowDirectionFront,
				day5part1.RowDirectionFront,
				day5part1.RowDirectionBack,
				day5part1.RowDirectionBack,
				day5part1.RowDirectionFront,
			},
			ColumnDirections: []day5part1.ColumnDirection{
				day5part1.ColumnDirectionRight,
				day5part1.ColumnDirectionRight,
				day5part1.ColumnDirectionRight,
			},
		},
		day5part1.BoardingPass{
			RowDirections: []day5part1.RowDirection{
				day5part1.RowDirectionFront,
				day5part1.RowDirectionFront,
				day5part1.RowDirectionFront,
				day5part1.RowDirectionBack,
				day5part1.RowDirectionBack,
				day5part1.RowDirectionBack,
				day5part1.RowDirectionFront,
			},
			ColumnDirections: []day5part1.ColumnDirection{
				day5part1.ColumnDirectionRight,
				day5part1.ColumnDirectionRight,
				day5part1.ColumnDirectionRight,
			},
		},
		day5part1.BoardingPass{
			RowDirections: []day5part1.RowDirection{
				day5part1.RowDirectionBack,
				day5part1.RowDirectionBack,
				day5part1.RowDirectionFront,
				day5part1.RowDirectionFront,
				day5part1.RowDirectionBack,
				day5part1.RowDirectionBack,
				day5part1.RowDirectionFront,
			},
			ColumnDirections: []day5part1.ColumnDirection{
				day5part1.ColumnDirectionRight,
				day5part1.ColumnDirectionLeft,
				day5part1.ColumnDirectionLeft,
			},
		},
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Seat IDs were incorrect, got: %v, want: %v", got, want)
	}
}

func TestGetMySeatIDBasedOnOtherBoardingPasses(t *testing.T) {
	boardingPasses := getBoardingPassesFromTextFile("../", "input.txt")
	want := 633
	got := getMySeatIDBasedOnOtherBoardingPasses(boardingPasses)

	if got != want {
		t.Errorf("My seat ID was incorrect, got: %d, want: %d", got, want)
	}
}
