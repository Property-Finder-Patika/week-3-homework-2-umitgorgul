package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	counterForSameLocation := 0           // counter if numbers are the same and also numbers are in same location
	counterForNotSameLocation := 0        // counter if numbers are the same but also numbers aren't in same location
	answer := rand.Intn(9999-1000) + 1000 // for creating a random number in rage of 1000 to 9999
	fmt.Println("Enter your guess: ")
	var guess string        // for user input
	fmt.Scanf("%v", &guess) // get user input
	if len(guess) != 4 {
		fmt.Println(`duzgun sayi ver programi durdur`)

	}
	if guess[0] == 0 {
		fmt.Println(`duzgun sayi ver  programi durdur`)

	}
	answerStr := strconv.Itoa(answer)     // int to string for iterate and check variables same
	for i := 0; i < len(answerStr); i++ { // for each character in random number
		for x := 0; x < len(guess); x++ { // for each character in guessed number
			a := answerStr[i]
			b := guess[x]
			if a == b && i == x { // check if characters are same and in same order
				counterForSameLocation++
			}
			if a == b && i != x { // check if characters are same and aren't in same order
				counterForNotSameLocation--
			}

		}
	}
	fmt.Println(answer)
	fmt.Println(counterForSameLocation, counterForNotSameLocation)
}
