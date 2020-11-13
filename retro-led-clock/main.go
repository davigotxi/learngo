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

var digits = [10]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}

func main() {

	//now := time.Now()

	//hour := now.Hour()

	for l := 0; l < 5; l++ {
		for d := 0; d < 10; d++ {
			fmt.Printf("%s ", digits[d][l])
		}
		fmt.Println()
	}
	fmt.Println()
}
