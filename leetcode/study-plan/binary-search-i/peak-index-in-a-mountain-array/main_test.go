package peak_index_in_a_mountain_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{arr: []int{0, 2, 1, 0}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PeakIndexInMountainArray(tt.args.arr), "PeakIndexInMountainArray(%v)", tt.args.arr)
		})
	}
}
