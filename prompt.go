package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func readPassword(prompt string) ([]byte, error) {
	fmt.Printf("%s", prompt)
	pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Printf("\n")

	return pw, err
}

func passwordPrompt() ([]byte, error) {
	empty := []byte{}

	pw, err := readPassword("Enter Password: ")
	if err != nil {
		return empty, err
	}

	verify, err := readPassword("Verify Password: ")
	if err != nil {
		return empty, err
	}

	if !bytes.Equal(pw, verify) {
		return empty, fmt.Errorf("Password do not match, please retry")
	}

	return pw, nil
}
