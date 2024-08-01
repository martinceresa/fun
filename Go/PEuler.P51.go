package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	sieve "gosolution/Sieve"
)

// Return 10**n
func powTen(n int) (res int){
	if n < 0 {
		res = 0
	} else {
		res = 1
		for i := 0; i < n ; i++ {
			res *= 10
		}
	}
	return
}


func repDigits (n int) [10]int {
	var res [10]int

	for n > 1 {
		d := n % 10
		res[d] ++
		n = n / 10
	}

	return res
}

func repDig(n int) (res bool) {

	res = false

	for _ , v := range repDigits(n) {
		if v > 1 {
			res = true
			break
		}
	}

	return res
}

func checkPrim(n int, famWanted int , primes []bool) (res bool) {
	digitsRep := repDigits(n)

	var minD = 9 - famWanted


	var famSize [10]int

	for e , v := range digitsRep {
		if e <= minD && v > 1 {
			var d int

			str_n := fmt.Sprintf("%d", n)

			// log.Println("Candidate ", str_n, e, v )
			famSize[e] = 1

			for d = e + 1 ; d < 10; d ++ {
				// Generate candidate i
				to_check := strings.ReplaceAll(str_n, fmt.Sprintf("%d", e), fmt.Sprintf("%d", d))
				// log.Println("To check", to_check)
				i , err := strconv.Atoi(to_check)
				if err != nil {
					log.Fatalln("Failed coming back from string")
				}
				//

				// Check if candidate is prime.
				if primes[i] {
					famSize[e]++
				}
			}
		}
	}

	// Computing result
	res = false

	for _ , v := range famSize {
		if v >= famWanted {
			res = true
			break
		}
	}

	return res
}

func main(){
	fmt.Println("Solution Project Euler Problem 51")

	var N int = 10000

	// Candidates:  Primes up to 1000000
	log.Println("Generating Primes")
	var primesMap []bool = sieve.PrimesUpToSquare(N)
	log.Println("DONE Generating Primes")

	// Force-brute algorithm
	var res = 0
	for e , v := range primesMap {
		if v { // e is prime.
			if repDig(e) && // discard elements without repeating digits.
				checkPrim(e, 8, primesMap) { // check if property holds.
				res = e
				// return first one found.
				break
			}
		}
	}

	fmt.Println("First one found ", res)
}
