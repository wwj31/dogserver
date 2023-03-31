package item

import (
	gogo "github.com/gogo/protobuf/proto"

	"server/common"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type Item struct {
	models.Model
	data   inner.ItemInfo
	modify bool
}

func New(base models.Model) *Item {
	mod := &Item{Model: base}
	mod.data.Items = make(map[int64]int64)
	mod.data.RID = base.Player.RID()

	return mod
}

func (s *Item) Data() gogo.Message {
	return &s.data
}

func (s *Item) OnLogin(first bool) {
	if first {
		s.Add(map[int64]int64{123: 999})
	}

	s.Player.Send2Client(s.itemInfoPush())
}

func (s *Item) OnLogout() {
}

func (s *Item) Enough(items map[int64]int64) bool {
	for id, need := range items {
		if s.data.Items[id] < common.Abs(need) {
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

	for _, v := range items {
		if v > 0 {
			return outer.ERROR_ITEM_USE_POSITIVE_NUM
		}
	}

	s.Player.Item().Add(items, true)
	return outer.ERROR_SUCCESS
}

func (s *Item) Add(items map[int64]int64, push ...bool) {
	defer func() { s.modify = true }()

	for id, count := range items {
		val, ok := s.data.Items[id]
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
		s.data.Items[id] = val
	}

	if len(push) > 0 && len(items) > 0 {
		s.Player.Send2Client(&outer.ItemChangeNotify{
			Items: items,
		})
	}
}

func (s *Item) itemInfoPush() *outer.ItemInfoPush {
	return &outer.ItemInfoPush{
		Items: s.data.Items,
	}
}
