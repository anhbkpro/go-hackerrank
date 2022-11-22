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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraysIntersection(tt.args.arr1, tt.args.arr2, tt.args.arr3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arraysIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
