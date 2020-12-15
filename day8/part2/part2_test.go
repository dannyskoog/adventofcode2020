package day8part2

import (
	day8part1 "adventofcode2020/day8/part1"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getInstructionsFromTextFile(path string, fileName string) []day8part1.Instruction {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	instructions := day8part1.ConvertStrArrToInstructions(strArr)

	return instructions
}

func TestIsProgramCorrupted(t *testing.T) {
	tests := []struct {
		in   []day8part1.Instruction
		want struct {
			value    int
			hasError bool
		}
	}{
		{
			in: []day8part1.Instruction{
				day8part1.Instruction{
					Operation: day8part1.None,
					Argument:  0,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  1,
				},
				day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  4,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  3,
				},
				day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  -3,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  -99,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  1,
				},
				day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  -4,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  6,
				},
			},
			want: struct {
				value    int
				hasError bool
			}{
				value:    -1,
				hasError: true,
			},
		},
		{
			in: []day8part1.Instruction{
				day8part1.Instruction{
					Operation: day8part1.None,
					Argument:  0,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  1,
				},
				day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  4,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  3,
				},
				day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  -3,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  -99,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  1,
				},
				day8part1.Instruction{
					Operation: day8part1.None,
					Argument:  -4,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  6,
				},
			},
			want: struct {
				value    int
				hasError bool
			}{
				value:    8,
				hasError: false,
			},
		},
	}

	for _, e := range tests {
		value, error := getAccumulatedValueWhenProgramEnds(e.in)

		if !reflect.DeepEqual(value, e.want.value) {
			t.Errorf("Value was incorrect, got: %d, want: %d", value, e.want.value)
		}

		if (error == nil && e.want.hasError) || (error != nil && !e.want.hasError) {
			t.Errorf("Error was incorrect, got: %v, want: %v", error, e.want.hasError)
		}
	}
}

func TestFindNextJumpOrNoneInstruction(t *testing.T) {
	tests := []struct {
		in struct {
			first  []day8part1.Instruction
			second int
		}
		want struct {
			index       int
			instruction day8part1.Instruction
		}
	}{
		{
			in: struct {
				first  []day8part1.Instruction
				second int
			}{
				first: []day8part1.Instruction{
					day8part1.Instruction{
						Operation: day8part1.None,
						Argument:  0,
					},
					day8part1.Instruction{
						Operation: day8part1.Accumulate,
						Argument:  1,
					},
					day8part1.Instruction{
						Operation: day8part1.Jump,
						Argument:  4,
					},
				},
				second: -1,
			},
			want: struct {
				index       int
				instruction day8part1.Instruction
			}{
				index: 2,
				instruction: day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  4,
				},
			},
		},
		{
			in: struct {
				first  []day8part1.Instruction
				second int
			}{
				first: []day8part1.Instruction{
					day8part1.Instruction{
						Operation: day8part1.Accumulate,
						Argument:  5,
					},
					day8part1.Instruction{
						Operation: day8part1.None,
						Argument:  -3,
					},
					day8part1.Instruction{
						Operation: day8part1.Jump,
						Argument:  2,
					},
				},
				second: -1,
			},
			want: struct {
				index       int
				instruction day8part1.Instruction
			}{
				index: 1,
				instruction: day8part1.Instruction{
					Operation: day8part1.None,
					Argument:  -3,
				},
			},
		},
	}

	for _, e := range tests {
		index, instruction := findNextJumpOrNoneInstruction(e.in.first, e.in.second)

		if index != e.want.index {
			t.Errorf("Next jump or none index was incorrect, got: %d, want: %d", index, e.want.index)
		}

		if !reflect.DeepEqual(instruction, e.want.instruction) {
			t.Errorf("Next jump or none instruction was incorrect, got: %+v, want: %+v", instruction, e.want.instruction)
		}
	}
}

func TestFixProgramAndGetAccumulatedValue(t *testing.T) {
	tests := []struct {
		in   []day8part1.Instruction
		want int
	}{
		{
			in: []day8part1.Instruction{
				day8part1.Instruction{
					Operation: day8part1.None,
					Argument:  0,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  1,
				},
				day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  4,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  3,
				},
				day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  -3,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  -99,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  1,
				},
				day8part1.Instruction{
					Operation: day8part1.Jump,
					Argument:  -4,
				},
				day8part1.Instruction{
					Operation: day8part1.Accumulate,
					Argument:  6,
				},
			},
			want: 8,
		},
		{
			in:   getInstructionsFromTextFile("../", "input.txt"),
			want: 797,
		},
	}

	for _, e := range tests {
		got := fixProgramAndGetAccumulatedValue(e.in)

		if got != e.want {
			t.Errorf("Accumulated value was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
