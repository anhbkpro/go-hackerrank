package main

import (
	"fmt"
)

/*
 * Complete the 'CountApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	countFruits(s, t, a, apples)
	countFruits(s, t, b, oranges)
}

func countFruits(s int32, t int32, a int32, fruits []int32) {
	var fruitsOnHouse = 0
	for _, ac := range fruits {
		ll := ac + a
		if ll >= s && ll <= t {
			fruitsOnHouse++
		}
	}
	fmt.Println(fruitsOnHouse)
}
