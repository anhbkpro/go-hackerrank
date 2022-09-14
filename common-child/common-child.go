package common_child

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'CommonChild' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func CommonChild(s1 string, s2 string) int32 {
	lsc := make([][]int32, len(s1)+1)
	for i := 0; i < len(lsc); i++ {
		lsc[i] = make([]int32, len(s2)+1)
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				lsc[i][j] = lsc[i-1][j-1] + 1
			} else {
				lsc[i][j] = int32(math.Max(float64(lsc[i-1][j]), float64(lsc[i][j-1])))
			}
		}
	}

	return lsc[len(s1)][len(s2)]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := CommonChild(s1, s2)

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
