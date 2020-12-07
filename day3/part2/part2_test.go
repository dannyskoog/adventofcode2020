package day3part2

import (
	day3part1 "adventofcode2020/day3/part1"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getTobogganMapFromTextFile(path string, fileName string) day3part1.TobogganMap {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	tobogganMap := day3part1.ConvertStrArrToTobogganMap(strArr)

	return tobogganMap
}

func TestCalculateProductOfEncounteredTreesInTobogganMapBySlopes(t *testing.T) {
	tobogganMap := getTobogganMapFromTextFile("../", "input.txt")
	tobogganMapSlopes := []day3part1.TobogganMapSlope{
		day3part1.TobogganMapSlope{
			RightSteps: 1,
			DownSteps:  1,
		},
		day3part1.TobogganMapSlope{
			RightSteps: 3,
			DownSteps:  1,
		},
		day3part1.TobogganMapSlope{
			RightSteps: 5,
			DownSteps:  1,
		},
		day3part1.TobogganMapSlope{
			RightSteps: 7,
			DownSteps:  1,
		},
		day3part1.TobogganMapSlope{
			RightSteps: 1,
			DownSteps:  2,
		},
	}

	want := 9709761600
	got := calculateProductOfEncounteredTreesInTobogganMapBySlopes(tobogganMap, tobogganMapSlopes)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Product was incorrect, got: %d, want: %d", got, want)
	}
}
