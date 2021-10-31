package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync/atomic"
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

func UIDToRID(uid int64, num int) int64 {
	return uid*10 + int64(num)
}

func RIDToUID(rid int64) int64 {
	return rid / 10
}

var objectIdCounter uint32 = 0

func GetNewId() (id uint64) {
	currentTime := time.Now().UTC().Unix()
	// 取出末尾18位
	i := atomic.AddUint32(&objectIdCounter, 1) & 0x3FFFF
	// 组装
	return uint64(currentTime)<<32 | uint64(i)<<14 | uint64(1)
}
