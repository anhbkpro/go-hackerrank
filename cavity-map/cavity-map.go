package cavity_map

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'CavityMap' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func CavityMap(grid []string) []string {
	var modifiedGrid []string
	for i, s := range grid {
		if i == 0 || i == len(grid)-1 {
			modifiedGrid = append(modifiedGrid, s)
		} else {
			var str []rune
			for j, v := range s {
				if j == 0 || j == len(s)-1 {
					str = append(str, v)
				} else {
					leftValue, _ := strconv.Atoi(string(s[j-1]))
					rightValue, _ := strconv.Atoi(string(s[j+1]))
					topValue, _ := strconv.Atoi(string(grid[i-1][j]))
					bottomValue, _ := strconv.Atoi(string(grid[i+1][j]))
					cellValue, _ := strconv.Atoi(string(s[j]))
					if cellValue > leftValue && cellValue > rightValue && cellValue > topValue && cellValue > bottomValue {
						str = append(str, 88) // rune of X is 88
					} else {
						str = append(str, v)
					}
				}
			}
			modifiedGrid = append(modifiedGrid, string(str))
		}
	}

	return modifiedGrid
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := CavityMap(grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
