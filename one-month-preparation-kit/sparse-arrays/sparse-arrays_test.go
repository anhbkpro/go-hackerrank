package sparse_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchingStrings(t *testing.T) {
	strings := []string{"aba", "baba", "aba", "xzxb"}
	queries := []string{"aba", "xzxb", "ab"}
	got := MatchingStrings(strings, queries)
	assert.Equal(t, []int32{2, 1, 0}, got)
}
