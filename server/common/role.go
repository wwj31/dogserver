package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	LOGIN_SESSION_KEY = "6^2f1onGbkK8iIQ%vaC4n*2#oK4rwj&aOa%"
)

func LoginSecurityCode(rid, uid, sid int64, newPlayer bool, timestamp int64) string {
	sum := md5.Sum([]byte(fmt.Sprintf("%v%v%v%v%v%v", rid, uid, sid, timestamp, newPlayer, LOGIN_SESSION_KEY)))
	return hex.EncodeToString(sum[:])
}

func IsLogin(rid, uid, sid int64, timestamp int64, newPlayer bool, securityCode string) bool {
	const MaxSessionTime = 7200
	if time.Now().Unix()-timestamp > MaxSessionTime {
		return false
	}
	if securityCode != LoginSecurityCode(rid, uid, sid, newPlayer, timestamp) {
		return false
	}
	return true
}
