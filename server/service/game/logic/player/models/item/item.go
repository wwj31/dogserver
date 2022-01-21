package item

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/expect"
	"server/common"
	"server/common/log"
	"server/db/table"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type Item struct {
	models.Model

	items inner.ItemMap
}

func New(rid uint64, base models.Model) *Item {
	item := &Item{
		Model: base,
		items: inner.ItemMap{Items: make(map[int64]int64, 10)},
	}

	if !base.Player.IsNewRole() {
		tItem := table.Item{RoleId: rid}
		err := base.Player.Gamer().Load(&tItem)
		expect.Nil(err)

		err = proto.Unmarshal(tItem.Bytes, &item.items)
		expect.Nil(err)
	}

	return item
}

func (s *Item) OnLogin() {
	s.Player.Send2Client(s.itemInfoPush())
}

func (s *Item) OnLogout() {
	s.save()
}

func (s *Item) Enough(items map[int64]int64) bool {
	for id, need := range items {
		if s.items.Items[id] < common.Abs(need) {
			return false
		}
	}
	return true
}

func (s *Item) Use(items map[int64]int64) outer.ERROR {
	if len(items) == 0 {
		return outer.ERROR_SUCCESS
	}

	if !s.Player.Item().Enough(items) {
		return outer.ERROR_ITEM_NOT_ENOUGH
	}

	s.Player.Item().Add(items, true)
	return outer.ERROR_SUCCESS
}

func (s *Item) Add(items map[int64]int64, push ...bool) {
	for id, count := range items {
		val, ok := s.items.Items[id]
		if ok {
			if count > 0 {
				val += count
			} else if count < 0 {
				val -= common.Min(val, common.Abs(count))
			} else {
				continue
			}
		} else {
			val = common.Max(0, count)
		}

		s.items.Items[id] = val
	}

	if len(push) > 0 && len(items) > 0 {
		s.Player.Send2Client(&outer.ItemChangeNotify{
			Items: items,
		})
	}

	s.save()
}

func (s *Item) itemInfoPush() *outer.ItemInfoPush {
	return &outer.ItemInfoPush{
		Items: s.items.Items,
	}
}

func (s *Item) save() {
	s.SetTable(&table.Item{
		RoleId:    s.Player.Role().RoleId(),
		Bytes:     s.marshal(&s.items),
		ItemCount: len(s.items.Items),
	})
}

func (s *Item) marshal(msg proto.Message) []byte {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		log.Errorw("proto marshal error", "err", err)
	}
	return bytes
}
