package find_peak_element

import "testing"

func TestFindPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single result",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 2,
		},
		{
			name: "multiple results",
			args: args{nums: []int{1, 2, 1, 3, 5, 6, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("FindPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
