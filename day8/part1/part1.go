package day8part1

import (
	"log"
	"strconv"
	"strings"
)

// Operation represents which action for the instruction to perform
type Operation string

const (
	// Accumulate represents an accumulative action
	Accumulate Operation = "acc"
	// Jump represents a moving action
	Jump = "jmp"
	// None represents a noop action
	None = "nop"
)

// Instruction represents which operation to take and with which argument
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

// ConvertStrArrToInstructions converts []string to []Instruction
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
