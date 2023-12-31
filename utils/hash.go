package utils

import (
	"crypto/sha1"
	"fmt"
)

// Hash string using SHA1
func Hash(text string) string {
	// s := "C3h1xJgUy4DRcFSB9Nyz9Bldrm-eWUG4ttIi5Gl5icb3-bkdwTEtfFobDjo3aBHmZVjschX5A-168jXk"

	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(text))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	res := fmt.Sprintf("%x\n", bs)

	return res
}
