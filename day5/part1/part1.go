package day5part1

import (
	"math"
	"sort"
)

// RowDirection represents in which direction to go when traversing rows
type RowDirection int

const (
	// RowDirectionFront represents a forward row direction
	RowDirectionFront RowDirection = iota
	// RowDirectionBack represents a backward row direction
	RowDirectionBack
)

// ColumnDirection represents in which direction to go when traversing columns
type ColumnDirection int

const (
	// ColumnDirectionLeft represents a left column direction
	ColumnDirectionLeft ColumnDirection = iota
	// ColumnDirectionRight represents a right column direction
	ColumnDirectionRight
)

// BoardingPass represents a boarding pass
type BoardingPass struct {
	RowDirections    []RowDirection
	ColumnDirections []ColumnDirection
}

// GetSeatIDFromRowAndColumnNumbers calculcates the seat ID based on a row- and column number
func GetSeatIDFromRowAndColumnNumbers(rowNumber int, columnNumber int) int {
	return rowNumber*8 + columnNumber
}

// GetRowNumberFromBoardingPass gets the row number for a boarding pass
func GetRowNumberFromBoardingPass(boardingPass BoardingPass) int {
	rangeStart, rangeEnd := 0.0, 127.0

	for _, rowDirection := range boardingPass.RowDirections {
		rangeReduction := math.Ceil((rangeEnd - rangeStart) / 2)

		if rowDirection == RowDirectionFront {
			rangeEnd -= rangeReduction
		} else {
			rangeStart += rangeReduction
		}
	}

	if rangeStart != rangeEnd {
		panic("rangeStart and rangeEnd were expected to be equal")
	}

	return int(rangeStart)
}

// GetColumnNumberFromBoardingPass gets the column number for a boarding pass
func GetColumnNumberFromBoardingPass(boardingPass BoardingPass) int {
	rangeStart, rangeEnd := 0.0, 7.0

	for _, columnDirection := range boardingPass.ColumnDirections {
		rangeReduction := math.Ceil((rangeEnd - rangeStart) / 2)

		if columnDirection == ColumnDirectionLeft {
			rangeEnd -= rangeReduction
		} else {
			rangeStart += rangeReduction
		}
	}

	if rangeStart != rangeEnd {
		panic("rangeStart and rangeEnd were expected to be equal")
	}

	return int(rangeStart)
}

// ConvertStrArrToBoardingPasses converts a []string to []BoardingPass
func ConvertStrArrToBoardingPasses(strArr []string) []BoardingPass {
	boardingPasses := []BoardingPass{}

	for _, str := range strArr {
		boardingPass := convertStrToBoardingPass(str)
		boardingPasses = append(boardingPasses, boardingPass)
	}

	return boardingPasses
}

func getHighestSeatIDFromBoardingPasses(boardingPasses []BoardingPass, numberOfRows int, numberOfColumns int) int {
	seatIDs := []int{}

	for _, boardingPass := range boardingPasses {
		rowNumber := GetRowNumberFromBoardingPass(boardingPass)
		columnNumber := GetColumnNumberFromBoardingPass(boardingPass)
		seatID := GetSeatIDFromRowAndColumnNumbers(rowNumber, columnNumber)
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)

	return seatIDs[len(seatIDs)-1]
}

func convertStrToBoardingPass(str string) BoardingPass {
	boardingPass := BoardingPass{
		RowDirections:    []RowDirection{},
		ColumnDirections: []ColumnDirection{},
	}

	for i, char := range str {
		charStr := string(char)

		if i < 7 {
			rowDirection := convertCharToRowDirection(charStr)
			boardingPass.RowDirections = append(boardingPass.RowDirections, rowDirection)
		} else {
			columnDirection := convertCharToColumnDirection(charStr)
			boardingPass.ColumnDirections = append(boardingPass.ColumnDirections, columnDirection)
		}
	}

	return boardingPass
}

func convertCharToRowDirection(char string) RowDirection {
	if char == "F" {
		return RowDirectionFront
	}

	return RowDirectionBack
}

func convertCharToColumnDirection(char string) ColumnDirection {
	if char == "L" {
		return ColumnDirectionLeft
	}

	return ColumnDirectionRight
}
