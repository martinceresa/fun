package main

import (
	"fmt"
	"slices"
	"math/big"
	"log"
)

/* * Lychrel Numbers! */

func isPal(i *big.Int) (isPal bool) {
	strRep := i.String()
	lenstrRep := len(strRep)

	isPal = true
	for i := 0 ; i*2 <= lenstrRep && isPal ; i++ {
		isPal = strRep[i] == strRep[lenstrRep - 1 - i]
	}

	if isPal {
		log.Println("Pal ", strRep)
	}

	return
}

func lynch(i *big.Int) (isLych bool) {

	isLych = true

	revI, add := new(big.Int), new(big.Int)

	for it := 0 ; it < 50 && isLych ; it++ {
		// Reverse number
		strRep := i.String()

		var runes []rune = []rune(strRep)
		slices.Reverse(runes)
		// reading reverse
		_ , ok := revI.SetString(string(runes), 10)
		if ! ok {
			log.Fatalln("Error reading reverse string")
		}
		//

		// Adding both numbers
		add.Add(i , revI)
		log.Println("Adding" , i , revI, add)

		isLych = ! isPal(add)

		// next step
		i.Set(add)
	}

	return

}

func main(){
	fmt.Println("Brute FORCEEEEE")

	count := 0
	for i := 0 ; i <= 10000; i ++ {
		bi := big.NewInt(int64(i))
		if lynch(bi) {
			count ++
		}
	}

	example := big.NewInt(int64(349))
	lynch(example)

	fmt.Println("Counted " , count)

	return
}
