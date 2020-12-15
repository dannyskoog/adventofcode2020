package day8part1

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/gobuffalo/packr/v2"
)

func getInstructionsFromTextFile(path string, fileName string) []instruction {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	strArr := strings.Split(str, "\r\n")
	instructions := convertStrArrToInstructions(strArr)

	return instructions
}

func TestConvertStrToOperation(t *testing.T) {
	tests := []struct {
		in   string
		want operation
	}{
		{in: "acc", want: accumulate},
		{in: "jmp", want: jump},
		{in: "nop", want: none},
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
		want instruction
	}{
		{
			in: "acc +1",
			want: instruction{
				operation: accumulate,
				argument:  1,
			},
		},
		{
			in: "jmp -3",
			want: instruction{
				operation: jump,
				argument:  -3,
			},
		},
		{
			in: "nop +0",
			want: instruction{
				operation: none,
				argument:  0,
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
	want := []instruction{
		instruction{
			operation: none,
			argument:  0,
		},
		instruction{
			operation: accumulate,
			argument:  1,
		},
		instruction{
			operation: jump,
			argument:  4,
		},
	}
	got := convertStrArrToInstructions([]string{
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
		in   []instruction
		want int
	}{
		{
			in: []instruction{
				instruction{
					operation: none,
					argument:  0,
				},
				instruction{
					operation: accumulate,
					argument:  1,
				},
				instruction{
					operation: jump,
					argument:  4,
				},
				instruction{
					operation: accumulate,
					argument:  3,
				},
				instruction{
					operation: jump,
					argument:  -3,
				},
				instruction{
					operation: accumulate,
					argument:  -99,
				},
				instruction{
					operation: accumulate,
					argument:  1,
				},
				instruction{
					operation: jump,
					argument:  -4,
				},
				instruction{
					operation: accumulate,
					argument:  6,
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
