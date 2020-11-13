package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	accessOK = "Access granted to %q.\n"
)

var users = map[string]string{
	"jack":  "1888",
	"david": "1985",
}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]
	if users[u] != p {
		fmt.Printf(errUser, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}
