package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"server/proto/outermsg/outer"
)

const (
	key = "6^2f1onGbkK8iIQ%vaC4n*2#oK4rwj&aOa%"
)

func LoginToken(req *outer.LoginReq) string {
	sum := md5.Sum([]byte(fmt.Sprintf("%v%v%v%v", req.PlatformUUID, req.PlatformName, req.ClientVersion, key)))
	return hex.EncodeToString(sum[:])
}

func LoginMD5(uid, rid uint64, new bool) string {
	sum := md5.Sum([]byte(fmt.Sprintf("%v%v%v%v", uid, rid, new, key)))
	return hex.EncodeToString(sum[:])
}
