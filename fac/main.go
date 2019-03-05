/*
Factorial 30.  We cannot use int or int64 for the result,
because fac(30) is very large.
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int64
	for n = 0; n <= 30; n++ {
		fmt.Printf("Fac(%d)=%d\n", n, fac(n))
	}
}

func fac(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	} else {
		return big.NewInt(n).Mul(big.NewInt(n), fac(n-1))
	}
}
