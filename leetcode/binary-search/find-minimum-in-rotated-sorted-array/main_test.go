package find_minimum_in_rotated_sorted_array

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{nums: []int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "2",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}},
			want: 0,
		},
		{
			name: "3",
			args: args{nums: []int{11, 13, 15, 17}},
			want: 11,
		},
		{
			name: "4",
			args: args{nums: []int{3, 1, 2}},
			want: 1,
		},
		{
			name: "array of 2 elements",
			args: args{nums: []int{2, 1}},
			want: 1,
		},
		{
			name: "array of 3 elements",
			args: args{nums: []int{2, 3, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
