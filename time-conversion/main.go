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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	isAM := isAM(s)
	hour := hour(s)
	fmt.Println(isAM)
	if isAM {
		if hour == 12 {
			return "00" + s[2:len(s)-2]
		}
		return s[:len(s)-2]
	} else {
		if hour == 12 {
			return s[:len(s)-2]
		}
		return fmt.Sprintf("%d%s", hour+12, s[2:len(s)-2])
	}
}

func hour(s string) int32 {
	hour, err := strconv.Atoi(s[:2])
	if err != nil {
		return 0
	}
	return int32(hour)
}

func isAM(s string) bool {
	fmt.Println(s[len(s)-2:])
	return s[len(s)-2:] == "AM"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
