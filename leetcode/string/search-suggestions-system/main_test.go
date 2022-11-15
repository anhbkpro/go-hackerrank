package search_suggestions_system

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	assert.Equal(t, [][]string{{}, {}, {}, {}, {}, {}, {}}, SuggestedProducts([]string{"havana"}, "tatiana"))
	assert.Equal(t,
		[][]string{
			{"mobile", "moneypot", "monitor"},
			{"mobile", "moneypot", "monitor"},
			{"mouse", "mousepad"},
			{"mouse", "mousepad"},
			{"mouse", "mousepad"},
		},
		SuggestedProducts(
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse"),
	)
	assert.Equal(t,
		[][]string{
			{"baggage", "bags", "banner"},
			{"baggage", "bags", "banner"},
			{"baggage", "bags"},
			{"bags"},
		}, SuggestedProducts(
			[]string{"bags", "baggage", "banner", "box", "cloths"},
			"bags"))
}
