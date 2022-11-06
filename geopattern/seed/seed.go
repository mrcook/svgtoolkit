// Package seed is used for transforming colours.
package seed

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

// Seed represents an encoded string to be used as a seed value.
type Seed struct {
	seed string
}

// New returns a Seed instance, with the given string encrypted as SHA-1.
func New(s string) *Seed {
	h := sha1.New()
	h.Write([]byte(s))

	result := h.Sum(nil)
	return &Seed{seed: fmt.Sprintf("%x", result)}
}

// ToInt returns a value taken from the seed at the given index, of the given length.
//
// TODO: ignore errors, panic, or return the error?
func (s *Seed) ToInt(index, length int) int {
	if length < 1 {
		length = 1
	}
	if index+length >= len(s.seed) {
		return 0
	}

	str := s.seed[index : index+length]

	hex, err := strconv.ParseInt(str, 16, 0)
	if err != nil {
		return 0
	}

	return int(hex)
}
