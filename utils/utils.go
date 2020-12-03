package utils

import (
	"log"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

// ReadLinesFromTextFile reads a text file and converts the lines into a []string
func ReadLinesFromTextFile(path string, fileName string) []string {
	str := readTextFile(path, fileName)
	lines := strings.Split(str, "\r\n")

	return lines
}

func readTextFile(path string, fileName string) string {
	box := packr.New("fileBox", path)
	str, err := box.FindString(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return str
}
