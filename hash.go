package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func hashEncode(b []byte, length int) (string, error) {
	if length < 0 {
		return "", fmt.Errorf("Length must be non-negative")
	}

	// sha256sum(bytes) -> 256 bits
	h := sha256.Sum256(b)
	// base64(hash) -> 42.6 runes
	b64 := base64.RawStdEncoding.EncodeToString(h[:])
	// remove '+' or '/' -> over 32 runes (probably)
	b64 = strings.Replace(b64, "+", "", -1)
	b64 = strings.Replace(b64, "/", "", -1)

	if len(b64) < length {
		// probably not reach here
		return "", fmt.Errorf("Retry another password")
	}

	return b64[:length], nil
}
