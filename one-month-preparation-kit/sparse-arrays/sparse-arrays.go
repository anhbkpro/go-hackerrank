package sparse_arrays

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'MatchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */

func MatchingStrings(strings []string, queries []string) []int32 {
	m := make(map[string]int32)
	for _, str := range strings {
		if _, ok := m[str]; ok {
			m[str]++
		} else {
			m[str] = 1
		}
	}

	var result []int32
	for _, query := range queries {
		if v, ok := m[query]; ok {
			result = append(result, v)
		} else {
			result = append(result, 0)
		}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stringsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var strs []string

	for i := 0; i < int(stringsCount); i++ {
		stringsItem := readLine(reader)
		strs = append(strs, stringsItem)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := MatchingStrings(strs, queries)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
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
