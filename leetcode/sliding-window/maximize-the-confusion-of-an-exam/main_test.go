package maximize_the_confusion_of_an_exam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxConsecutiveAnswers(t *testing.T) {
	assert.Equal(t, 5, MaxConsecutiveAnswers("TTFTTFTT", 1))
	assert.Equal(t, 4, MaxConsecutiveAnswers("TTFF", 2))
	assert.Equal(t, 3, MaxConsecutiveAnswers("TFFT", 1))
}
