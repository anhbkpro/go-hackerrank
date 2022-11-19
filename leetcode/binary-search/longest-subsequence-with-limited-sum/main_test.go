package longest_subsequence_with_limited_sum

import (
	"reflect"
	"testing"
)

func Test_answerQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{nums: []int{4, 5, 2, 1}, queries: []int{3, 10, 21}},
			want: []int{2, 3, 4},
		},
		{
			name: "second",
			args: args{nums: []int{2, 3, 4, 5}, queries: []int{1}},
			want: []int{0},
		},
		{
			name: "third",
			args: args{nums: []int{736411, 184882, 914641, 37925, 214915}, queries: []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}},
			want: []int{2, 2, 1, 1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
