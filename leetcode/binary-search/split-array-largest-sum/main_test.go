package split_array_largest_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_splitArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{nums: []int{7, 2, 5, 10, 8}, k: 2},
			want: 18,
		},
		{
			name: "2",
			args: args{nums: []int{1, 2, 3, 4, 5}, k: 2},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, splitArray(tt.args.nums, tt.args.k), "splitArray(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
