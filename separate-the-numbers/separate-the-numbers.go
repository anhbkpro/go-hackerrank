package separate_the_numbers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'SeparateNumbers' function below.
 *
 * The function accepts STRING s as parameter.
 */

func SeparateNumbers(s string) {
	for i := 1; i < len(s)/2+1; i++ {
		subString := s[:i]
		if sequential(s, subString) {
			fmt.Println("YES " + subString)
			return
		}
	}
	fmt.Println("NO")
}

func sequential(s, subString string) bool {
	if len(s) == 0 {
		return true
	}

	if strings.HasPrefix(s, subString) {
		l := len(subString)
		i, _ := strconv.Atoi(subString)
		return sequential(s[l:], fmt.Sprint(i+1))
	}

	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		SeparateNumbers(s)
	}
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
