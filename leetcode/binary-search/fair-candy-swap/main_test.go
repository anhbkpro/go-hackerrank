package fair_candy_swap

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	type args struct {
		aliceSizes []int
		bobSizes   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{aliceSizes: []int{1, 1}, bobSizes: []int{2, 2}},
			want: []int{1, 2},
		},
		{
			name: "2",
			args: args{aliceSizes: []int{1, 2}, bobSizes: []int{2, 3}},
			want: []int{1, 2},
		},
		{
			name: "3",
			args: args{aliceSizes: []int{2}, bobSizes: []int{1, 3}},
			want: []int{2, 3},
		},
		{
			name: "4",
			args: args{aliceSizes: []int{1, 2, 5}, bobSizes: []int{2, 4}},
			want: []int{5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.args.aliceSizes, tt.args.bobSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
