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
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
	sMap := make(map[int32]rune)
	for i, v := range s {
		sMap[int32(i)] = v
	}

	diffIndex := 0
	var long string
	var short string
	if len(s) > len(t) {
		long = s
		short = t
	} else {
		long = t
		short = s
	}

	for i, v := range long {
		if i < len(short) && v == int32(short[i]) {
			continue
		} else {
			diffIndex = i
			break
		}
	}

	if diffIndex == 0 && int32(len(s)+len(t)) < k || s == t {
		return "Yes"
	}

	diff := k - int32(len(s)-diffIndex+len(t)-diffIndex)

	if diff == 0 || (diff > 0 && diff%2 == 0) {
		return "Yes"
	}

	remaining := k - int32(len(long)-len(short))
	if remaining > 0 && 2*int32(len(short)) < remaining {
		return "Yes"
	}

	return "No"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

	fmt.Fprintf(writer, "%s\n", result)

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
