package check_if_n_and_its_double_exist

import "testing"

func Test_checkIfExist(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{arr: []int{10, 2, 5, 3}},
			want: true,
		},
		{
			name: "second",
			args: args{arr: []int{3, 1, 7, 11}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.args.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
