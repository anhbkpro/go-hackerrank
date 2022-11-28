package the_k_weakest_rows_in_a_matrix

import (
	"reflect"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{mat: [][]int{
				{1, 1, 0, 0, 0},  // 2 _ 0 = 2*len(mat[0]) + 0 = 2*5 + 0 = 10
				{1, 1, 1, 1, 0},  // 4	   = 4*10 + 1 = 41
				{1, 0, 0, 0, 0},  // 1 _ 2 soldiers*10 + index = 10 + 1 = 11
				{1, 1, 0, 0, 0},  // 2 _ 3 = 2*10 + 3 = 23
				{1, 1, 1, 1, 1}}, // 5
				k: 3},
			want: []int{2, 0, 3},
		},
		{
			name: "2",
			args: args{mat: [][]int{
				{1, 1},
				{1, 0},
				{1, 0},
				{1, 1},
				{0, 0},
				{1, 1},
			}, // 5
				k: 1},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kWeakestRowsPriorityQueue(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kWeakestRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
