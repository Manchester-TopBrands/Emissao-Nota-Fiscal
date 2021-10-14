package auth

import (
	"crypto/rand"
)

// GenKey ...
func genKey() []byte {
	b := make([]byte, 12)
	rand.Read(b)
	return b
}
