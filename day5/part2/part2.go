package day5part2

import (
	day5part1 "adventofcode2020/day5/part1"
)

type seatIDs map[int]bool

func getMySeatIDBasedOnOtherBoardingPasses(boardingPasses []day5part1.BoardingPass) int {
	seatIDs := getSeatIDsFromBoardingPasses(boardingPasses)

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			seatID := day5part1.GetSeatIDFromRowAndColumnNumbers(i, j)

			if !seatIDs[seatID] && seatIDs[seatID-1] && seatIDs[seatID+1] {
				return seatID
			}
		}
	}

	return -1
}

func getSeatIDsFromBoardingPasses(boardingPasses []day5part1.BoardingPass) seatIDs {
	seatIDs := make(seatIDs)

	for _, boardingPass := range boardingPasses {
		rowNumber := day5part1.GetRowNumberFromBoardingPass(boardingPass)
		columnNumber := day5part1.GetColumnNumberFromBoardingPass(boardingPass)
		seatID := day5part1.GetSeatIDFromRowAndColumnNumbers(rowNumber, columnNumber)
		seatIDs[seatID] = true
	}

	return seatIDs
}
