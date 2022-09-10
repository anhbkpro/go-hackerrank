package halloween_sale

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'HowManyGames' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER d
 *  3. INTEGER m
 *  4. INTEGER s
 */

func HowManyGames(p int32, d int32, m int32, s int32) int32 {
	budget := int32(0)
	gameUnit := int32(0)
	for budget <= s {
		cost := p - d*gameUnit
		if cost < m {
			cost = m
		}
		budget += cost
		gameUnit++
	}

	return gameUnit - 1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	pTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	p := int32(pTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	m := int32(mTemp)

	sTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	s := int32(sTemp)

	answer := HowManyGames(p, d, m, s)

	fmt.Fprintf(writer, "%d\n", answer)

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
