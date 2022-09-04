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
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isBalanced(s string) string {
	var stack Stack
	for _, r := range s {
		if isOpeningBrace(r) {
			stack.Push(string(r))
		} else if isClosingBrace(r) {
			poppedOpeningBrace, ok := stack.Pop()
			if ok {
				if isNotAMatch(poppedOpeningBrace, string(r)) {
					return "NO"
				}
			} else {
				return "NO"
			}
		}
	}

	if len(stack) > 0 {
		return "NO"
	}

	return "YES"
}

type Stack []string

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isOpeningBrace(c rune) bool {
	for _, v := range "([{" {
		if v == c {
			return true
		}
	}

	return false
}

func isClosingBrace(c rune) bool {
	for _, v := range ")]}" {
		if v == c {
			return true
		}
	}

	return false
}

func isNotAMatch(openingBrace, closingBrace string) bool {
	m := map[string]string{"(": ")", "[": "]", "{": "}"}
	if v, ok := m[openingBrace]; ok {
		return v != closingBrace
	}

	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
