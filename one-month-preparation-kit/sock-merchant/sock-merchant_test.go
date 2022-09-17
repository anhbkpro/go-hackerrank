package sock_merchant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSockMerchant(t *testing.T) {
	input := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	got := SockMerchant(9, input)
	assert.Equal(t, int32(3), got)
}
