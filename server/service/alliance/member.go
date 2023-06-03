package alliance

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"

	"server/common/log"
	"server/common/mongodb"
)

type Member struct {
	*Alliance `bson:"-"`

	RID      string   `bson:"_id"`
	ShortId  int64    `bson:"short_id"`
	Position Position `bson:"position"`
}

func (m *Member) Save() {
	if _, err := mongodb.Ins.Collection(m.Coll()).UpdateByID(context.Background(),
		m.RID,
		bson.M{"$set": m},
		options.Update().SetUpsert(true)); err != nil {
		log.Errorw("member save failed", "err", err, "rid", m.RID, "shortId", m.ShortId, "position", m.Position)
		return
	}
}
