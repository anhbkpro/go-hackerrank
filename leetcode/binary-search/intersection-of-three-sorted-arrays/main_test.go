package intersection_of_three_sorted_arrays

import (
	"reflect"
	"testing"
)

func Test_arraysIntersection(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		arr3 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "3 pointers",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{1, 2, 5, 7, 9},
				arr3: []int{1, 3, 4, 5, 8},
			},
			want: []int{1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraysIntersection3Pointers(tt.args.arr1, tt.args.arr2, tt.args.arr3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arraysIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
