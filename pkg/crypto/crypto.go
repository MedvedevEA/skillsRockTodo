package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func GetHash(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
func CheckHash(text string, hash string) bool {
	textHash := GetHash(text)
	return textHash == hash
}

func GenerateSecureToken() string {
	return GetHash(strconv.FormatInt(rand.Int63(), 10))
}
