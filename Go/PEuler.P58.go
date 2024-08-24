package main

import (
	"fmt"
	sieve "gosolution/Sieve"
)

func compn(n int) (res []int) {
	res = make([]int, n)

	s, it  := 2 , 0

	res[0] = 1

	for i := 1 ; i < n ; i++ {
		if it == 4 {
			s += 2
			it = 0
		}
		res[i] = res[i - 1] + s
		it ++
	}

	return
}

func totalPrimes(primes []bool, n int) (pr int, s int, total int) {
	it , l := 0 , 1
	pr , s = 0 , 2

	for total = 1 ; total <= n ; total++ {
		if it == 4 {
			s += 2
			it = 0
		}
		l = l + s
		it ++
		if l >= len(primes) {
			break
		}
		if primes[l] {
			pr ++
		}
	}

	return
}

const EN = 100000

func main() {

	fmt.Println("Checking Initial suquence numbers.")

	// test := compn(15)

	primes := sieve.PrimesUpToSquare(EN)

	for n := 100; n < len(primes); n ++ {
		pms , s, total := totalPrimes(primes, n)
		var ratio float64 = float64(pms) / float64(total)
		if ratio < 0.1 {
			fmt.Printf("Ratio [%d]/[%d]= [%f] -- Side [%d]",pms,total, ratio, s)
			break
		}
	}


	return
}
