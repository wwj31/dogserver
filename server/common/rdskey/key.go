package rdskey

import "fmt"

func LockLoginKey(key string) string {
	return "lock:login:" + key
}

func SessionKey(rid string) string {
	return "session:" + rid
}

func ShortIDKey() string {
	return "shortId"
}

func PlayerInfoKey(shortId int64) string {
	return fmt.Sprintf("playerinfo:%v", shortId)
}
