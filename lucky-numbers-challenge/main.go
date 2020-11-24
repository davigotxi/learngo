package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Reqs
// 1) Special bonus message is player wins on 1st round
// 2) Random won/lost messages
// 3) Player to guess 2 numbers instead of one. When either of them
//    picked by computer, print the winning message

const (
	maxTurns = 5
	usage    = `Welcome to the Luck Number Game!

The program will pick %d random numbers.
Your mission is to guess one of those numbers, you can pick 2 numbers

The greater your numbers are, the harder it gets.

Wanna play?`
)

var winBonusMsg = "YOU WON ON THE FIRST TURN!!!"

var winMsgs = [...]string{
	"WIN Message 1",
	"WIN Message 2",
	"WIN Message 3"}

var lostMsgs = [...]string{
	"LOST Message 1",
	"LOST Message 2",
	"LOST Message 3"}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func guessArg(s string) (int, error) {

	guess, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Not a number")
		return guess, err
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number")
		return guess, errors.New("Cannot work with negative numbers")
	}
	return guess, nil
}

func main() {

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess1, err := guessArg(args[0])
	if err != nil {
		return
	}

	guess2, err := guessArg(args[1])
	if err != nil {
		return
	}

	for turn := 0; turn < maxTurns; turn++ {

		n := rand.Intn(max(guess1, guess2) + 1)

		if n == guess2 || n == guess1 {
			if turn == 0 {
				fmt.Println(winBonusMsg)
			} else {

				fmt.Println(winMsgs[rand.Intn(len(winMsgs))])
			}
			return
		}
	}

	fmt.Println(lostMsgs[rand.Intn(len(lostMsgs))])
}
