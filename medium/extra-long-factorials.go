package main

import (
	"fmt"
	"math/big"
)

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func extraLongFactorials(n int32) {
	var fact big.Int
	fact.MulRange(1, int64(n))
	fmt.Println(&fact)
}

func main() {
	extraLongFactorials(20)
}
