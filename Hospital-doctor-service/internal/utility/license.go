package utility

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateUniqueLicense generates a random unique license number of specified length
func GenerateUniqueLicense(length int) (string, error) {
	license := make([]byte, length)
	for i := 0; i < length; i++ {
		// Generate a random index for the charset
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		license[i] = charset[index.Int64()]
	}
	return string(license), nil
}
