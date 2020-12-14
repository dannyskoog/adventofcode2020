package day7part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getBagsFromTextFile(path string, fileName string) BagMap {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	bags := ConvertStrArrToBags(strArr)

	return bags
}

func TestConvertStrToBagRule(t *testing.T) {
	tests := []struct {
		in   string
		want bag
	}{
		{
			"muted black bags contain 4 vibrant indigo bags, 2 wavy crimson bags, 4 light cyan bags, 5 dim salmon bags.",
			bag{
				name: "muted black",
				innerBags: []InnerBag{
					InnerBag{
						Name:   "vibrant indigo",
						Amount: 4,
					},
					InnerBag{
						Name:   "wavy crimson",
						Amount: 2,
					},
					InnerBag{
						Name:   "light cyan",
						Amount: 4,
					},
					InnerBag{
						Name:   "dim salmon",
						Amount: 5,
					},
				},
			},
		},
		{
			"plaid crimson bags contain 1 plaid plum bag.",
			bag{
				name: "plaid crimson",
				innerBags: []InnerBag{
					InnerBag{
						Name:   "plaid plum",
						Amount: 1,
					},
				},
			},
		},
		{
			"dim gold bags contain no other bags.",
			bag{
				name:      "dim gold",
				innerBags: []InnerBag{},
			},
		},
	}

	for _, e := range tests {
		got := convertStrToBag(e.in)

		if !reflect.DeepEqual(got, e.want) {
			t.Errorf("Bag rule was incorrect, got: %+v, want: %+v", got, e.want)
		}
	}

}

func TestConvertStrArrToBagRules(t *testing.T) {
	want := BagMap{
		"muted violet": []InnerBag{
			InnerBag{
				Name:   "pale red",
				Amount: 3,
			},
			InnerBag{
				Name:   "dull red",
				Amount: 5,
			},
			InnerBag{
				Name:   "light cyan",
				Amount: 5,
			},
		},
		"wavy aqua": []InnerBag{},
		"posh plum": []InnerBag{
			InnerBag{
				Name:   "faded green",
				Amount: 3,
			},
		},
	}
	got := ConvertStrArrToBags([]string{
		"muted violet bags contain 3 pale red bags, 5 dull red bags, 5 light cyan bags.",
		"wavy aqua bags contain no other bags.",
		"posh plum bags contain 3 faded green bags.",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Bag rules were incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestAreInnerBagsContainingBag(t *testing.T) {
	bags := BagMap{
		"bright lavender": []InnerBag{},
		"clear lime": []InnerBag{
			InnerBag{
				Name:   "dim olive",
				Amount: 4,
			},
			InnerBag{
				Name:   "pale plum",
				Amount: 4,
			},
		},
		"pale plum": []InnerBag{
			InnerBag{
				Name:   "dim olive",
				Amount: 1,
			},
		},
		"dim olive": []InnerBag{
			InnerBag{
				Name:   "bright lavender",
				Amount: 1,
			},
		},
	}

	tests := []struct {
		in struct {
			first  []InnerBag
			second string
		}
		want bool
	}{
		{
			in: struct {
				first  []InnerBag
				second string
			}{
				first:  bags["clear lime"],
				second: "bright lavender",
			},
			want: true,
		},
		{
			in: struct {
				first  []InnerBag
				second string
			}{
				first:  bags["pale plume"],
				second: "clear lime",
			},
			want: false,
		},
		{
			in: struct {
				first  []InnerBag
				second string
			}{
				first:  bags["pale plum"],
				second: "dim olive",
			},
			want: true,
		},
		{
			in: struct {
				first  []InnerBag
				second string
			}{
				first:  bags["bright lavender"],
				second: "pale plum",
			},
			want: false,
		},
	}

	for _, e := range tests {
		got := areInnerBagsContainingBag(bags, e.in.first, e.in.second)

		if got != e.want {
			t.Errorf("Inner bags containing value for (1st parameter: %v, 2nd parameter: %v) was incorrect, got: %v, want: %v", e.in.first, e.in.second, got, e.want)
		}
	}
}

func TestGetBagsByContainedBag(t *testing.T) {
	bags := BagMap{
		"bright lavender": []InnerBag{},
		"clear lime": []InnerBag{
			InnerBag{
				Name:   "bright lavender",
				Amount: 1,
			},
			InnerBag{
				Name:   "dim olive",
				Amount: 4,
			},
			InnerBag{
				Name:   "pale plum",
				Amount: 4,
			},
		},
		"pale plum": []InnerBag{
			InnerBag{
				Name:   "dim olive",
				Amount: 1,
			},
		},
		"dim olive": []InnerBag{
			InnerBag{
				Name:   "bright lavender",
				Amount: 1,
			},
		},
	}
	want := []string{"clear lime", "pale plum"}
	got := getBagNamesByContainedBag(bags, "dim olive")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Bags were incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestGetBagNamesByContainedBagCount(t *testing.T) {
	bags := getBagsFromTextFile("../", "input.txt")
	want := 172
	got := len(getBagNamesByContainedBag(bags, "shiny gold"))

	if got != want {
		t.Errorf("Count was incorrect, got: %d, want: %d", got, want)
	}
}
