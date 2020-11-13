package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func printClock(hour, min, sec, splitSec int) {
	screen.MoveTopLeft()
	clockBuffer := [...]placeholder{
		blank, blank, blank, blank, blank, blank, blank, blank, blank, blank,
		digits[hour/10],
		digits[hour%10],
		separator,
		digits[min/10],
		digits[min%10],
		separator,
		digits[sec/10],
		digits[sec%10],
		dot,
		digits[splitSec],
	}
	startPosition := sec % 20
	clock := [...]placeholder{
		clockBuffer[startPosition%20],
		clockBuffer[(startPosition+1)%20],
		clockBuffer[(startPosition+2)%20],
		clockBuffer[(startPosition+3)%20],
		clockBuffer[(startPosition+4)%20],
		clockBuffer[(startPosition+5)%20],
		clockBuffer[(startPosition+6)%20],
		clockBuffer[(startPosition+7)%20],
		clockBuffer[(startPosition+8)%20],
		clockBuffer[(startPosition+9)%20],
	}

	for line := range clock[0] {
		for index, digit := range clock {
			next := clock[index][line]
			if (digit == separator || digit == dot) && sec%2 == 0 {
				next = "   "
			}
			fmt.Print(next, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printAlarm() {
	screen.MoveTopLeft()
	for _, line := range alarm {
		fmt.Println(line)
	}
	fmt.Println()
}

func main() {

	screen.Clear()
	for {

		now := time.Now()
		hour, min, sec, nano := now.Hour(), now.Minute(), now.Second(), now.Nanosecond()
		splitSec := nano / 1e8

		if sec%10 == 0 {
			printAlarm()
		} else {
			printClock(hour, min, sec, splitSec)
		}

		time.Sleep(time.Millisecond * 100)
	}
}
