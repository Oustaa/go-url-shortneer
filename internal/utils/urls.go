package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeURL(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))

	hashInBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}
