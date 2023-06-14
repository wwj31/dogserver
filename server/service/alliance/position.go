package alliance

import "server/proto/outermsg/outer"

type Position int32

const (
	Normal       Position = iota + 1 // 普通成员
	ViceCaptain                      // 小队长
	Captain                          // 队长
	Manager                          // 管理员
	DeputyMaster                     // 副盟主
	Master                           // 盟主
)

func (p Position) Int32() int32 {
	return int32(p)
}

func (p Position) Pb() outer.Position {
	return outer.Position(p)
}
