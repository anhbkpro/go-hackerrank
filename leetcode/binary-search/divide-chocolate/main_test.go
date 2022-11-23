package divide_chocolate

import "testing"

func Test_maximizeSweetness(t *testing.T) {
	type args struct {
		sweetness []int
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5 cuts",
			args: args{sweetness: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, k: 5},
			want: 6,
		},
		{
			name: "8 cuts",
			args: args{sweetness: []int{5, 6, 7, 8, 9, 1, 2, 3, 4}, k: 8},
			want: 1,
		},
		{
			name: "2 cuts",
			args: args{sweetness: []int{1, 2, 2, 1, 2, 2, 1, 2, 2}, k: 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeSweetness(tt.args.sweetness, tt.args.k); got != tt.want {
				t.Errorf("maximizeSweetness() = %v, want %v", got, tt.want)
			}
		})
	}
}
