package missing_number

import "testing"

func TestMissingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4",
			args: args{
				[]int{3, 0, 1},
			},
			want: 2,
		},
		{
			name: "10",
			args: args{
				[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNumber(tt.args.nums); got != tt.want {
				t.Errorf("MissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
