package day8part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getInstructionsFromTextFile(path string, fileName string) []Instruction {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	instructions := ConvertStrArrToInstructions(strArr)

	return instructions
}

func TestConvertStrToOperation(t *testing.T) {
	tests := []struct {
		in   string
		want Operation
	}{
		{in: "acc", want: Accumulate},
		{in: "jmp", want: Jump},
		{in: "nop", want: None},
	}

	for _, e := range tests {
		got := convertStrToOperation(e.in)

		if got != e.want {
			t.Errorf("Operation was incorrect for %s, got: %s, want: %s", e.in, got, e.want)
		}
	}
}

func TestConvertStrToInstruction(t *testing.T) {
	tests := []struct {
		in   string
		want Instruction
	}{
		{
			in: "acc +1",
			want: Instruction{
				Operation: Accumulate,
				Argument:  1,
			},
		},
		{
			in: "jmp -3",
			want: Instruction{
				Operation: Jump,
				Argument:  -3,
			},
		},
		{
			in: "nop +0",
			want: Instruction{
				Operation: None,
				Argument:  0,
			},
		},
	}

	for _, e := range tests {
		got := convertStrToInstruction(e.in)

		if !reflect.DeepEqual(got, e.want) {
			t.Errorf("Instruction was incorrect for %s, got: %+v, want: %+v", e.in, got, e.want)
		}
	}
}

func TestConvertStrArrToInstructions(t *testing.T) {
	want := []Instruction{
		Instruction{
			Operation: None,
			Argument:  0,
		},
		Instruction{
			Operation: Accumulate,
			Argument:  1,
		},
		Instruction{
			Operation: Jump,
			Argument:  4,
		},
	}
	got := ConvertStrArrToInstructions([]string{
		"nop +0",
		"acc +1",
		"jmp +4",
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Instructions were incorrect, got: %+v, want: %+v", got, want)
	}
}

/*
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
*/

func TestGetAccumulatedValueBeforeInstructionsAreRepeated(t *testing.T) {
	tests := []struct {
		in   []Instruction
		want int
	}{
		{
			in: []Instruction{
				Instruction{
					Operation: None,
					Argument:  0,
				},
				Instruction{
					Operation: Accumulate,
					Argument:  1,
				},
				Instruction{
					Operation: Jump,
					Argument:  4,
				},
				Instruction{
					Operation: Accumulate,
					Argument:  3,
				},
				Instruction{
					Operation: Jump,
					Argument:  -3,
				},
				Instruction{
					Operation: Accumulate,
					Argument:  -99,
				},
				Instruction{
					Operation: Accumulate,
					Argument:  1,
				},
				Instruction{
					Operation: Jump,
					Argument:  -4,
				},
				Instruction{
					Operation: Accumulate,
					Argument:  6,
				},
			},
			want: 5,
		},
		{
			in:   getInstructionsFromTextFile("../", "input.txt"),
			want: 1782,
		},
	}

	for _, e := range tests {
		got := getAccumulatedValueBeforeInstructionsAreRepeated(e.in)

		if !reflect.DeepEqual(got, e.want) {
			t.Errorf("Accumulated value was incorrect, got: %d, want: %d", got, e.want)
		}
	}
}
