package minimum_number_of_days_to_make_m_bouquets

import "testing"

//Runtime: 125 ms, faster than 93.75% of Go online submissions for Minimum Number of Days to Make m Bouquets.
//Memory Usage: 8.9 MB, less than 87.50% of Go online submissions for Minimum Number of Days to Make m Bouquets.
func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3",
			args: args{bloomDay: []int{1, 10, 3, 10, 2}, m: 3, k: 1},
			want: 3,
		},
		{
			name: "-1",
			args: args{bloomDay: []int{1, 10, 3, 10, 2}, m: 3, k: 2},
			want: -1,
		},
		{
			name: "12",
			args: args{bloomDay: []int{7, 7, 7, 7, 12, 7, 7}, m: 2, k: 3},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
