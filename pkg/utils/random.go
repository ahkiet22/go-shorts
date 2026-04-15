package util

import (
	"crypto/sha256"
	"encoding/binary"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func EncodeBase62(num uint64) string {
	if num == 0 {
		return string(charset[0])
	}

	result := make([]byte, 0)

	for num > 0 {
		result = append([]byte{charset[num%62]}, result...)
		num /= 62
	}

	return string(result)
}

func GenerateShortCodeFromURL(url string) string {
	hash := sha256.Sum256([]byte(url))

	num := binary.BigEndian.Uint64(hash[:8])

	return EncodeBase62(num)
}
