package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputFile_ReadLines(t *testing.T) {
	type fields struct {
		filename string
		content  string
	}

	tests := []struct {
		name    string
		fields  fields
		results []string
	}{
		{
			name:    "ok",
			fields:  fields{content: "test1\ntest2\ntest3"},
			results: []string{"test1", "test2", "test3"},
		},
		{
			name:    "ok with empty lines",
			fields:  fields{content: "test1\ntest2\n\ntest3\n"},
			results: []string{"test1", "test2", "test3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InputFile{content: tt.fields.content}

			for line := range i.ReadLines {
				assert.Equal(t, tt.results[0], line)

				tt.results = tt.results[1:]
			}

			assert.Empty(t, tt.results)
		})
	}
}
