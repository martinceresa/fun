package main

import (
	"fmt"
	"math/big"
)

type Rational struct {
	enum *big.Int
	denum *big.Int
}

func upTo(n int) (res []Rational) {
	res = make([]Rational, n)
	res[0] = Rational{enum : big.NewInt(1), denum: big.NewInt(2)}
	for i:= 1; i < n ; i ++ {
		//
		enum := new(big.Int)
		enum.Set(res[i-1].denum)
		//
		denum := new(big.Int)
		denum.Mul(res[i-1].denum, big.NewInt(2))
		denum.Add(denum, res[i-1].enum)

		res[i] = Rational{enum: enum, denum: denum}
	}

	return
}

func SqRootTwoChain(n int, pn []Rational) (sqroot []Rational){
	sqroot = make([]Rational, len(pn))
	for p , v := range pn {
		//
		enum := new(big.Int)
		enum.Set(v.enum)
		enum.Add(enum, v.denum)

		denum := new(big.Int)
		denum.Set(v.denum)
		//
		sqroot[p] = Rational{ enum: enum, denum: denum}
	}

	return
}

func main() {
	fmt.Println("Solving P57")

	uptoPNs := upTo(1000)
	sqrootsFracs := SqRootTwoChain(1000, uptoPNs)

	candidates := 0

	for _ , p := range sqrootsFracs {
		enum := p.enum.String()
		denum := p.denum.String()
		fmt.Printf("%s / %s\n",enum, denum)
		if len(enum) > len(denum) {candidates ++}
	}

	fmt.Println("Condidates " , candidates)

	return
}
