package day11part2

import (
	day11part1 "adventofcode2020/day11/part1"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getSeatGridFromTextFile(path string, fileName string) day11part1.SeatGrid {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	grid := day11part1.ConvertStrArrToGrid(strArr)

	return grid
}

func TestGetNumberOfOccupiedSeatInSightByDirection(t *testing.T) {
	grid := day11part1.ConvertStrArrToGrid([]string{
		".##.##L",
		"#.#.#.#",
		"##...##",
		".#LL...",
		"##LL.##",
		"#.#.#.#",
		".##.##.",
	})
	tests := []struct {
		in struct {
			y         int
			x         int
			direction direction
		}
		want int
	}{
		{
			in: struct {
				y         int
				x         int
				direction direction
			}{
				y:         3,
				x:         3,
				direction: left,
			},
			want: 0,
		},
		{
			in: struct {
				y         int
				x         int
				direction direction
			}{
				y:         4,
				x:         2,
				direction: right,
			},
			want: 0,
		},
		{
			in: struct {
				y         int
				x         int
				direction direction
			}{
				y:         0,
				x:         6,
				direction: top,
			},
			want: 0,
		},
		{
			in: struct {
				y         int
				x         int
				direction direction
			}{
				y:         3,
				x:         3,
				direction: bottomLeft,
			},
			want: 0,
		},
	}

	for _, e := range tests {
		got := getNumberOfOccupiedSeatsInSightByDirection(grid, e.in.y, e.in.x, e.in.direction)

		if got != e.want {
			t.Errorf("Number of occupied seats for y: %d and x: %d were incorrect, got: %d, want: %d", e.in.y, e.in.x, got, e.want)
		}
	}
}

func TestGetNumberOfOccupiedSeatsInSight(t *testing.T) {
	grid := day11part1.ConvertStrArrToGrid([]string{
		".##.##L",
		"#.#.#.#",
		"##...##",
		".#LL...",
		"##LL.##",
		"#.#.#.#",
		".##.##.",
	})
	tests := []struct {
		in struct {
			y int
			x int
		}
		want int
	}{
		{
			in: struct {
				y int
				x int
			}{
				y: 2,
				x: 2,
			},
			want: 5,
		},
		{
			in: struct {
				y int
				x int
			}{
				y: 5,
				x: 5,
			},
			want: 6,
		},
		{
			in: struct {
				y int
				x int
			}{
				y: 0,
				x: 6,
			},
			want: 2,
		},
	}

	for _, e := range tests {
		got := getNumberOfOccupiedSeatsInSight(grid, e.in.y, e.in.x)

		if got != e.want {
			t.Errorf("Number of occupied seats for y: %d and x: %d were incorrect, got: %d, want: %d", e.in.y, e.in.x, got, e.want)
		}
	}
}

func TestSimulateSeatingArea(t *testing.T) {
	grid := day11part1.ConvertStrArrToGrid([]string{
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
	want := day11part1.SeatGrid{
		0: {0: day11part1.OccupiedSeat, 1: day11part1.Floor, 2: day11part1.EmptySeat, 3: day11part1.OccupiedSeat, 4: day11part1.Floor, 5: day11part1.EmptySeat, 6: day11part1.OccupiedSeat, 7: day11part1.Floor, 8: day11part1.EmptySeat, 9: day11part1.OccupiedSeat},
		1: {0: day11part1.OccupiedSeat, 1: day11part1.EmptySeat, 2: day11part1.EmptySeat, 3: day11part1.EmptySeat, 4: day11part1.EmptySeat, 5: day11part1.EmptySeat, 6: day11part1.EmptySeat, 7: day11part1.Floor, 8: day11part1.EmptySeat, 9: day11part1.EmptySeat},
		2: {0: day11part1.EmptySeat, 1: day11part1.Floor, 2: day11part1.EmptySeat, 3: day11part1.Floor, 4: day11part1.EmptySeat, 5: day11part1.Floor, 6: day11part1.Floor, 7: day11part1.OccupiedSeat, 8: day11part1.Floor, 9: day11part1.Floor},
		3: {0: day11part1.OccupiedSeat, 1: day11part1.OccupiedSeat, 2: day11part1.EmptySeat, 3: day11part1.OccupiedSeat, 4: day11part1.Floor, 5: day11part1.OccupiedSeat, 6: day11part1.EmptySeat, 7: day11part1.Floor, 8: day11part1.EmptySeat, 9: day11part1.OccupiedSeat},
		4: {0: day11part1.EmptySeat, 1: day11part1.Floor, 2: day11part1.EmptySeat, 3: day11part1.OccupiedSeat, 4: day11part1.Floor, 5: day11part1.EmptySeat, 6: day11part1.EmptySeat, 7: day11part1.Floor, 8: day11part1.EmptySeat, 9: day11part1.OccupiedSeat},
		5: {0: day11part1.OccupiedSeat, 1: day11part1.Floor, 2: day11part1.EmptySeat, 3: day11part1.EmptySeat, 4: day11part1.EmptySeat, 5: day11part1.EmptySeat, 6: day11part1.OccupiedSeat, 7: day11part1.Floor, 8: day11part1.EmptySeat, 9: day11part1.EmptySeat},
		6: {0: day11part1.Floor, 1: day11part1.Floor, 2: day11part1.OccupiedSeat, 3: day11part1.Floor, 4: day11part1.EmptySeat, 5: day11part1.Floor, 6: day11part1.Floor, 7: day11part1.Floor, 8: day11part1.Floor, 9: day11part1.Floor},
		7: {0: day11part1.EmptySeat, 1: day11part1.EmptySeat, 2: day11part1.EmptySeat, 3: day11part1.OccupiedSeat, 4: day11part1.OccupiedSeat, 5: day11part1.OccupiedSeat, 6: day11part1.EmptySeat, 7: day11part1.EmptySeat, 8: day11part1.EmptySeat, 9: day11part1.OccupiedSeat},
		8: {0: day11part1.OccupiedSeat, 1: day11part1.Floor, 2: day11part1.EmptySeat, 3: day11part1.EmptySeat, 4: day11part1.EmptySeat, 5: day11part1.EmptySeat, 6: day11part1.EmptySeat, 7: day11part1.OccupiedSeat, 8: day11part1.Floor, 9: day11part1.EmptySeat},
		9: {0: day11part1.OccupiedSeat, 1: day11part1.Floor, 2: day11part1.EmptySeat, 3: day11part1.OccupiedSeat, 4: day11part1.EmptySeat, 5: day11part1.EmptySeat, 6: day11part1.OccupiedSeat, 7: day11part1.Floor, 8: day11part1.EmptySeat, 9: day11part1.OccupiedSeat},
	}
	got := simulateSeatingArea(grid)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Seating area was incorrect, got: %+v, want: %+v", got, want)
	}
}

func TestGetOccupiedSeatsCountByGrid(t *testing.T) {
	tests := []struct {
		in   day11part1.SeatGrid
		want int
	}{
		{
			in: simulateSeatingArea(
				day11part1.ConvertStrArrToGrid([]string{
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
			want: 26,
		},
		{
			in: simulateSeatingArea(
				getSeatGridFromTextFile("../", "input.txt"),
			),
			want: 1937,
		},
	}

	for _, e := range tests {
		got := day11part1.GetOccupiedSeatsCountByGrid(e.in)

		if got != e.want {
			t.Errorf("Occupied seats count was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
