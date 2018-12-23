package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get length
	if len(os.Args) != 2 {
		fmt.Println("Usage: pwenc [length]")
		os.Exit(1)
	}
	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read password
	pw, err := passwordPrompt()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Encode password
	h, err := hashEncode(pw, length)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s (%d)\n", h, len(h))
}
