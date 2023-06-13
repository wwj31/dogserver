package outer

//go install github.com/wwj31/spawner@latest\
//go:generate go run github.com/wwj31/spawner@latest -pool=false

func (s Msg) Int32() int32 {
	return int32(s)
}

func (s Msg) UInt32() uint32 {
	return uint32(s)
}

func (s Msg) Int64() uint64 {
	return uint64(s)
}

func (s Msg) UInt64() uint64 {
	return uint64(s)
}

func (s GameType) Int32() int32 {
	return int32(s)
}

func (s GameType) UInt32() uint32 {
	return uint32(s)
}

func (s GameType) Int64() uint64 {
	return uint64(s)
}

func (s GameType) UInt64() uint64 {
	return uint64(s)
}
