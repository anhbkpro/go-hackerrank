package find_target_indices_after_sorting_array

import (
	"reflect"
	"testing"
)

func Test_targetIndices(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "3 matches",
			args: args{nums: []int{1, 2, 2, 2, 3}, target: 2},
			want: []int{1, 2, 3},
		},
		{
			name: "2 matches",
			args: args{nums: []int{1, 2, 5, 2, 3}, target: 2},
			want: []int{1, 2},
		},
		{
			name: "1 matches",
			args: args{nums: []int{1, 2, 5, 2, 3}, target: 3},
			want: []int{3},
		},
		{
			name: "1 matches",
			args: args{nums: []int{1, 2, 5, 2, 3}, target: 5},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := targetIndices(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("targetIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
