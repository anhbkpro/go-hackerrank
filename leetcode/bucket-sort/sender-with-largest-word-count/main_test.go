package sender_with_largest_word_count

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestWordCount(t *testing.T) {
	assert.Equal(t, "abfedA", LargestWordCountDiscuss(
		[]string{"Hello", "Hello", "Hello", "Hello"},
		[]string{"abFedc", "abfEdc", "abfeDc", "abfedA"}))
	assert.Equal(t, "Alice", LargestWordCountDiscuss(
		[]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"},
		[]string{"Alice", "userTwo", "userThree", "Alice"}))
	assert.Equal(t, "Charlie", LargestWordCount(
		[]string{"How is leetcode for everyone", "Leetcode is useful for practice"},
		[]string{"Bob", "Charlie"}))
}
