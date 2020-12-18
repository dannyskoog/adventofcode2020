package day11part2

import (
	day11part1 "adventofcode2020/day11/part1"
)

type direction int

const (
	right = iota
	left
	top
	topRight
	topLeft
	bottom
	bottomRight
	bottomLeft
)

func simulateSeatingArea(grid day11part1.SeatGrid) day11part1.SeatGrid {
	updatedGrid := day11part1.DeepCopyGrid(grid)
	isGridUpdated := false

	for y := range grid {
		for x := range grid[y] {
			switch grid[y][x] {
			case day11part1.EmptySeat:
				if getNumberOfOccupiedSeatsInSight(grid, y, x) == 0 {
					updatedGrid[y][x] = day11part1.OccupiedSeat
					isGridUpdated = true
				}
				break
			case day11part1.OccupiedSeat:
				if getNumberOfOccupiedSeatsInSight(grid, y, x) >= 5 {
					updatedGrid[y][x] = day11part1.EmptySeat
					isGridUpdated = true
				}
				break
			}
		}
	}

	if !isGridUpdated {
		return grid
	}

	return simulateSeatingArea(updatedGrid)
}

func getNumberOfOccupiedSeatsInSight(grid day11part1.SeatGrid, y int, x int) int {
	numberOfOccupiedSeats := 0
	directions := []direction{right, left, top, topRight, topLeft, bottom, bottomRight, bottomLeft}

	for _, direction := range directions {
		numberOfOccupiedSeats += getNumberOfOccupiedSeatsInSightByDirection(grid, y, x, direction)
	}

	return numberOfOccupiedSeats
}

func getNumberOfOccupiedSeatsInSightByDirection(grid day11part1.SeatGrid, y int, x int, direction direction) int {
	switch direction {
	case right:
		x++
		break
	case left:
		x--
		break
	case top:
		y--
		break
	case topRight:
		y--
		x++
		break
	case topLeft:
		y--
		x--
		break
	case bottom:
		y++
		break
	case bottomRight:
		y++
		x++
		break
	case bottomLeft:
		y++
		x--
		break
	}

	if grid[y][x] == "" || grid[y][x] == day11part1.EmptySeat {
		return 0
	}

	if grid[y][x] == day11part1.OccupiedSeat {
		return 1
	}

	return getNumberOfOccupiedSeatsInSightByDirection(grid, y, x, direction)
}
