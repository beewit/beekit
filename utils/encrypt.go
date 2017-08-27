package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1Encode(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}
