package main

import (
	"fmt"
	"log"
	// "sort"
	"os"
	"bufio"
	"strings"
	// "strconv"
	//
	"slices"
	"cmp"
)

/******************************************/
// * Type definitions
//
// Suits definition
//
type suit int

const (
	Club suit = iota
	Diamonds
	Spades
	Hearts
)

type hand int
const (
	HCard hand = iota
	OnePair
	TwoPair
	ThreeK
	Straight
	Flush
	FullHouse
	FourK
	StraightFlush
	RoyalFlush

)

// Cargs
type Card struct {
	Num int
	Suit suit
}

// Hand def
const handSize int = 5
const cardNums int = 15
type Hand []Card
/******************************************/
//
// Sort Hand
type OrdHand Hand
func (a OrdHand) Len() int { return len(a)}
func (a OrdHand) Swap(i, j int){ a[i], a[j] = a[j], a[i]}
func (a OrdHand) Less(i, j int) (res bool) {
	cL, cR := a[i], a[j]

	res = false

	if cL.Num < cR.Num {
		res = true
	}
	// else if cL.Num == cR.Num {
	// 	res = cL.Suit < cL.Suit
	// }

	return
}
// sort.Sort(OrdHand(hand))
//
type Winner bool
const (
	PL Winner = false
	PR Winner = true
)
//

/**********************/
// Card operations
func cardSucc(cL , cR Card) bool {
	return cL.Num + 1 == cR.Num
}

func cardSameSuit(cL, cR Card) bool {
	return cL.Suit == cR.Suit
}


func checkReps(h Hand) (reps []int , rep bool) {
	rep = false
	reps = make([]int, cardNums)

	for i := 0; i < cardNums; i++ {
		reps[i] = 0
	}

	for i:= 0; i < handSize ; i++ {
		if h[i].Num > cardNums {
			log.Fatalln("Bad Card", h[i])
		}
		reps[h[i].Num] ++
		if reps[h[i].Num] > 1 {
			rep = true
		}
	}

	return
}

// Functions, we always assume hands are ordered.
func succCheck(h Hand, cond func(Card,Card) bool) bool {
	var i int = 0
	for i = 0; i < handSize - 1 &&
		cond(h[i], h[i+1]); i++ {}
	return i == handSize - 1
}
//
func straight(h Hand) bool {
	return succCheck(h, cardSucc)
}

func sameSuit(h Hand) bool {
	return succCheck(h, cardSameSuit)
}

// * Poker Hands definitions
//
func straightFlush(h Hand) bool {
	return straight(h) && sameSuit(h)
}

func royalFlush(h Hand) bool {
	return h[0].Num == 10 &&  straightFlush(h)
}


func checkFour(h Hand) (res bool, num int) {
	reps, rep := checkReps(h)

	res, num = false, 0

	if rep {
		res = false
		for i:= 0; i < cardNums && ! res; i++  {
			res, num = reps[i] == 4 , i
		}
	}

	return
}

func checkFullH(h Hand) (res bool, treenum int , pairnum int){
	reps, rep := checkReps(h)
	res2 := false
	res  = false
	pairnum , treenum = 0 , 0
	if rep {
		for i:= 0; i < cardNums ; i ++ {
			if reps[i] == 2 {
				res2, pairnum = true, i
			}
			if reps[i] == 3 {
				res, treenum = true, i
			}
		}
	}

	res = res2 && res
	return
}

func checkThreeK(h Hand) (res bool, num int) {
	reps, rep := checkReps(h)
	res , num = false, 0
	if rep {
		for i:= 0 ; i < cardNums && ! res; i ++ {
			res, num = reps[i] == 3, i
		}
	}
	return
}

func checkTwo(h Hand) (res bool, num int) {
	reps, rep := checkReps(h)
	res = false
	if rep {
		for i:= 0 ; i < cardNums && ! res; i ++ {
			res, num = reps[i] == 2, i
		}
	}
	return
}

func checkTTwo(h Hand) (res bool, numl , numr int) {
	reps, rep := checkReps(h)
	res, numl, numr = false, 0 , 0
	res2 := false
	if rep {
		for i:= 0 ; i < cardNums && ! res; i ++ {
			if !res2 && reps[i] == 2 {
				res2 = true
				numl = i
			} else if reps[i] == 2 {
				res = true
				numr = i
			}
		}
	}

	return
}

/**********************/

