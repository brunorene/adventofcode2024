package common

import (
	"fmt"
	"os"
)

type InputFile struct {
	content string
}

func ReadInput(filename string) (*InputFile, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	return &InputFile{content: string(content)}, nil
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
