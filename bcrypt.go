package main

import (
	"crypto/sha256"
)

func hashfunc(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
