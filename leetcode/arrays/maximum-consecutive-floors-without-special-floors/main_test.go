package maximum_consecutive_floors_without_special_floors

import "testing"

func Test_maxConsecutive(t *testing.T) {
	type args struct {
		bottom  int
		top     int
		special []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				bottom:  2,
				top:     9,
				special: []int{4, 6},
			},
			want: 3,
		},
		{
			name: "second",
			args: args{
				bottom:  6,
				top:     8,
				special: []int{7, 6, 8},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutive(tt.args.bottom, tt.args.top, tt.args.special); got != tt.want {
				t.Errorf("maxConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
