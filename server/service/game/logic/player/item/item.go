package item

import (
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/expect"
	"server/common"
	"server/db/table"
	"server/proto/message"
	"server/service/game/logic/model"
)

type Item struct {
	model.Model

	items *message.ItemMap
}

func New(rid uint64, base model.Model) *Item {
	tItem := table.Item{RoleId: rid}

	// 不是新号，找db要数据
	if !base.Player.IsNewRole() {
		err := base.Player.Load(&tItem)
		expect.Nil(err)
	}

	items := &message.ItemMap{Items: make(map[int64]int64, 10)}
	if tItem.Items != nil {
		err := proto.Unmarshal(tItem.Items, items)
		expect.Nil(err)
	}

	role := &Item{
		Model: base,
		items: items,
	}

	return role
}

func (s *Item) OnLogin() {
	s.Player.Send2Client(s.itemInfoPush())
}

func (s *Item) OnStop() {
	s.save()
}

func (s *Item) Add(items map[int64]int64, push ...bool) {
	for id, count := range items {
		val := count
		if old, ok := s.items.Items[id]; ok {
			if count > 0 {
				val += count
			} else if count < 0 {
				val -= common.Min(old, -count)
			} else {
				continue
			}
		}
		s.items.Items[id] = val
	}

	if len(push) > 0 {
		s.Player.Send2Client(&message.ItemMap{
			Items: items,
		})
	}

	s.save()
}

func (s *Item) itemInfoPush() *message.ItemInfoPush {
	return &message.ItemInfoPush{
		Items: s.items.Items,
	}
}

func (s *Item) save() {
	s.SetTable(&table.Item{
		RoleId: s.Player.Role().RoleId(),
		Items:  s.marshal(s.items),
	})
}

func (s *Item) marshal(msg proto.Message) []byte {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		s.Log().KV("err", err).Error("proto marshal error")
	}
	return bytes
}
