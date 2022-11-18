package search_in_rotated_sorted_array

import "testing"

func Test_search(t *testing.T) {
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
			name: "exists",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			want: 4,
		},
		{
			name: "exists",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
			want: -1,
		},
		{
			name: "single-element array",
			args: args{nums: []int{1}, target: 1},
			want: 0,
		},
		{
			name: "2-elements array",
			args: args{nums: []int{1, 3}, target: 3},
			want: 1,
		},
		{
			name: "3-elements array",
			args: args{nums: []int{3, 5, 1}, target: 3},
			want: 0,
		},
		{
			name: "3-elements array",
			args: args{nums: []int{5, 1, 3}, target: 5},
			want: 0,
		},
		{
			name: "2-elements array",
			args: args{nums: []int{1, 3}, target: 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
