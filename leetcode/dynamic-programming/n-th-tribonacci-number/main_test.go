package n_th_tribonacci_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTribonacci(t *testing.T) {
	assert.Equal(t, 4, Tribonacci(4))
}
