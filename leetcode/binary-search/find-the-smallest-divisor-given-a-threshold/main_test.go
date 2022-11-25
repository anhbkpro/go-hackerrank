package find_the_smallest_divisor_given_a_threshold

import "testing"

//Runtime: 47 ms, faster than 90.00% of Go online submissions for Find the Smallest Divisor Given a Threshold.
//Memory Usage: 6.9 MB, less than 100.00% of Go online submissions for Find the Smallest Divisor Given a Threshold.
func Test_smallestDivisor(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{nums: []int{1, 2, 5, 9}, threshold: 6},
			want: 5,
		},
		{
			name: "44",
			args: args{nums: []int{44, 22, 33, 11, 1}, threshold: 5},
			want: 44,
		},
		{
			name: "1",
			args: args{nums: []int{21212, 10101, 12121}, threshold: 1000000},
			want: 1,
		},
		{
			name: "3",
			args: args{nums: []int{2, 3, 5, 7, 11}, threshold: 11},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestDivisor(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("smallestDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}
