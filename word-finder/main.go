package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" +
	"lazy cat jumps again and again and again"

func main() {

	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
		s := strings.ToLower(q)
	search:
		for i, w := range words {
			switch s {
			case "and", "or", "the":
				break search
			}
			if s == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}
