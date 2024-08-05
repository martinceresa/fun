package math

import ( "math/big" )

/********************************/
// Factorials
func createFactorialMem()(func (int) (int64)) {
	var mem []int64

	var factorial = func (a int) (fact_a int64) {
		len_mem := len(mem)

		// If |fact(a)| is not in memory (|mem|)
		if a > len_mem {
			// Compute intermediate steps
			for i := len_mem; i < a; i ++ {
				mem = append(mem, mem[i-1] * int64(i))
			}
		}
		// Return |fact(a) = mem[a-1]|
		fact_a = mem[a-1]
		return
	}

	return factorial
}

// return an slice |factorial[i] = fact(i)|
func ComputeFactorials(i int)(factorial []*big.Int){
	factorial = make([]*big.Int, i+1)
	factorial[0] = big.NewInt(1)

	for j := 1; j <= i; j ++ {
		factorial[j] = big.NewInt(0)
		factorial[j].Mul(factorial[j-1], big.NewInt(int64(j)))
	}

	return
}
/********************************/

/********************************/
// Super naive approach to  combinatorics
func Naive_nCk(n , k int, factorials []*big.Int) (res *big.Int) {
	// Decs and Inits
	temp := big.NewInt(0)
	res = big.NewInt(0)
	//
	temp.Mul(factorials[k], factorials[n-k])
	res.Div(factorials[n], temp)

	return
}

// If we expand factorials to be a matrix maybe we can just compute
// C(n,k) = n * (n-1) * ... * (n-k+1) / (n-k) * .. * 2
// C(n,k) = M[k,n] = M[(n-k+1), n-k] (or something similar)
