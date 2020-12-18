package day11part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getSeatGridFromTextFile(path string, fileName string) SeatGrid {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	grid := ConvertStrArrToGrid(strArr)

	return grid
}

func TestConvertCharToPositionType(t *testing.T) {
	tests := []struct {
		in   string
		want PositionType
	}{
		{
			in:   ".",
			want: Floor,
		},
		{
			in:   "L",
			want: EmptySeat,
		},
		{
			in:   "#",
			want: OccupiedSeat,
		},
	}

	for _, e := range tests {
		got := convertCharToPositionType(e.in)

		if got != e.want {
			t.Errorf("Position type for %s was incorrect, got: %s, want: %s", e.in, got, e.want)
		}
	}
}

func TestConvertStrArrToGrid(t *testing.T) {
	want := SeatGrid{
		0: {0: OccupiedSeat, 1: Floor, 2: OccupiedSeat, 3: EmptySeat, 4: Floor, 5: EmptySeat, 6: OccupiedSeat, 7: Floor, 8: OccupiedSeat, 9: OccupiedSeat},
		1: {0: OccupiedSeat, 1: EmptySeat, 2: EmptySeat, 3: EmptySeat, 4: OccupiedSeat, 5: EmptySeat, 6: EmptySeat, 7: Floor, 8: EmptySeat, 9: OccupiedSeat},
		2: {0: EmptySeat, 1: Floor, 2: EmptySeat, 3: Floor, 4: EmptySeat, 5: Floor, 6: Floor, 7: OccupiedSeat, 8: Floor, 9: Floor},
		3: {0: OccupiedSeat, 1: EmptySeat, 2: EmptySeat, 3: EmptySeat, 4: Floor, 5: OccupiedSeat, 6: OccupiedSeat, 7: Floor, 8: EmptySeat, 9: OccupiedSeat},
		4: {0: OccupiedSeat, 1: Floor, 2: EmptySeat, 3: EmptySeat, 4: Floor, 5: EmptySeat, 6: EmptySeat, 7: Floor, 8: EmptySeat, 9: EmptySeat},
		5: {0: OccupiedSeat, 1: Floor, 2: EmptySeat, 3: EmptySeat, 4: OccupiedSeat, 5: EmptySeat, 6: OccupiedSeat, 7: Floor, 8: OccupiedSeat, 9: OccupiedSeat},
		6: {0: Floor, 1: Floor, 2: EmptySeat, 3: Floor, 4: EmptySeat, 5: Floor, 6: Floor, 7: Floor, 8: Floor, 9: Floor},
		7: {0: OccupiedSeat, 1: EmptySeat, 2: OccupiedSeat, 3: EmptySeat, 4: EmptySeat, 5: EmptySeat, 6: EmptySeat, 7: OccupiedSeat, 8: EmptySeat, 9: OccupiedSeat},
		8: {0: OccupiedSeat, 1: Floor, 2: EmptySeat, 3: EmptySeat, 4: EmptySeat, 5: EmptySeat, 6: EmptySeat, 7: EmptySeat, 8: Floor, 9: EmptySeat},
		9: {0: OccupiedSeat, 1: Floor, 2: OccupiedSeat, 3: EmptySeat, 4: OccupiedSeat, 5: EmptySeat, 6: OccupiedSeat, 7: Floor, 8: OccupiedSeat, 9: OccupiedSeat},
	}
	got := ConvertStrArrToGrid([]string{
		"#.#L.L#.##",
		"#LLL#LL.L#",
		"L.L.L..#..",
		"#LLL.##.L#",
		"#.LL.LL.LL",
		"#.LL#L#.##",
		"..L.L.....",
		"#L#LLLL#L#",
		"#.LLLLLL.L",
		"#.#L#L#.##",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Grid was incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestGetNumberOfOccupiedAdjacentSeats(t *testing.T) {
	grid := ConvertStrArrToGrid([]string{
		"#.#L.L#.##",
		"#LLL#LL.L#",
		"L.L.L..#..",
		"#LLL.##.L#",
		"#.LL.LL.LL",
		"#.LL#L#.##",
		"..L.L.....",
		"#L#LLLL#L#",
		"#.LLLLLL.L",
		"#.#L#L#.##",
	})
	tests := []struct {
		in struct {
			grid SeatGrid
			y    int
			x    int
		}
		want int
	}{
		{
			in: struct {
				grid SeatGrid
				y    int
				x    int
			}{
				grid: grid,
				y:    3,
				x:    2,
			},
			want: 0,
		},
		{
			in: struct {
				grid SeatGrid
				y    int
				x    int
			}{
				grid: grid,
				y:    0,
				x:    0,
			},
			want: 1,
		},
		{
			in: struct {
				grid SeatGrid
				y    int
				x    int
			}{
				grid: grid,
				y:    5,
				x:    7,
			},
			want: 2,
		},
		{
			in: struct {
				grid SeatGrid
				y    int
				x    int
			}{
				grid: grid,
				y:    1,
				x:    8,
			},
			want: 4,
		},
	}

	for _, e := range tests {
		got := getNumberOfOccupiedAdjacentSeats(e.in.grid, e.in.y, e.in.x)

		if got != e.want {
			t.Errorf("Number was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}

func TestSimulateSeatingArea(t *testing.T) {
	grid := ConvertStrArrToGrid([]string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	})
	want := SeatGrid{
		0: {0: OccupiedSeat, 1: Floor, 2: OccupiedSeat, 3: EmptySeat, 4: Floor, 5: EmptySeat, 6: OccupiedSeat, 7: Floor, 8: OccupiedSeat, 9: OccupiedSeat},
		1: {0: OccupiedSeat, 1: EmptySeat, 2: EmptySeat, 3: EmptySeat, 4: OccupiedSeat, 5: EmptySeat, 6: EmptySeat, 7: Floor, 8: EmptySeat, 9: OccupiedSeat},
		2: {0: EmptySeat, 1: Floor, 2: OccupiedSeat, 3: Floor, 4: EmptySeat, 5: Floor, 6: Floor, 7: OccupiedSeat, 8: Floor, 9: Floor},
		3: {0: OccupiedSeat, 1: EmptySeat, 2: OccupiedSeat, 3: OccupiedSeat, 4: Floor, 5: OccupiedSeat, 6: OccupiedSeat, 7: Floor, 8: EmptySeat, 9: OccupiedSeat},
		4: {0: OccupiedSeat, 1: Floor, 2: OccupiedSeat, 3: EmptySeat, 4: Floor, 5: EmptySeat, 6: EmptySeat, 7: Floor, 8: EmptySeat, 9: EmptySeat},
		5: {0: OccupiedSeat, 1: Floor, 2: OccupiedSeat, 3: EmptySeat, 4: OccupiedSeat, 5: EmptySeat, 6: OccupiedSeat, 7: Floor, 8: OccupiedSeat, 9: OccupiedSeat},
		6: {0: Floor, 1: Floor, 2: EmptySeat, 3: Floor, 4: EmptySeat, 5: Floor, 6: Floor, 7: Floor, 8: Floor, 9: Floor},
		7: {0: OccupiedSeat, 1: EmptySeat, 2: OccupiedSeat, 3: EmptySeat, 4: OccupiedSeat, 5: OccupiedSeat, 6: EmptySeat, 7: OccupiedSeat, 8: EmptySeat, 9: OccupiedSeat},
		8: {0: OccupiedSeat, 1: Floor, 2: EmptySeat, 3: EmptySeat, 4: EmptySeat, 5: EmptySeat, 6: EmptySeat, 7: EmptySeat, 8: Floor, 9: EmptySeat},
		9: {0: OccupiedSeat, 1: Floor, 2: OccupiedSeat, 3: EmptySeat, 4: OccupiedSeat, 5: EmptySeat, 6: OccupiedSeat, 7: Floor, 8: OccupiedSeat, 9: OccupiedSeat},
	}
	got := simulateSeatingArea(grid)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Seating area was incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestGetOccupiedSeatsCountByGrid(t *testing.T) {
	tests := []struct {
		in   SeatGrid
		want int
	}{
		{
			in: simulateSeatingArea(
				ConvertStrArrToGrid([]string{
					"L.LL.LL.LL",
					"LLLLLLL.LL",
					"L.L.L..L..",
					"LLLL.LL.LL",
					"L.LL.LL.LL",
					"L.LLLLL.LL",
					"..L.L.....",
					"LLLLLLLLLL",
					"L.LLLLLL.L",
					"L.LLLLL.LL",
				}),
			),
			want: 37,
		},
		{
			in: simulateSeatingArea(
				getSeatGridFromTextFile("../", "input.txt"),
			),
			want: 2152,
		},
	}

	for _, e := range tests {
		got := GetOccupiedSeatsCountByGrid(e.in)

		if got != e.want {
			t.Errorf("Occupied seats count was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
