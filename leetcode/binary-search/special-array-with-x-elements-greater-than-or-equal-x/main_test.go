package special_array_with_x_elements_greater_than_or_equal_x

import "testing"

func Test_specialArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{nums: []int{3, 5}},
			want: 2,
		},
		{
			name: "second",
			args: args{nums: []int{0, 0}},
			want: -1,
		},
		{
			name: "third",
			args: args{nums: []int{0, 4, 3, 0, 4}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialArrayBs(tt.args.nums); got != tt.want {
				t.Errorf("specialArrayBs() = %v, want %v", got, tt.want)
			}
		})
	}
}
