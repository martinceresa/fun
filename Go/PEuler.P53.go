package main

import (
	"fmt"
	"log"
	math "gosolution/Math"
	"math/big"
)

/*
 * Main ideas
 * C(n,k) = C(n, n-k) --> We have repeated elements.
 * In a way, it is like Pascal's Triangle?/Pyramid?, elements should describe a bell.
 * C(n+1, k ) > C(n, k)
 * k \in 1,n/2,  k1 < k2 <= n/n => C(n,k1) < C(n,k2)
 *
 * Tl;dr: explore half of elements, and once one is bigger than 1kk the rest
 * (until the half) are also greater.
 *
 * First approach was more efficient, but I fucked it up setting threshold to 100000 (instead of 1kk)
 * if n % 2 == 0, then 2(n-k) - 1
 * if n % 2 == 1, then 2(n-k)
 * But same thing goes iterating through the array :shrug:
 */

var threshold *big.Int = big.NewInt(1000000)
const N int = 100

func main(){

	fmt.Println("Computing PEruler Problem 53")

	// Compute all factorials up to 100!
	// We may not need them all but I think it should be fast enough
	factorials := math.ComputeFactorials(100)
	// log.Println("Factorials done", factorials)

	var absdom [N+1][N+1]int
	for p , _ := range absdom {
		for q , _ := range absdom[p] {
			absdom[p][q] = 0
		}
	}

	// C(n,0) = C(n,n) = 1
	for n := 0; n <= N; n++ {
		absdom[n][0] = -1
		absdom[n][n] = -1
	}

	for n := 1 ; n <= N; n ++ {
		var k int
		// Only bserve to n/2 (see facts above.)
		var topK int

		if n % 2 == 0 {
			topK = n / 2
		} else {
			topK = (n / 2) + 1
		}

		for k = 1; k <= topK ; k ++ {
			// Compute combinatorials when needed.
			if absdom[n][k] == 0 {
				if math.Naive_nCk(n,k,factorials).Cmp(threshold) > 0 {
					// C(n,k) >= threshold
					absdom[n][k] = 1

					// n1 < n2 => C(n1, k) < C(n2, k)
					for j := n+1; j <= N; j ++ {
						absdom[j][k] = 1
					}
				} else {
					// C(n,k) <= threshold
					absdom[n][k] = -1
				}
			}
		}
		// C(n,k) = C(n, n - k)
		for j := n; j > topK ; j-- {
			absdom[n][j] = absdom[n][n - j]
		}
	}

	fmt.Println("Since I suck!")

	othercount := 0
	for n := 1; n <= N ; n ++ {
		for k := 1; k <= n; k ++ {
			if absdom[n][k] > 0 {
				othercount ++
			}

			if absdom[n][k] == 0 {
				log.Fatalln("Not computed value?", n , k)
			}
		}
	}

	fmt.Println("Abs Dom counted ", othercount)
	return
}
