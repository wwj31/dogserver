package alliance

type Position int64

const (
	Normal        Position = iota + 1 // 普通成员
	DeputyCaptain                     // 小队长
	Captain                           // 队长
	Manager                           // 管理员
	DeputyMaster                      // 副盟主
	Master                            // 盟主
)
