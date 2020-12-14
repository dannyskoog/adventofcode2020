package day7part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getBagsFromTextFile(path string, fileName string) bagMap {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	bags := convertStrArrToBags(strArr)

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
				innerBags: []innerBag{
					innerBag{
						name:   "vibrant indigo",
						amount: 4,
					},
					innerBag{
						name:   "wavy crimson",
						amount: 2,
					},
					innerBag{
						name:   "light cyan",
						amount: 4,
					},
					innerBag{
						name:   "dim salmon",
						amount: 5,
					},
				},
			},
		},
		{
			"plaid crimson bags contain 1 plaid plum bag.",
			bag{
				name: "plaid crimson",
				innerBags: []innerBag{
					innerBag{
						name:   "plaid plum",
						amount: 1,
					},
				},
			},
		},
		{
			"dim gold bags contain no other bags.",
			bag{
				name:      "dim gold",
				innerBags: []innerBag{},
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
	want := bagMap{
		"muted violet": []innerBag{
			innerBag{
				name:   "pale red",
				amount: 3,
			},
			innerBag{
				name:   "dull red",
				amount: 5,
			},
			innerBag{
				name:   "light cyan",
				amount: 5,
			},
		},
		"wavy aqua": []innerBag{},
		"posh plum": []innerBag{
			innerBag{
				name:   "faded green",
				amount: 3,
			},
		},
	}
	got := convertStrArrToBags([]string{
		"muted violet bags contain 3 pale red bags, 5 dull red bags, 5 light cyan bags.",
		"wavy aqua bags contain no other bags.",
		"posh plum bags contain 3 faded green bags.",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Bag rules were incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestAreInnerBagsContainingBag(t *testing.T) {
	bags := bagMap{
		"bright lavender": []innerBag{},
		"clear lime": []innerBag{
			innerBag{
				name:   "dim olive",
				amount: 4,
			},
			innerBag{
				name:   "pale plum",
				amount: 4,
			},
		},
		"pale plum": []innerBag{
			innerBag{
				name:   "dim olive",
				amount: 1,
			},
		},
		"dim olive": []innerBag{
			innerBag{
				name:   "bright lavender",
				amount: 1,
			},
		},
	}

	tests := []struct {
		in struct {
			first  []innerBag
			second string
		}
		want bool
	}{
		{
			in: struct {
				first  []innerBag
				second string
			}{
				first:  bags["clear lime"],
				second: "bright lavender",
			},
			want: true,
		},
		{
			in: struct {
				first  []innerBag
				second string
			}{
				first:  bags["pale plume"],
				second: "clear lime",
			},
			want: false,
		},
		{
			in: struct {
				first  []innerBag
				second string
			}{
				first:  bags["pale plum"],
				second: "dim olive",
			},
			want: true,
		},
		{
			in: struct {
				first  []innerBag
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
	bags := bagMap{
		"bright lavender": []innerBag{},
		"clear lime": []innerBag{
			innerBag{
				name:   "bright lavender",
				amount: 1,
			},
			innerBag{
				name:   "dim olive",
				amount: 4,
			},
			innerBag{
				name:   "pale plum",
				amount: 4,
			},
		},
		"pale plum": []innerBag{
			innerBag{
				name:   "dim olive",
				amount: 1,
			},
		},
		"dim olive": []innerBag{
			innerBag{
				name:   "bright lavender",
				amount: 1,
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
