package day11part1

// PositionType represents the type of a position in a grid
type PositionType string

const (
	// Floor represents floor in a grid
	Floor PositionType = "."
	// EmptySeat represents an empty seat in a grid
	EmptySeat = "L"
	// OccupiedSeat represents an occupied saet in a grid
	OccupiedSeat = "#"
)

// SeatGrid represents a grid of seats
type SeatGrid map[int]map[int]PositionType

// GetOccupiedSeatsCountByGrid returns occupied seats count by grid
func GetOccupiedSeatsCountByGrid(grid SeatGrid) int {
	occupiedSeatsCount := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == OccupiedSeat {
				occupiedSeatsCount++
			}
		}
	}

	return occupiedSeatsCount
}

func simulateSeatingArea(grid SeatGrid) SeatGrid {
	updatedGrid := DeepCopyGrid(grid)
	isGridUpdated := false

	for y := range grid {
		for x := range grid[y] {
			switch grid[y][x] {
			case EmptySeat:
				if getNumberOfOccupiedAdjacentSeats(grid, y, x) == 0 {
					updatedGrid[y][x] = OccupiedSeat
					isGridUpdated = true
				}
				break
			case OccupiedSeat:
				if getNumberOfOccupiedAdjacentSeats(grid, y, x) >= 4 {
					updatedGrid[y][x] = EmptySeat
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

func getNumberOfOccupiedAdjacentSeats(grid SeatGrid, y int, x int) int {
	numberOfOccupiedSeats := 0
	adjacentSeats := []PositionType{
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
		if seat == OccupiedSeat {
			numberOfOccupiedSeats++
		}
	}

	return numberOfOccupiedSeats
}

// ConvertStrArrToGrid converts []string to seatGrid
func ConvertStrArrToGrid(strArr []string) SeatGrid {
	grid := make(SeatGrid, len(strArr))

	for i, str := range strArr {
		grid[i] = make(map[int]PositionType, len(str))

		for j, char := range str {
			positionType := convertCharToPositionType(string(char))
			grid[i][j] = positionType
		}
	}

	return grid
}

func convertCharToPositionType(char string) PositionType {
	switch char {
	case string(Floor):
		return Floor
	case string(EmptySeat):
		return EmptySeat
	case string(OccupiedSeat):
		return OccupiedSeat
	}

	panic("Could not convert char to position type")
}

// DeepCopyGrid performs a deep copy of a grid
func DeepCopyGrid(source SeatGrid) SeatGrid {
	target := make(SeatGrid, len(source))

	for outerKey, outerValue := range source {
		target[outerKey] = make(map[int]PositionType, len(outerValue))

		for innerKey, innerValue := range outerValue {
			target[outerKey][innerKey] = innerValue
		}
	}

	return target
}
