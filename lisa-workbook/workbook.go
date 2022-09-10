package lisa_workbook

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'Workbook' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY arr
 */

func Workbook(n int32, k int32, arr []int32) int32 {
	totalPage := int32(0)
	specialNumber := int32(0)
	m := map[int32][]int32{}
	for c, v := range arr {
		fmt.Println(fmt.Sprintf("========== [START] CHAPTER %d ========", c+1))
		fmt.Printf("Start page index of CHAPTER %d at [%d]\n", c+1, totalPage+1)
		pages := int32(math.Ceil(float64(v) / float64(k)))
		previousPage := totalPage
		totalPage += pages
		fmt.Printf("End page index of CHAPTER %d at [%d] \n", c+1, totalPage)

		var problems []int32
		for p := int32(0); p < v; p++ {
			problems = append(problems, p+1)
		}
		fmt.Println(fmt.Sprintf("CHAPTER %d has %d problems: %v", c+1, v, problems))
		for i := previousPage; i < totalPage; i++ {
			m[i+1] = problems[(i-previousPage)*k : int(math.Min(float64(v), float64((i-previousPage+1)*k)))]
			fmt.Println(fmt.Sprintf("*** Page %d (m[%d]) holds problems: %v", i+1, i+1, m[i+1]))
			for _, v1 := range m[i+1] {
				if i+1 == v1 {
					specialNumber++
					fmt.Println("+++ Meet special page", i+1, " => Total special pages", specialNumber)
				}
			}
		}
		fmt.Println(fmt.Sprintf("========== [END] CHAPTER %d ==========\n", c+1))
	}

	return specialNumber
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := Workbook(n, k, arr)

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
