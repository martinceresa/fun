package main

import (
	"fmt"
	"math/big"
)

/* * Problem: find bigiest sum digit from number a^b where a,b < 10 */

func sum_digit(number *big.Int) (sumd *big.Int) {
	sumd = big.NewInt(int64(0))
	temp := new(big.Int)

	// Runify string
	numStr := number.String()

	// for each rune
	for _, v := range numStr {
		// Big Int rune first
		temp.SetString(string(v), 10)
		// Add it to sumd.
		sumd.Add(sumd, temp)
	}

	return
}

func main(){
	fmt.Println("Beginning loop")

	bint := new(big.Int)
	bmax := big.NewInt(0)

	for a := 1; a < 100; a ++ {
		for b := 1; b < 100; b ++ {
			bint.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			bcurr := sum_digit(bint)
			if bmax.Cmp(bcurr) < 0 {
				bmax.Set(bcurr)
			}
		}
	}

	fmt.Println("Bigest dig sum is: ", bmax)

	return
}
