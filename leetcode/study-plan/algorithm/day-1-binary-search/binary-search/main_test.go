package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single value array",
			args: args{nums: []int{1}, target: 1},
			want: 0,
		},
		{
			name: "target exists",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9},
			want: 4,
		},
		{
			name: "target does not exist",
			args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 7},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Search(tt.args.nums, tt.args.target), "Search(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
