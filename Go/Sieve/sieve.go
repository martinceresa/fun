package sieve

// Normal Sieve of Erathostenes.
// No wheel optimization

// Primes up to Square(|n|) returns prime numbers less than |n*n|
func PrimesUpToSquare(n int) []bool {
	var square_n = n*n
	var bmap []bool = make([]bool,square_n+1)

	// Initialize bmap
	bmap[0] , bmap[1] = false, false
	for p := 2; p < len(bmap); p++ {
		bmap[p] = true
	}

	// Compute sieve
	for a := 2 ; a <= n; a ++ {
		if bmap[a] {
			for b := a*a; b <= square_n; b = b+a {
				bmap[b] = false
			}
		}
	}

	// return bmap
	return bmap
}

// Duplicating limit of a sieve.
// Assuming |bmap| is a sieve where bmap[i] = true <=> i is prime.
func ExtendingSieve(bmap []bool) []bool{
	var n int = len(bmap)

	var square_n = n*n
	var newbmap []bool = make([]bool, square_n+1)

	// Init just in case
	newbmap[0], newbmap[1] = false, false
	// all new candidates to true
	for e := n; e < square_n; e++ {
		newbmap[e] = true
	}

	var lastPrime int
	for e , v := range bmap {
		if v {
			// e is prime
			newbmap[e] = v
			// normal sieve
			var lastE int
			for lastE = e*e; lastE < n; lastE = lastE + e {}
			for ne := lastE ; ne < square_n ; ne = ne + e {
				newbmap[ne] = false
			}

			lastPrime = e
		}
	}

	for ne := lastPrime+1; ne <= n ; ne ++ {
		if newbmap[ne] {
			for b := ne*ne; b <= square_n ; b = b+ne {
				newbmap[b] = false
			}
		}
	}

	return newbmap
}


// func main(){

// 	var a int = 100
// 	fmt.Println("Primes up to ", a*a)

// 	var bmap []bool = primesUpToSquare(a)

// 	bmap = extendingSieve(bmap)

// 	for e, v := range bmap {
// 		if v {
// 			fmt.Printf("%d ", e)
// 		}
// 	}
// }
