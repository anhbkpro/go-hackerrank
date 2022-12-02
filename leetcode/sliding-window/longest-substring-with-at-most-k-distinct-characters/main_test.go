package longest_substring_with_at_most_k_distinct_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestKSubstr(t *testing.T) {
	assert.Equal(t, 7, LongestKSubstrV2("cbebebe", 3))
	assert.Equal(t, 5, LongestKSubstrV2("ababcbcca", 2))
	assert.Equal(t, 1, LongestKSubstrV2("abcde", 1))
}

func TestLongestKSubstr1(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2",
			args: args{
				s: "aabbcc",
				k: 1,
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				s: "aabbcc",
				k: 2,
			},
			want: 4,
		},
		{
			name: "6",
			args: args{
				s: "aabbcc",
				k: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LongestKSubstr(tt.args.s, tt.args.k), "LongestKSubstr(%v, %v)", tt.args.s, tt.args.k)
		})
	}
}
