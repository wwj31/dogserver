package alliance

import (
	"time"

	"server/common"
)

type Member struct {
	RID       string          `bson:"RID"`
	ShortId   int64           `bson:"short_id"`
	Name      string          `bson:"name"`
	Position  Position        `bson:"position"`
	OnlineAt  time.Time       `bson:"online_at"`
	OfflineAt time.Time       `bson:"offline_at"`
	GSession  common.GSession `bson:"-"`
}
