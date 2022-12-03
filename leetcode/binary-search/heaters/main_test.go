package heaters

import "testing"

func Test_findRadius(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				houses:  []int{1, 2, 3},
				heaters: []int{2},
			},
			want: 1,
		},
		{
			name: "1",
			args: args{
				houses:  []int{1, 2, 3, 4},
				heaters: []int{1, 4},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				houses:  []int{1, 5},
				heaters: []int{2},
			},
			want: 3,
		},
		{
			name: "0",
			args: args{
				houses:  []int{1},
				heaters: []int{1, 2, 3, 4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
