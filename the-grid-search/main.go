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
 * Complete the 'gridSearch' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING_ARRAY G
 *  2. STRING_ARRAY P
 */

func gridSearch(G []string, P []string) string {
	for i, s := range G {
		for j, ss := range s {
			if string(ss) == string(P[0][0]) {
				r := isMatch(i, j, G, P)
				if r {
					return "YES"
				}
			}
		}
	}
	return "NO"
}

func isMatch(pi, pj int, G []string, P []string) bool {
	if len(G[0])-pj < len(P[0]) {
		return false
	}
	for i := pi; i < pi+len(P); i++ {
		for j := pj; j < pj+len(P[i-pi]); j++ {
			if i-pi < len(P) && j-pj < len(P[0]) && G[i][j] != P[i-pi][j-pj] {
				return false
			}
		}
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		RTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		R := int32(RTemp)

		//CTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		//checkError(err)
		//C := int32(CTemp)

		var G []string

		for i := 0; i < int(R); i++ {
			GItem := readLine(reader)
			G = append(G, GItem)
		}

		secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		rTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		r := int32(rTemp)

		//cTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
		//checkError(err)
		//c := int32(cTemp)

		var P []string

		for i := 0; i < int(r); i++ {
			PItem := readLine(reader)
			P = append(P, PItem)
		}

		result := gridSearch(G, P)

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
