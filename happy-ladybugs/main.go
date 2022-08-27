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
 * Complete the 'happyLadybugs' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING b as parameter.
 */

func happyLadybugs(b string) string {
	fmt.Println("\nCheck happy lady bugs: ", b)
	bugMap := make(map[rune]int32)
	for _, c := range b {
		if _, ok := bugMap[c]; ok {
			bugMap[c]++
		} else {
			bugMap[c] = 1
		}
	}
	fmt.Println("bug map:", bugMap)

	underscoreRune := int32(95)
	_, hasEmptyCell := bugMap[underscoreRune]
	if !canMakePair(bugMap) {
		return "NO"
	}
	if hasEmptyCell {
		return "YES"
	}
	return "NO"
}

func canMakePair(m map[rune]int32) bool {
	for k, v := range m {
		if v < 2 && k != int32(95) {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		fmt.Sprintf("%v", n)
		b := readLine(reader)

		result := happyLadybugs(b)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
