package common

import (
	"fmt"
	"os"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

type InputFile struct {
	content string
}

func ReadInput(filename string) (*InputFile, error) {
	content, err := os.ReadFile(filename)
	CheckError(err)

	return &InputFile{content: string(content)}, nil
}

func (i *InputFile) Read() string {
	return i.content
}

func (i *InputFile) ReadLines(yield func(string) bool) {
	var line string

	for _, letter := range i.content {
		if letter == '\n' {
			if line == "" {
				continue
			}

			yield(line)

			line = ""

			continue
		}

		line += string(letter)
	}

	if line == "" {
		return
	}

	yield(line)
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
