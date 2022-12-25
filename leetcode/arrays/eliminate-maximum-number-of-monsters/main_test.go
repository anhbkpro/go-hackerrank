package eliminate_maximum_number_of_monsters

import "testing"

func Test_eliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				dist:  []int{1, 3, 4},
				speed: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			name: "second",
			args: args{
				dist:  []int{1, 1, 2, 3},
				speed: []int{1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "third",
			args: args{
				dist:  []int{3, 2, 4},
				speed: []int{5, 3, 2},
			},
			want: 1,
		},
		{
			name: "fourth",
			args: args{
				dist:  []int{3, 5, 7, 4, 5},
				speed: []int{2, 3, 6, 3, 2},
			},
			want: 2,
		},
		{
			name: "fifth",
			args: args{
				dist:  []int{4, 2, 8},
				speed: []int{2, 1, 4},
			},
			want: 2,
		},
		{
			name: "sixth",
			args: args{
				dist:  []int{4, 3, 4},
				speed: []int{1, 1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eliminateMaximum(tt.args.dist, tt.args.speed); got != tt.want {
				t.Errorf("eliminateMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
