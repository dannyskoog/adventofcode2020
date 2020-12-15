package day8part2

import (
	day8part1 "adventofcode2020/day8/part1"
	"errors"
	"fmt"
)

func fixProgramAndGetAccumulatedValue(instructions []day8part1.Instruction) int {
	lastModifiedIndex := -1
	instructionsCopy := make([]day8part1.Instruction, len(instructions))

	for {
		value, error := getAccumulatedValueWhenProgramEnds(instructionsCopy)

		if error == nil {
			return value
		}

		copy(instructionsCopy, instructions)

		index, instruction := findNextJumpOrNoneInstruction(instructionsCopy, lastModifiedIndex)
		lastModifiedIndex = index

		fmt.Println(instructionsCopy[index] == instructions[index])

		if instruction.Operation == day8part1.Jump {
			instructionsCopy[index] = day8part1.Instruction{
				Operation: day8part1.None,
				Argument:  instruction.Argument,
			}
		} else {
			instructionsCopy[index] = day8part1.Instruction{
				Operation: day8part1.Jump,
				Argument:  instruction.Argument,
			}
		}
	}
}

func findNextJumpOrNoneInstruction(instructions []day8part1.Instruction, startIndex int) (int, day8part1.Instruction) {
	for index, instruction := range instructions {
		if index > startIndex {
			if instruction.Operation == day8part1.Jump || (instruction.Operation == day8part1.None && instruction.Argument != 0) {
				return index, instruction
			}
		}
	}

	panic("No next jump or none instruction was found")
}

func getAccumulatedValueWhenProgramEnds(instructions []day8part1.Instruction) (int, error) {
	visitedIndices := make(map[int]bool)
	currentIndex := 0
	accumulatedValue := 0
	programTerminateIndex := len(instructions)

	for {
		if currentIndex == programTerminateIndex {
			return accumulatedValue, nil
		} else if visitedIndices[currentIndex] {
			return -1, errors.New("Program did never end")
		}

		currentInstruction := instructions[currentIndex]
		visitedIndices[currentIndex] = true

		switch currentInstruction.Operation {
		case day8part1.Jump:
			currentIndex += currentInstruction.Argument
			break
		case day8part1.Accumulate:
			currentIndex++
			accumulatedValue += currentInstruction.Argument
			break
		case day8part1.None:
			currentIndex++
			break
		}

	}
}
