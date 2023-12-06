package alliance

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"

	"server/common/log"
	"server/common/mongodb"
	"server/proto/outermsg/outer"
)

type Manifest struct {
	*Alliance `bson:"-"`

	UID        string `bson:"_id"`
	GameType   int32
	GameParams *outer.GameParams
}

func (m *Manifest) Save() {
	if _, err := mongodb.Ins.Collection(m.ManifestColl()).UpdateByID(context.Background(),
		m.UID,
		bson.M{"$set": m},
		options.Update().SetUpsert(true)); err != nil {
		log.Errorw("member save failed", "err", err, "rid", m.UID, "param", m.GameParams.String())
		return
	}
}

func (m *Manifest) Delete() {
	if _, err := mongodb.Ins.Collection(m.ManifestColl()).DeleteOne(context.Background(), bson.M{"_id": m.UID}); err != nil {
		log.Errorw("member save failed", "err", err, "rid", m.UID, "param", m.GameParams.String())
		return
	}
	delete(m.manifests, m.UID)

}

func (m *Manifest) ArgEqual(other *Manifest) bool {
	return m.GameParams.Mahjong.String() == other.GameParams.Mahjong.String() &&
		m.GameParams.FasterRun.String() == other.GameParams.FasterRun.String() &&
		m.GameParams.NiuNiu.String() == other.GameParams.NiuNiu.String()
}

func (m *Manifest) ToPB() *outer.Manifest {
	return &outer.Manifest{
		Id:         m.UID,
		GameType:   outer.GameType(m.GameType),
		GameParams: m.GameParams,
	}
}
