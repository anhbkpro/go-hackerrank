package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	l := len(s)
	rows := math.Sqrt(float64(l))
	f := math.Floor(rows)
	c := math.Ceil(rows)
	columns := rows
	if f != c {
		columns = rows + 1
	}
	rowsInt := int(rows)
	columnsInt := int(columns)

	if rowsInt*columnsInt < l {
		rows++
	}
	columnMap := make(map[int]string)

	for i := 0; i < columnsInt; i++ {
		for j := 0; j < l; j++ {
			if j%columnsInt == i {
				columnMap[i] = columnMap[i] + string(s[j])
			}
		}
	}
	fmt.Println(columnMap)
	return valuesString(columnMap)
}

func valuesString(m map[int]string) string {
	values := make([]string, 0, len(m))
	for i := 0; i < len(m); i++ {
		values = append(values, m[i])
	}
	return strings.Join(values, " ")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

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
