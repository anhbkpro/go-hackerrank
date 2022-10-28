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
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */

func beautifulDays(i int, j int, k int) int {
	var beautifulDays int
	for s := i; s <= j; s++ {
		if (s-reverse(s))%k == 0 {
			beautifulDays++
		}
	}
	return beautifulDays
}

func reverse(input int) int {
	str := strconv.Itoa(input)
	rs := reverseStr(str)
	i, err := strconv.Atoi(rs)
	if err != nil {
		return 0
	}
	return i
}

func reverseStr(str string) string {
	runes := []rune(str)
	if len(runes) == 1 {
		return string(runes[0])
	}
	return str[len(runes)-1:] + reverseStr(string(runes[:len(runes)-1]))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	iTemp, err := strconv.Atoi(firstMultipleInput[0])
	checkError(err)
	i := iTemp

	jTemp, err := strconv.Atoi(firstMultipleInput[1])
	checkError(err)
	j := jTemp

	kTemp, err := strconv.Atoi(firstMultipleInput[2])
	checkError(err)
	k := kTemp

	result := beautifulDays(i, j, k)

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
