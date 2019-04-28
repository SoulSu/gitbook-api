package server

import (
	"crypto/rand"
	"encoding/hex"
)

func GenRequestID() string {
	return realRandom(30)
}

func realRandom(l int) string {
	b := make([]byte, l)
	_, _ = rand.Read(b)

	return hex.EncodeToString(b)
}
