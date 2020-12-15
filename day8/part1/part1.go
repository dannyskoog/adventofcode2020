package day8part1

import (
	"log"
	"strconv"
	"strings"
)

type Operation string

const (
	Accumulate Operation = "acc"
	Jump                 = "jmp"
	None                 = "nop"
)

type Instruction struct {
	Operation Operation
	Argument  int
}

func getAccumulatedValueBeforeInstructionsAreRepeated(instructions []Instruction) int {
	visitedIndices := make(map[int]bool)
	currentIndex := 0
	accumulatedValue := 0

	for {
		if visitedIndices[currentIndex] {
			break
		}

		visitedIndices[currentIndex] = true

		currentInstruction := instructions[currentIndex]

		switch currentInstruction.Operation {
		case Jump:
			currentIndex += currentInstruction.Argument
			break
		case Accumulate:
			currentIndex++
			accumulatedValue += currentInstruction.Argument
		case None:
			currentIndex++
		}
	}

	return accumulatedValue
}

func ConvertStrArrToInstructions(strArr []string) []Instruction {
	instructions := []Instruction{}

	for _, str := range strArr {
		instruction := convertStrToInstruction(str)
		instructions = append(instructions, instruction)
	}

	return instructions
}

func convertStrToInstruction(str string) Instruction {
	splitted := strings.Split(str, " ")
	operation := convertStrToOperation(splitted[0])
	argument, _ := strconv.Atoi(splitted[1])

	return Instruction{
		operation,
		argument,
	}
}

func convertStrToOperation(str string) Operation {
	var operation Operation

	switch str {
	case string(Accumulate):
		operation = Accumulate
		break
	case string(Jump):
		operation = Jump
		break
	case string(None):
		operation = None
		break
	default:
		log.Fatal("String didn't match any operation")
	}

	return operation
}
