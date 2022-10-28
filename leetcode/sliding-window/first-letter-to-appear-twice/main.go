package first_letter_to_appear_twice

func RepeatedCharacter(s string) byte {
	freqChar := make(map[byte]struct{})
	for _, b := range []byte(s) {
		if _, ok := freqChar[b]; ok {
			return b
		} else {
			freqChar[b] = struct{}{}
		}
	}
	return 0
}
