package top_k_frequent_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	assert.Equal(t, []string{"i", "love"}, TopKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
