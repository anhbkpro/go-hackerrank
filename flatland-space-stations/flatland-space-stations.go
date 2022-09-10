package flatland_space_stations

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// FlatlandSpaceStations Complete the FlatlandSpaceStations function below.
func FlatlandSpaceStations(n int32, c []int32) int32 {
	m := make(map[int32]struct{})
	for _, v := range c {
		m[v] = struct{}{}
	}

	var distances []int32
	for i := int32(0); i < n; i++ {
		if _, ok := m[i]; ok {
			distances = append(distances, 0)
		} else {
			nearestStation := int32(math.MaxInt32)
			for _, v := range c {
				if int32(math.Abs(float64(i-v))) < nearestStation {
					nearestStation = int32(math.Abs(float64(i - v)))
				}
			}
			distances = append(distances, nearestStation)
		}
	}

	maxValue := int32(0)
	for _, d := range distances {
		if maxValue < d {
			maxValue = d
		}
	}

	return maxValue
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(m); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := FlatlandSpaceStations(n, c)

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
