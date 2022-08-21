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
 * Complete the 'larrysArray' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY A as parameter.
 */

func larrysArray(A []int32) string {
	numberPositionMap := make(map[int32]int32)
	for i := int32(0); i < int32(len(A)); i++ {
		numberPositionMap[A[i]] = i
	}

	fmt.Println(numberPositionMap)

	var index = int32(0)
	for index < int32(len(A)) {
		if A[index] == index+1 {
			index++
			continue
		}
		nextMinPosition := numberPositionMap[index+1]
		if nextMinPosition-index >= 2 {
			temp := A[nextMinPosition-2]
			A[nextMinPosition-2] = A[nextMinPosition-1]
			A[nextMinPosition-1] = A[nextMinPosition]
			A[nextMinPosition] = temp

			numberPositionMap[A[nextMinPosition-2]] = nextMinPosition - 2
			numberPositionMap[A[nextMinPosition-1]] = nextMinPosition - 1
			numberPositionMap[A[nextMinPosition]] = nextMinPosition
		} else if nextMinPosition-index == 1 {
			if nextMinPosition == int32(len(A)-1) {
				return "NO"
			}
			temp := A[nextMinPosition-1]
			A[nextMinPosition-1] = A[nextMinPosition]
			A[nextMinPosition] = A[nextMinPosition+1]
			A[nextMinPosition+1] = temp

			numberPositionMap[A[nextMinPosition-1]] = nextMinPosition - 1
			numberPositionMap[A[nextMinPosition]] = nextMinPosition
			numberPositionMap[A[nextMinPosition+1]] = nextMinPosition + 1
		}
	}
	if index == int32(len(A)) {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var A []int32

		for i := 0; i < int(n); i++ {
			AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			A = append(A, AItem)
		}

		result := larrysArray(A)

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
