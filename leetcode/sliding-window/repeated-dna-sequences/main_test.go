package repeated_dna_sequences

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	assert.Equal(t, []string{"AAAAACCCCC", "CCCCCAAAAA"}, FindRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	assert.Equal(t, []string{"AAAAAAAAAA"}, FindRepeatedDnaSequences("AAAAAAAAAAAAA"))
	assert.Equal(t, []string{"AAAAAAAAAA"}, FindRepeatedDnaSequences("AAAAAAAAAAA"))
}
