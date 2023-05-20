package alliance

import "server/common"

type Member struct {
	RID      string          `bson:"RID"`
	ShortId  int64           `bson:"short_id"`
	Name     string          `bson:"name"`
	Position Position        `bson:"position"`
	GSession common.GSession `bson:"-"`
}
