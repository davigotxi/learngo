package main

import (
	"fmt"
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

var digits = [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}

func main() {

	//now := time.Now()

	//hour := now.Hour()

	for _, digit := range digits {
		for _, line := range digit {
			fmt.Println(line)
		}
		fmt.Println()
	}

	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line], " ")
		}
		fmt.Println()
	}
	fmt.Println()

}
