package common

type Pos uint32

func GetNum(x, y uint16) Pos {
	return Pos(x)*10000 + Pos(y)
}

func (pos Pos) GetPos() (x, y uint16) {
	x = uint16(pos / 10000)
	y = uint16(pos % 10000)
	return
}
