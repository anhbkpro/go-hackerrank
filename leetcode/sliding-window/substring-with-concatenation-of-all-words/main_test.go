package substring_with_concatenation_of_all_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	assert.Equal(t, []int{13}, FindSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
	assert.Equal(t, []int{8}, FindSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	assert.Equal(t, []int{0, 9}, FindSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	assert.Equal(t, []int{}, FindSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	assert.Equal(t, []int{6, 9, 12}, FindSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
}
