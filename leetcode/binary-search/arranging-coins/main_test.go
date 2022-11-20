package arranging_coins

import "testing"

func Test_arrangeCoins(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 row",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "2 rows",
			args: args{n: 5},
			want: 2,
		},
		{
			name: "3 rows",
			args: args{n: 8},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
