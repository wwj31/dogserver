package event

import "server/service/game/logic/player/models/role/typ"

type ChangeAttribute struct {
	Type   typ.Attribute
	OldVal int64
	NewVal int64
}
