package sqrtx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "even square",
			args: args{x: 4},
			want: 2,
		},
		{
			name: "odd square",
			args: args{x: 11},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MySqrt(tt.args.x), "MySqrt(%v)", tt.args.x)
		})
	}
}
