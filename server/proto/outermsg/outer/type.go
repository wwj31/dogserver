package outer

//go install github.com/wwj31/spawner@v0.0.6
//go:generate spawner -pool=true

func (s MSG) Int32() int32 {
	return int32(s)
}

func (s MSG) UInt32() uint32 {
	return uint32(s)
}

func (s MSG) Int64() uint64 {
	return uint64(s)
}

func (s MSG) UInt64() uint64 {
	return uint64(s)
}
