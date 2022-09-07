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
 * Complete the 'weightedUniformStrings' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER_ARRAY queries
 */

func weightedUniformStrings(s string, queries []int32) []string {
	runeMap := make(map[rune]int32)
	weightedMap := make(map[int32]struct{})
	var rv []string
	for _, r := range s {
		if _, ok := runeMap[r]; ok {
			runeMap[r]++
		} else {
			runeMap[r] = 1
		}
	}

	for k, v := range runeMap {
		for i := int32(0); i < v; i++ {
			fmt.Println(fmt.Sprintf("--- %d*%s=%d", i+1, string(k), (i+1)*(k-97+1)))
			weightedMap[(i+1)*(k-97+1)] = struct{}{}
		}
	}

	for _, query := range queries {
		if _, ok := weightedMap[query]; ok {
			rv = append(rv, "Yes")
		} else {
			rv = append(rv, "No")
		}
	}

	return rv
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	s := readLine(reader)

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := weightedUniformStrings(s, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
