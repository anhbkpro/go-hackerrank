package repeated_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedString(t *testing.T) {
	input := "aba"
	got := RepeatedString(input, 10)
	assert.Equal(t, int64(7), got)

	input2 := "a"
	got2 := RepeatedString(input2, 1000000000000)
	assert.Equal(t, int64(1000000000000), got2)

	input3 := "abcac"
	got3 := RepeatedString(input3, 10)
	assert.Equal(t, int64(4), got3)

	input4 := "x"
	got4 := RepeatedString(input4, 970770)
	assert.Equal(t, int64(0), got4)

	input5 := "epsxyyflvrrrxzvnoenvpegvuonodjoxfwdmcvwctmekpsnamchznsoxaklzjgrqruyzavshfbmuhdwwmpbkwcuomqhiyvuztwvq"
	got5 := RepeatedString(input5, 549382313570)
	assert.Equal(t, int64(16481469408), got5)
}
