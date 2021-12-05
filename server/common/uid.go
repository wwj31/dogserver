package common

import (
	"github.com/sony/sonyflake"
	"github.com/wwj31/dogactor/expect"
)

type UID struct {
	insId     uint16
	sonyflake *sonyflake.Sonyflake
}

func NewUID(insId uint16) UID {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{MachineID: func() (uint16, error) {
		return insId, nil
	}})
	return UID{
		sonyflake: sf,
	}
}

func (s *UID) Uuid() uint64 {
	uuid, err := s.sonyflake.NextID()
	expect.Nil(err)
	return uuid
}
