package single_row_keyboard

import "testing"

func Test_calculateTime(t *testing.T) {
	type args struct {
		keyboard string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				keyboard: "abcdefghijklmnopqrstuvwxyz",
				word:     "cba",
			},
			want: 4,
		},
		{
			name: "second",
			args: args{
				keyboard: "pqrstuvwxyzabcdefghijklmno",
				word:     "leetcode",
			},
			want: 73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTime(tt.args.keyboard, tt.args.word); got != tt.want {
				t.Errorf("calculateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
