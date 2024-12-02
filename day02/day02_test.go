package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reportsToCheck(t *testing.T) {
	type args struct {
		report []int
	}

	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{
			name: "ok",
			args: args{report: []int{1, 2, 3, 4, 5, 6}},
			wantResult: [][]int{
				{1, 2, 3, 4, 5, 6},
				{2, 3, 4, 5, 6},
				{1, 3, 4, 5, 6},
				{1, 2, 4, 5, 6},
				{1, 2, 3, 5, 6},
				{1, 2, 3, 4, 6},
				{1, 2, 3, 4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantResult, reportsToCheck(tt.args.report), "reportsToCheck(%v)", tt.args.report)
		})
	}
}
