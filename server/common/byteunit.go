package common

// go install golang.org/x/tools/cmd/stringer

//go:generate stringer -type ByteUnit
type ByteUnit int64

const (
	B  ByteUnit = 1
	KB          = 1024 * B
	MB          = 1024 * KB
	GB          = 1024 * MB
	TB          = 1024 * GB
)
