package item

import (
	"server/common"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/game/logic/player/models"
)

type Item struct {
	models.Model
	items  inner.ItemInfo
	modify bool
}

func New(base models.Model) *Item {
	mod := &Item{
		Model: base,
		items: inner.ItemInfo{Items: make(map[int64]int64, 10)},
	}

	//first
	mod.Add(map[int64]int64{123: 999})
	return mod
}

func (s *Item) OnLogin() {
	s.Player.Send2Client(s.itemInfoPush())
}

func (s *Item) OnLogout() {
}

func (s *Item) OnSave() {
}

func (s *Item) Enough(items map[int64]int64) bool {
	for id, need := range items {
		if s.items.Items[id] < common.Abs64(need) {
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
		val, ok := s.items.Items[id]
		if ok {
			if count > 0 {
				val += count
			} else if count < 0 {
				val -= common.Min64(val, common.Abs64(count))
			} else {
				continue
			}
		} else {
			val = common.Max64(0, count)
		}
		s.items.Items[id] = val
	}

	if len(push) > 0 && len(items) > 0 {
		s.Player.Send2Client(&outer.ItemChangeNotify{
			Items: items,
		})
	}
}

func (s *Item) itemInfoPush() *outer.ItemInfoPush {
	return &outer.ItemInfoPush{
		Items: s.items.Items,
	}
}
