package rdskey

func LockLoginKey(key string) string {
	return "lock:login:{" + key + "}"
}

func SessionKey(rid string) string {
	return "session:{" + rid + "}"
}
