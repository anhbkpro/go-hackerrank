package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, IsPalindrome(121))
	assert.Equal(t, false, IsPalindrome(-121))
	assert.Equal(t, false, IsPalindrome(20))
}
