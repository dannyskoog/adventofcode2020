package day3part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getTobogganMapFromTextFile(path string, fileName string) TobogganMap {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	tobogganMap := ConvertStrArrToTobogganMap(strArr)

	return tobogganMap
}

func TestConvertStrToTobogganMapRow(t *testing.T) {
	want := tobogganMapRow{
		positions: []tobogganMapRowPosition{tree, tree, tree, openSquare, tree, tree, openSquare, tree},
	}
	got := convertStrToTobogganMapRow("###.##.#")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Toboggan map row was incorrect, got: %v, want: %v", got, want)
	}
}

func TestConvertStrArrToTobogganMap(t *testing.T) {
	want := TobogganMap{
		rows: []tobogganMapRow{
			tobogganMapRow{
				positions: []tobogganMapRowPosition{tree, openSquare, tree, openSquare, openSquare, tree, openSquare, openSquare},
			},
			tobogganMapRow{
				positions: []tobogganMapRowPosition{openSquare, tree, openSquare, openSquare, openSquare, openSquare, openSquare, openSquare},
			},
		},
	}
	got := ConvertStrArrToTobogganMap([]string{"#.#..#..", ".#......"})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Toboggan map rows were incorrect, got: %v, want: %v", got, want)
	}
}

func TestCalculateNumberOfTreesEncounteredInTobogganMapBySlope(t *testing.T) {
	tobogganMap := getTobogganMapFromTextFile("../", "input.txt")

	want := 278
	got := CalculateNumberOfTreesEncounteredInTobogganMapBySlope(tobogganMap, TobogganMapSlope{RightSteps: 3, DownSteps: 1})

	if got != want {
		t.Errorf("Number of trees were incorrect, got: %d, want: %d", got, want)
	}
}
