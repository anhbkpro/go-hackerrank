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
 * Complete the 'hackerrankInString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func hackerrankInString(s string) string {
	fmt.Println("Checking: ", s)
	r := []rune(s)
	indexStart := 0
	count := 0
	for _, item := range "hackerrank" {
		fmt.Println("Iterating", string(item))
		for i := indexStart; i < len(s); i++ {
			indexStart++
			if r[i] == item {
				count++
				if len([]rune("hackerrank")) == count {
					return "YES"
				}
				fmt.Println("--- Matched: ", string(r[i]), indexStart)
				break
			} else {
				fmt.Println("Skipped: ", string(r[i]), indexStart)
			}
		}
	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := hackerrankInString(s)

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
