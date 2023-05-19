package req_test

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/spf13/cast"
)

const secretKey = "62f1onGbkK8iIQBvaC4n*2#oK4rwj&aOa5"

func addSign(r *http.Request) {
	t := time.Now().Unix()
	now := cast.ToString(t)
	token := calculateSignature(now, secretKey)
	r.Header.Set("X-Time", now)
	r.Header.Set("X-Signature", token)
}

func calculateSignature(time, secretKey string) string {
	str := time + secretKey
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}
