package maximum_candies_allocated_to_k_children

import "testing"

func Test_maximumCandies(t *testing.T) {
	type args struct {
		candies []int
		k       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{candies: []int{5, 8, 6}, k: 3},
			want: 5,
		},
		{
			name: "5",
			args: args{candies: []int{5, 2}, k: 11},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCandies(tt.args.candies, tt.args.k); got != tt.want {
				t.Errorf("maximumCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
