package rdsop

import (
	"context"

	"server/common/log"
	"server/common/rds"
)

// DeleteAlliance 删除联盟，加入集合
func DeleteAlliance(allianceId int32) {
	_, err := rds.Ins.SAdd(context.Background(), DeleteAlliancesKey(), allianceId).Result()
	if err != nil {
		log.Errorw("DeleteAlliance rds failed", "err", err)
	}
}

// IsAllianceDeleted 联盟是否被删除
func IsAllianceDeleted(allianceId int32) bool {
	deleted, err := rds.Ins.SIsMember(context.Background(), DeleteAlliancesKey(), allianceId).Result()
	if err != nil {
		log.Errorw("IsAllianceDeleted rds failed", "err", err)
		return false
	}
	return deleted
}
