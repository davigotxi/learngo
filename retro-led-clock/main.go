package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func printClock(hour, min, sec, splitSec int) {
	screen.MoveTopLeft()
	clock := [...]placeholder{
		digits[hour/10], digits[hour%10],
		separator,
		digits[min/10], digits[min%10],
		separator,
		digits[sec/10], digits[sec%10],
		dot,
		digits[splitSec],
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
