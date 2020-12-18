package day11part1

type positionType string

const (
	floor        positionType = "."
	emptySeat                 = "L"
	occupiedSeat              = "#"
)

type seatGrid map[int]map[int]positionType

func getOccupiedSeatsCountByGrid(grid seatGrid) int {
	occupiedSeatsCount := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == occupiedSeat {
				occupiedSeatsCount++
			}
		}
	}

	return occupiedSeatsCount
}

func simulateSeatingArea(grid seatGrid) seatGrid {
	updatedGrid := deepCopyGrid(grid)
	isGridUpdated := false

	for y := range grid {
		for x := range grid[y] {
			switch grid[y][x] {
			case emptySeat:
				if getNumberOfOccupiedAdjacentSeats(grid, y, x) == 0 {
					updatedGrid[y][x] = occupiedSeat
					isGridUpdated = true
				}
				break
			case occupiedSeat:
				if getNumberOfOccupiedAdjacentSeats(grid, y, x) >= 4 {
					updatedGrid[y][x] = emptySeat
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

func getNumberOfOccupiedAdjacentSeats(grid seatGrid, y int, x int) int {
	numberOfOccupiedSeats := 0
	adjacentSeats := []positionType{
		grid[y][x+1],   // Right
		grid[y][x-1],   // Left
		grid[y-1][x],   // Top
		grid[y-1][x+1], // Top right
		grid[y-1][x-1], // Top left
		grid[y+1][x],   // Bottom
		grid[y+1][x+1], // Bottom right
		grid[y+1][x-1], // Bottom left
	}

	for _, seat := range adjacentSeats {
		if seat == occupiedSeat {
			numberOfOccupiedSeats++
		}
	}

	return numberOfOccupiedSeats
}

func convertStrArrToGrid(strArr []string) seatGrid {
	grid := make(seatGrid, len(strArr))

	for i, str := range strArr {
		grid[i] = make(map[int]positionType, len(str))

		for j, char := range str {
			positionType := convertCharToPositionType(string(char))
			grid[i][j] = positionType
		}
	}

	return grid
}

func convertCharToPositionType(char string) positionType {
	switch char {
	case string(floor):
		return floor
	case string(emptySeat):
		return emptySeat
	case string(occupiedSeat):
		return occupiedSeat
	}

	panic("Could not convert char to position type")
}

func deepCopyGrid(source seatGrid) seatGrid {
	target := make(seatGrid, len(source))

	for outerKey, outerValue := range source {
		target[outerKey] = make(map[int]positionType, len(outerValue))

		for innerKey, innerValue := range outerValue {
			target[outerKey][innerKey] = innerValue
		}
	}

	return target
}
