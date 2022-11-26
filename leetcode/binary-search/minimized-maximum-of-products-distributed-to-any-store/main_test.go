package minimized_maximum_of_products_distributed_to_any_store

import "testing"

func Test_minimizedMaximum(t *testing.T) {
	type args struct {
		n          int
		quantities []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3",
			args: args{n: 6, quantities: []int{11, 6}},
			want: 3,
		},
		{
			name: "5",
			args: args{n: 7, quantities: []int{15, 10, 10}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizedMaximum(tt.args.n, tt.args.quantities); got != tt.want {
				t.Errorf("minimizedMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
