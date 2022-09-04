package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'alternate' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func alternate(s string) int32 {
	charMap := make(map[rune]struct{})
	var distinctStr []rune
	for _, r := range s {
		if _, ok := charMap[r]; !ok {
			charMap[r] = struct{}{}
			fmt.Println(string(r))
			distinctStr = append(distinctStr, r)
		}
	}
	fmt.Println("charMap:", charMap)
	fmt.Println("distinctStr:", string(distinctStr))
	is := new(IteratedString)

	for i, c := range distinctStr {
		for j := i + 1; j < len(distinctStr); j++ {
			is.Pairs = append(is.Pairs, Pair{c, distinctStr[j]})
		}
	}

	fmt.Printf("Pairs: %v\n", is.Pairs)
	fmt.Printf("Total pairs: %v\n", len(is.Pairs)) // n*(n-1)/2
	longestLength := int32(0)
	for _, pair := range is.Pairs {
		var resultingStr []rune
		fmt.Printf("\n- iterating {%v, %v}\n", string(pair.First), string(pair.Second))
		for _, r := range s {
			if r == pair.First || r == pair.Second {
				resultingStr = append(resultingStr, r)
			}
		}
		fmt.Println("+++ resultingStr", string(resultingStr), ", length =", len(resultingStr))
		if isAlternatingStr(resultingStr) && longestLength < int32(len(resultingStr)) {
			longestLength = int32(len(resultingStr))
			fmt.Println("=> longestLength", longestLength)
		}
	}
	return longestLength
}

func isAlternatingStr(s []rune) bool {
	if len(s) < 2 {
		return true
	}
	first := s[0]
	second := s[1]
	for i := 0; i < len(s); i++ {
		if i%2 == 0 && s[i] != first {
			return false
		}
		if i%2 == 1 && s[i] != second {
			return false
		}
	}
	return true
}

type IteratedString struct {
	Pairs []Pair
}

type Pair struct {
	First  rune
	Second rune
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	lTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	l := int32(lTemp)

	fmt.Sprintf("Length %d", l)

	s := readLine(reader)

	result := alternate(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
