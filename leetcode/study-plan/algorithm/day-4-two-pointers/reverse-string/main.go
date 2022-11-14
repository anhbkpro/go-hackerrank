package reverse_string

import "fmt"

func ReverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		temp := s[r]
		s[r] = s[l]
		s[l] = temp
		l++
		r--
	}
	fmt.Println(s)
}
