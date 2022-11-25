package find_smallest_common_element_in_all_rows

import "testing"

func Test_smallestCommonElement(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{mat: [][]int{{1, 2, 3, 4, 5}, {2, 4, 5, 8, 10}, {3, 5, 7, 9, 11}, {1, 3, 5, 7, 9}}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestCommonElement(tt.args.mat); got != tt.want {
				t.Errorf("smallestCommonElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
