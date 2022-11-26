package cutting_ribbons

import "testing"

func Test_maxLength(t *testing.T) {
	type args struct {
		ribbons []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{ribbons: []int{9, 7, 5}, k: 3},
			want: 5,
		},
		{
			name: "100000",
			args: args{ribbons: []int{100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 1, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000}, k: 49},
			want: 100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.args.ribbons, tt.args.k); got != tt.want {
				t.Errorf("maxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
