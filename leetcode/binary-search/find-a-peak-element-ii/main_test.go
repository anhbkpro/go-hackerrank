package find_a_peak_element_ii

import (
	"reflect"
	"testing"
)

func TestFindPeakGrid(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "2x2",
			args: args{mat: [][]int{{1, 4}, {3, 2}}},
			want: []int{1, 0},
		},
		{
			name: "3x3",
			args: args{mat: [][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}},
			want: []int{1, 1},
		},
		{
			name: "5x5",
			args: args{mat: [][]int{{47, 30, 35, 8, 25}, {6, 36, 19, 41, 40}, {24, 37, 13, 46, 5}, {3, 43, 15, 50, 19}, {6, 15, 7, 25, 18}}},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPeakGrid(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPeakGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
