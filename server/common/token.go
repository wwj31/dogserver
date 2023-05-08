package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const (
	key = "6^2f1onGbkK8iIQ%vaC4n*2#oK4rwj&aOa%"
)

func EnterGameToken(rid string, new bool) string {
	sum := md5.Sum([]byte(fmt.Sprintf("%v%v%v", rid, new, key)))
	return hex.EncodeToString(sum[:])
}
