package day7part2

import (
	day7part1 "adventofcode2020/day7/part1"
	"log"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getBagsFromTextFile(path string, fileName string) day7part1.BagMap {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	bags := day7part1.ConvertStrArrToBags(strArr)

	return bags
}

func TestGetInnerBagsCount(t *testing.T) {
	bags := day7part1.BagMap{
		"bright lavender": []day7part1.InnerBag{},
		"clear lime": []day7part1.InnerBag{
			day7part1.InnerBag{
				Name:   "bright lavender",
				Amount: 1,
			},
			day7part1.InnerBag{
				Name:   "dim olive",
				Amount: 4,
			},
			day7part1.InnerBag{
				Name:   "pale plum",
				Amount: 4,
			},
		},
		"pale plum": []day7part1.InnerBag{
			day7part1.InnerBag{
				Name:   "dim olive",
				Amount: 1,
			},
		},
		"dim olive": []day7part1.InnerBag{
			day7part1.InnerBag{
				Name:   "bright lavender",
				Amount: 1,
			},
		},
	}
	tests := []struct {
		in   string
		want int
	}{
		{
			in:   "clear lime",
			want: 21,
		},
		{
			in:   "bright lavender",
			want: 0,
		},
		{
			in:   "dim olive",
			want: 1,
		},
	}

	for _, e := range tests {
		got := getInnerBagsCount(bags, bags[e.in])

		if got != e.want {
			t.Errorf("Inner bags count for %s was incorrect, got: %d, want: %d", e.in, got, e.want)
		}
	}
}

func TestGetInnerBagsCountByContainingBag(t *testing.T) {
	bags := getBagsFromTextFile("../", "input.txt")
	want := 39645
	got := getInnerBagsCountByContainingBag(bags, "shiny gold")

	if got != want {
		t.Errorf("Inner bags count by containing bag was incorrect, got: %d, want: %d", got, want)
	}
}
