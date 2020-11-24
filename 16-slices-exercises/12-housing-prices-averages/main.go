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
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// Solve this exercise by using your previous solution for
	// the "Housing Prices" exercise.
	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 15*5))

	var (
		locs                                   []string
		sizes, beds, baths, prices             []int
		sumSizes, sumBeds, sumBaths, sumPrices float64
	)

	houses := strings.Split(data, "\n")
	for i, house := range houses {
		cols := strings.Split(house, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)
		sumSizes += float64(n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)
		sumBeds += float64(n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)
		sumBaths += float64(n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
		sumPrices += float64(n)

		fmt.Printf("%-15v%-15v%-15v%-15v%-15v\n", locs[i], sizes[i], beds[i], baths[i], prices[i])
	}
	fmt.Printf("%s\n", strings.Repeat("=", 15*5))

	fmt.Printf("%-15v%-15.2f%-15.2f%-15.2f%-15.2f\n", "",
		sumSizes/float64(len(houses)),
		sumBeds/float64(len(houses)),
		sumBaths/float64(len(houses)),
		sumPrices/float64(len(houses)))

}
