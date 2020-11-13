package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

// ░ █ ▄ ▀

type placeholder [5]string

var zero = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}
var one = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}
var two = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}
var three = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}
var four = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}
var five = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}
var six = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}
var seven = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}
var eight = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}
var nine = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}
var separator = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var digits = [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}

func main() {

	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			separator,
			digits[min/10], digits[min%10],
			separator,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == separator && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}

}
