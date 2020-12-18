package day11part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getSeatGridFromTextFile(path string, fileName string) seatGrid {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	grid := convertStrArrToGrid(strArr)

	return grid
}

func TestConvertCharToPositionType(t *testing.T) {
	tests := []struct {
		in   string
		want positionType
	}{
		{
			in:   ".",
			want: floor,
		},
		{
			in:   "L",
			want: emptySeat,
		},
		{
			in:   "#",
			want: occupiedSeat,
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
	want := seatGrid{
		0: {0: occupiedSeat, 1: floor, 2: occupiedSeat, 3: emptySeat, 4: floor, 5: emptySeat, 6: occupiedSeat, 7: floor, 8: occupiedSeat, 9: occupiedSeat},
		1: {0: occupiedSeat, 1: emptySeat, 2: emptySeat, 3: emptySeat, 4: occupiedSeat, 5: emptySeat, 6: emptySeat, 7: floor, 8: emptySeat, 9: occupiedSeat},
		2: {0: emptySeat, 1: floor, 2: emptySeat, 3: floor, 4: emptySeat, 5: floor, 6: floor, 7: occupiedSeat, 8: floor, 9: floor},
		3: {0: occupiedSeat, 1: emptySeat, 2: emptySeat, 3: emptySeat, 4: floor, 5: occupiedSeat, 6: occupiedSeat, 7: floor, 8: emptySeat, 9: occupiedSeat},
		4: {0: occupiedSeat, 1: floor, 2: emptySeat, 3: emptySeat, 4: floor, 5: emptySeat, 6: emptySeat, 7: floor, 8: emptySeat, 9: emptySeat},
		5: {0: occupiedSeat, 1: floor, 2: emptySeat, 3: emptySeat, 4: occupiedSeat, 5: emptySeat, 6: occupiedSeat, 7: floor, 8: occupiedSeat, 9: occupiedSeat},
		6: {0: floor, 1: floor, 2: emptySeat, 3: floor, 4: emptySeat, 5: floor, 6: floor, 7: floor, 8: floor, 9: floor},
		7: {0: occupiedSeat, 1: emptySeat, 2: occupiedSeat, 3: emptySeat, 4: emptySeat, 5: emptySeat, 6: emptySeat, 7: occupiedSeat, 8: emptySeat, 9: occupiedSeat},
		8: {0: occupiedSeat, 1: floor, 2: emptySeat, 3: emptySeat, 4: emptySeat, 5: emptySeat, 6: emptySeat, 7: emptySeat, 8: floor, 9: emptySeat},
		9: {0: occupiedSeat, 1: floor, 2: occupiedSeat, 3: emptySeat, 4: occupiedSeat, 5: emptySeat, 6: occupiedSeat, 7: floor, 8: occupiedSeat, 9: occupiedSeat},
	}
	got := convertStrArrToGrid([]string{
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
	grid := convertStrArrToGrid([]string{
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
			grid seatGrid
			y    int
			x    int
		}
		want int
	}{
		{
			in: struct {
				grid seatGrid
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
				grid seatGrid
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
				grid seatGrid
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
				grid seatGrid
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
	grid := convertStrArrToGrid([]string{
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
	want := seatGrid{
		0: {0: occupiedSeat, 1: floor, 2: occupiedSeat, 3: emptySeat, 4: floor, 5: emptySeat, 6: occupiedSeat, 7: floor, 8: occupiedSeat, 9: occupiedSeat},
		1: {0: occupiedSeat, 1: emptySeat, 2: emptySeat, 3: emptySeat, 4: occupiedSeat, 5: emptySeat, 6: emptySeat, 7: floor, 8: emptySeat, 9: occupiedSeat},
		2: {0: emptySeat, 1: floor, 2: occupiedSeat, 3: floor, 4: emptySeat, 5: floor, 6: floor, 7: occupiedSeat, 8: floor, 9: floor},
		3: {0: occupiedSeat, 1: emptySeat, 2: occupiedSeat, 3: occupiedSeat, 4: floor, 5: occupiedSeat, 6: occupiedSeat, 7: floor, 8: emptySeat, 9: occupiedSeat},
		4: {0: occupiedSeat, 1: floor, 2: occupiedSeat, 3: emptySeat, 4: floor, 5: emptySeat, 6: emptySeat, 7: floor, 8: emptySeat, 9: emptySeat},
		5: {0: occupiedSeat, 1: floor, 2: occupiedSeat, 3: emptySeat, 4: occupiedSeat, 5: emptySeat, 6: occupiedSeat, 7: floor, 8: occupiedSeat, 9: occupiedSeat},
		6: {0: floor, 1: floor, 2: emptySeat, 3: floor, 4: emptySeat, 5: floor, 6: floor, 7: floor, 8: floor, 9: floor},
		7: {0: occupiedSeat, 1: emptySeat, 2: occupiedSeat, 3: emptySeat, 4: occupiedSeat, 5: occupiedSeat, 6: emptySeat, 7: occupiedSeat, 8: emptySeat, 9: occupiedSeat},
		8: {0: occupiedSeat, 1: floor, 2: emptySeat, 3: emptySeat, 4: emptySeat, 5: emptySeat, 6: emptySeat, 7: emptySeat, 8: floor, 9: emptySeat},
		9: {0: occupiedSeat, 1: floor, 2: occupiedSeat, 3: emptySeat, 4: occupiedSeat, 5: emptySeat, 6: occupiedSeat, 7: floor, 8: occupiedSeat, 9: occupiedSeat},
	}
	got := simulateSeatingArea(grid)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Seating area was incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestGetOccupiedSeatsCountByGrid(t *testing.T) {
	tests := []struct {
		in   seatGrid
		want int
	}{
		{
			in: simulateSeatingArea(
				convertStrArrToGrid([]string{
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
		got := getOccupiedSeatsCountByGrid(e.in)

		if got != e.want {
			t.Errorf("Occupied seats count was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
