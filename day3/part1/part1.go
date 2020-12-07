package day3part1

// TobogganMapSlope represents a slope to take on a TobogganMap
type TobogganMapSlope struct {
	RightSteps int
	DownSteps  int
}

type tobogganMapRowPosition int

const (
	openSquare tobogganMapRowPosition = iota
	tree
)

type tobogganMapRow struct {
	positions []tobogganMapRowPosition
}

// TobogganMap represents a map with rows filled with positions of either open squares or trees
type TobogganMap struct {
	rows []tobogganMapRow
}

// CalculateNumberOfTreesEncounteredInTobogganMapBySlope calculcates the number of trees encountered in a TobogganMap by slope
func CalculateNumberOfTreesEncounteredInTobogganMapBySlope(tobogganMap TobogganMap, slope TobogganMapSlope) int {
	numberOfTrees := 0
	x := 0

	for y := 0; y < len(tobogganMap.rows); y += slope.DownSteps {
		if tobogganMap.rows[y].positions[x] == tree {
			numberOfTrees++
		}

		x = (x + slope.RightSteps) % len(tobogganMap.rows[y].positions)
	}

	return numberOfTrees
}

// ConvertStrArrToTobogganMap converts a []string into a TobogganMap
func ConvertStrArrToTobogganMap(strArr []string) TobogganMap {
	tobogganMap := TobogganMap{
		rows: []tobogganMapRow{},
	}

	for _, str := range strArr {
		mapRow := convertStrToTobogganMapRow(str)
		tobogganMap.rows = append(tobogganMap.rows, mapRow)
	}

	return tobogganMap
}

func convertStrToTobogganMapRow(str string) tobogganMapRow {
	mapRow := tobogganMapRow{
		positions: []tobogganMapRowPosition{},
	}

	for _, char := range str {
		charStr := string(char)
		var position tobogganMapRowPosition

		if charStr == "." {
			position = openSquare
		} else {
			position = tree
		}

		mapRow.positions = append(mapRow.positions, position)
	}

	return mapRow
}
