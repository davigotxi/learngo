// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	//   We've a string that contains even and odd numbers.

	//   1. Convert the string to an []int
	//
	var numbers []int
	for _, v := range strings.Split(data, " ") {
		n, _ := strconv.Atoi(v)
		numbers = append(numbers, n)
	}

	//   2. Print the slice
	fmt.Printf("%-20v: %v\n", "Numbers", numbers)

	//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
	evenNumbers := numbers[:3]

	//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
	oddNumbers := numbers[3:]

	//   5. Slice it for the two numbers at the middle
	middleNumbers := numbers[2:4]

	//   6. Slice it for the first two numbers
	firstNumbers := numbers[:2]

	//   7. Slice it for the last two numbers (use the len function)
	lastNumbers := numbers[len(numbers)-2:]

	//   8. Slice the evens slice for the last one number
	lastEvenNumbers := evenNumbers[len(evenNumbers)-1:]

	//   9. Slice the odds slice for the last two numbers
	lastOddNumbers := oddNumbers[len(oddNumbers)-2:]

	fmt.Printf("%-20v: %v\n", "Even Numbers", evenNumbers)
	fmt.Printf("%-20v: %v\n", "Odd Numbers", oddNumbers)
	fmt.Printf("%-20v: %v\n", "Middle Numbers", middleNumbers)
	fmt.Printf("%-20v: %v\n", "First Numbers", firstNumbers)
	fmt.Printf("%-20v: %v\n", "Last Numbers", lastNumbers)
	fmt.Printf("%-20v: %v\n", "Last Even Numbers", lastEvenNumbers)
	fmt.Printf("%-20v: %v\n", "Last Odd Numbers", lastOddNumbers)

}
