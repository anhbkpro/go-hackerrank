package longest_increasing_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4",
			args: args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLIS(tt.args.nums), "lengthOfLIS(%v)", tt.args.nums)
		})
	}
}
