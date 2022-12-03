package longest_palindromic_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", LongestPalindrome("babad"))
	assert.Equal(t, "bb", LongestPalindrome("cbbd"))
}

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4",
			args: args{s: "bbbab"},
			want: 4,
		},
		{
			name: "2",
			args: args{s: "abbc"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestPalindromeSubseq(tt.args.s), "longestPalindromeSubseq(%v)", tt.args.s)
		})
	}
}