func bestHand(h Hand) (res hand, metadata []int) {
	// Sort hand
	// sort.Sort(OrdHand(h))

	metadata = make([]int, 2)

	// Lowest hand is just /High Card/
	res, metadata[0] = HCard, h[handSize-1].Num

	var isH bool = false

	if royalFlush(h) {
		log.Println("Royal Flush")
		res = RoyalFlush
	} else if straightFlush(h) {
		log.Println("Stright Flush")
		res = StraightFlush
	} else if isH, metadata[0] = checkFour(h); isH {
		log.Println("Four of a Kind", metadata)
		res = FourK
	} else if isH, metadata[0], metadata[1] = checkFullH(h); isH {
		log.Println("Full House", metadata)
		res = FullHouse
	} else if sameSuit(h) {
		log.Println("Same suit")
		res = Flush
	} else if straight(h) {
		log.Println("Straight")
		res = Straight
	} else if isH, metadata[0] = checkThreeK(h); isH {
		log.Println("Three of a Kind", metadata)
		res = ThreeK
	} else if isH, metadata[0], metadata[1] = checkTTwo(h); isH {
		log.Println("Two pairs", metadata)
		res = TwoPair
	} else if isH, metadata[0] = checkTwo(h) ; isH {
		log.Println("One pair of", metadata[0])
		res = OnePair
	}

	if res == HCard {
		metadata[0] = h[handSize-1].Num
		log.Println("Highest Card", metadata[0])
	}

	return
}

func whoWins( hL, hR Hand) (winner Winner) {

	lHand , lmetadata := bestHand(hL)
	rHand , rmetadata := bestHand(hR)

	if lHand > rHand {
		winner = PL
	} else if lHand < rHand {
		winner = PR
	} else {
		if lmetadata[0] > rmetadata[0] {
			winner = PL
		} else if lmetadata[0] < rmetadata[0] {
			winner = PR
		// } else if lHand == TwoPair || lHand == FullHouse {
		// 	if lmetadata[1] < rmetadata[1] {
		// 		winner = PL
		// 	} else if lmetadata[1] > rmetadata[1] {
		// 		winner = PR
		// 	} else {
		// 		winner = highestCard(hL,hR)
		// 	}
		} else {
			winner = highestCard(hL,hR)
		}
	}

	return
}

func highestCard(hl , hr Hand) (winner Winner) {
	res := false

	for i := handSize - 1; i >= 0 && ! res; i -- {
		if hl[i].Num > hr[i].Num {
			res = true
			winner = PL
			log.Println("Highgest card winner Player 1", hl[i].Num , hr[i].Num)
		} else if hl[i].Num < hr[i].Num {
			res = true
			winner = PR
			log.Println("Highgest card winner Player 2", hl[i].Num , hr[i].Num)
		}
	}

	if ! res {
		log.Fatalln("No highgest card?")
	}

	return
}

func parseCards(cards []string) (h Hand){
	h = make([]Card, handSize)
	for p, v := range cards {
		switch v[0]{
			case 'A':
				h[p].Num = 14
			case 'K':
				h[p].Num = 13
			case 'Q':
				h[p].Num = 12
			case 'J':
				h[p].Num = 11
			case 'T':
				h[p].Num = 10
			default:
				h[p].Num = int(v[0]) - '0'
		}

		switch v[1] {
		case 'H':
			h[p].Suit = Hearts
		case 'C':
			h[p].Suit = Club
		case 'S':
			h[p].Suit = Spades
		case 'D':
			h[p].Suit = Diamonds
		default:
			log.Fatalln("Unkwon suit:", v)
		}
	}

	slices.SortFunc(h, func(cl , cr Card) int {
		return cmp.Compare(cl.Num, cr.Num)
	})

	return
}

func parseLine(twohands string) (hl, hr Hand) {
	cards := strings.Split(twohands, " ")

	hl = parseCards(cards[:5])
	hr = parseCards(cards[5:])

	return
}

func test() {
	testHands :=  make([]string, 5)

	testHands[0] = "5H 5C 6S 7S KD 2C 3S 8S 8D TD"
	testHands[1] = "5D 8C 9S JS AC 2C 5C 7D 8S QH"
	testHands[2] = "2D 9C AS AH AC 3D 6D 7D TD QD"
	testHands[3] = "4D 6S 9H QH QC 3D 6D 7H QD QS"
	testHands[4] = "2H 2D 4C 4D 4S 3C 3D 3S 9S 9D"

	plWon := 0
	for i:= 0; i < 5; i++ {
		hl, hr := parseLine(testHands[i])
		/* ** Decide who wins*/
		if whoWins(hl,hr) == PL {
			/* ** Count */
			fmt.Println("Winner is Player 1")
			plWon ++
		} else {
			fmt.Println("Winner is Player 2")
		}
	}

	fmt.Println("Player 1:", plWon)

	return
}

func main(){

	file , err := os.Open("../ProjectEuler54/hands.txt")
	if err != nil {
		log.Fatalln("Error Openning file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	/* * Read line */
	wonPL, hands := 0 , 0
	for scanner.Scan() {
		hands ++
		// Read line (up to 64k characters)
		hl, hr := parseLine(scanner.Text())
		fmt.Printf("Hands: PL[%v] || PR[%v]\n", hl, hr)
		/* ** Decide who wins*/
		if whoWins(hl,hr) == PL {
			/* ** Count */
			wonPL ++
			fmt.Println("Player 1 wins")
		} else {
			fmt.Println("Player 2 wins")
		}
	}

	if err := scanner.Err() ; err != nil {
		log.Fatalln("Error scanning",err)
	}

	fmt.Println("We counted", wonPL , " out of ", hands)

	// test()

	return
}
