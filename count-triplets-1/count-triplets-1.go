package count_triplets_1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// CountTriplets Complete the CountTriplets function below.
func CountTriplets(arr []int64, r int64) int64 {
	result := int64(0)
	r2 := make(map[int64]int64)
	r3 := make(map[int64]int64)

	for _, v := range arr {
		if _, ok := r3[v]; ok {
			result += r3[v]
		}
		if _, ok := r2[v]; ok {
			r3[v*r] += r2[v]
		}
		r2[v*r] += 1
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := CountTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
