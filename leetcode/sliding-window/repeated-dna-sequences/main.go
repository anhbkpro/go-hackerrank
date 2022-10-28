package repeated_dna_sequences

func FindRepeatedDnaSequences(s string) []string {
	seen := make(map[string]struct{})
	repeated := make(map[string]struct{})
	var ans []string
	for i := 0; i+10 <= len(s); i++ {
		ten := s[i : i+10]
		if _, ok := seen[ten]; ok {
			repeated[ten] = struct{}{}
		} else {
			seen[ten] = struct{}{}
		}
	}
	for k := range repeated {
		ans = append(ans, k)
	}
	return ans
}
