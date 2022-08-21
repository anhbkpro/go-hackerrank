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
 * [Medium]
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

func timeInWords(h int32, m int32) string {
	var hoursMap = map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
	}

	var minutesMap = map[int32]string{
		1:  "one minute",
		2:  "two minutes",
		3:  "three minutes",
		4:  "four minutes",
		5:  "five minutes",
		6:  "six minutes",
		7:  "seven minutes",
		8:  "eight minutes",
		9:  "nine minutes",
		10: "ten minutes",
		11: "eleven minutes",
		12: "twelve minutes",
		13: "thirteen minutes",
		14: "fourteen minutes",
		15: "quarter",
		16: "sixteen minutes",
		17: "seventeen minutes",
		18: "eighteen minutes",
		19: "nineteen minutes",
		20: "twenty minutes",
		21: "twenty one minutes",
		22: "twenty two minutes",
		23: "twenty three minutes",
		24: "twenty four minutes",
		25: "twenty five minutes",
		26: "twenty six minutes",
		27: "twenty seven minutes",
		28: "twenty eight minutes",
		29: "twenty nine minutes",
		30: "half",
	}

	if m == 0 {
		return fmt.Sprintf("%s o' clock", hoursMap[h])
	} else if m <= 30 && m >= 1 {
		return fmt.Sprintf("%s past %s", minutesMap[m], hoursMap[h])
	} else {
		return fmt.Sprintf("%s to %s", minutesMap[60-m], hoursMap[h+1])
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

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
