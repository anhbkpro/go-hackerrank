package longest_common_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				text1: "ace",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "0",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestCommonSubsequence(tt.args.text1, tt.args.text2), "longestCommonSubsequence(%v, %v)", tt.args.text1, tt.args.text2)
		})
	}
}
