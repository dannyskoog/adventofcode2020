package day5part1

import (
	"math"
	"sort"
)

type rowDirection int

const (
	rowDirectionFront rowDirection = iota
	rowDirectionBack
)

type columnDirection int

const (
	columnDirectionLeft columnDirection = iota
	columnDirectionRight
)

type boardingPass struct {
	rowDirections    []rowDirection
	columnDirections []columnDirection
}

func getHighestSeatIDFromBoardingPasses(boardingPasses []boardingPass) int {
	seatIDs := []int{}

	for _, boardingPass := range boardingPasses {
		seatID := getSeatIDFromBoardingPass(boardingPass)
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)

	return seatIDs[len(seatIDs)-1]
}

func getSeatIDFromBoardingPass(boardingPass boardingPass) int {
	rowNumber := getRowNumberFromBoardingPass(boardingPass)
	columnNumber := getColumnNumberFromBoardingPass(boardingPass)
	seatID := rowNumber*8 + columnNumber

	return seatID
}

func getRowNumberFromBoardingPass(boardingPass boardingPass) int {
	rangeStart, rangeEnd := 0.0, 127.0

	for _, rowDirection := range boardingPass.rowDirections {
		rangeReduction := math.Ceil((rangeEnd - rangeStart) / 2)

		if rowDirection == rowDirectionFront {
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

func getColumnNumberFromBoardingPass(boardingPass boardingPass) int {
	rangeStart, rangeEnd := 0.0, 7.0

	for _, columnDirection := range boardingPass.columnDirections {
		rangeReduction := math.Ceil((rangeEnd - rangeStart) / 2)

		if columnDirection == columnDirectionLeft {
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

func convertStrArrToBoardingPasses(strArr []string) []boardingPass {
	boardingPasses := []boardingPass{}

	for _, str := range strArr {
		boardingPass := convertStrToBoardingPass(str)
		boardingPasses = append(boardingPasses, boardingPass)
	}

	return boardingPasses
}

func convertStrToBoardingPass(str string) boardingPass {
	boardingPass := boardingPass{
		rowDirections:    []rowDirection{},
		columnDirections: []columnDirection{},
	}

	for i, char := range str {
		charStr := string(char)

		if i < 7 {
			rowDirection := convertCharToRowDirection(charStr)
			boardingPass.rowDirections = append(boardingPass.rowDirections, rowDirection)
		} else {
			columnDirection := convertCharToColumnDirection(charStr)
			boardingPass.columnDirections = append(boardingPass.columnDirections, columnDirection)
		}
	}

	return boardingPass
}

func convertCharToRowDirection(char string) rowDirection {
	if char == "F" {
		return rowDirectionFront
	}

	return rowDirectionBack
}

func convertCharToColumnDirection(char string) columnDirection {
	if char == "L" {
		return columnDirectionLeft
	}

	return columnDirectionRight
}
