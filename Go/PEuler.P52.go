package main

import (
	"fmt"
	// "log"
	// "strings"
	"math"
)

func numToPrimeStrCod( num uint64 ) (cod uint64){
	str_num := fmt.Sprintf("%d", num)
	cod = 1

	for _, r := range str_num {
		switch r {
		case '0':
			cod *= 2
		case '1':
			cod *= 3
		case '2':
			cod *= 5
		case '3':
			cod *= 7
		case '4':
			cod *= 11
		case '5':
			cod *= 13
		case '6':
			cod *= 17
		case '7':
			cod *= 19
		case '8':
			cod *= 23
		case '9':
			cod *= 29
		}
	}

	return
}

// Mult property + 1
const multProp int = 7

func main(){

	var candidate uint64

	var ccand uint64
	var cnext uint64
	var i int

	// fmt.Println("Checking cod 125874 and 251748", numToPrimeStrCod(125874),numToPrimeStrCod(251748))

	for candidate = 1; candidate < math.MaxUint64; candidate++ {
		// log.Println("Candidate", candidate)
		ccand = numToPrimeStrCod(candidate)
		for i = 2; i < multProp ; i ++ {
			cnext = numToPrimeStrCod(uint64(i) * candidate)
			if cnext != ccand {
				break
			}
		}

		if i == multProp {
			break
		}
	}

	fmt.Println("We found something", candidate)
}
