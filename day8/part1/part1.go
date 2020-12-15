package day8part1

import (
	"log"
	"strconv"
	"strings"
)

type operation string

const (
	accumulate operation = "acc"
	jump                 = "jmp"
	none                 = "nop"
)

type instruction struct {
	operation operation
	argument  int
}

func getAccumulatedValueBeforeInstructionsAreRepeated(instructions []instruction) int {
	visitedIndices := make(map[int]bool)
	currentIndex := 0
	accumulatedValue := 0

	for {
		if visitedIndices[currentIndex] {
			break
		}

		visitedIndices[currentIndex] = true

		currentInstruction := instructions[currentIndex]

		switch currentInstruction.operation {
		case jump:
			currentIndex += currentInstruction.argument
			break
		case accumulate:
			currentIndex++
			accumulatedValue += currentInstruction.argument
		case none:
			currentIndex++
		}
	}

	return accumulatedValue
}

func convertStrArrToInstructions(strArr []string) []instruction {
	instructions := []instruction{}

	for _, str := range strArr {
		instruction := convertStrToInstruction(str)
		instructions = append(instructions, instruction)
	}

	return instructions
}

func convertStrToInstruction(str string) instruction {
	splitted := strings.Split(str, " ")
	operation := convertStrToOperation(splitted[0])
	argument, _ := strconv.Atoi(splitted[1])

	return instruction{
		operation,
		argument,
	}
}

func convertStrToOperation(str string) operation {
	var operation operation

	switch str {
	case string(accumulate):
		operation = accumulate
		break
	case string(jump):
		operation = jump
		break
	case string(none):
		operation = none
		break
	default:
		log.Fatal("String didn't match any operation")
	}

	return operation
}
