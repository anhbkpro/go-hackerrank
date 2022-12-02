package add_two_numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	// test same length inputs
	l13 := ListNode{3, nil}
	l12 := ListNode{4, &l13}
	l11 := ListNode{2, &l12}
	l23 := ListNode{4, nil}
	l22 := ListNode{6, &l23}
	l21 := ListNode{5, &l22}
	r23 := ListNode{8, nil}
	r22 := ListNode{0, &r23}
	r21 := ListNode{7, &r22}

	// test different length inputs
	t2l12 := ListNode{9, nil}
	t2l11 := ListNode{9, &t2l12}

	t2l23 := ListNode{9, nil}
	t2l22 := ListNode{9, &t2l23}
	t2l21 := ListNode{9, &t2l22}

	t2r4 := ListNode{1, nil}
	t2r3 := ListNode{0, &t2r4}
	t2r2 := ListNode{9, &t2r3}
	t2r1 := ListNode{8, &t2r2}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "same length",
			args: args{
				l1: &l11,
				l2: &l21,
			},
			want: &r21,
		},
		{
			name: "different length",
			args: args{
				l1: &t2l11,
				l2: &t2l21,
			},
			want: &t2r1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
