package comm

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func GetSha256(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}