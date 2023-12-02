// Code generated by "spawner -pool=false"; DO NOT EDIT.

package outer

type factory func() interface{}

func Spawner(name string, newPool ...bool) (interface{}, bool) {
	f, ok := spawner[name]
	if !ok {
		return nil, ok
	}
	return f(), true
}

func Put(name string, x interface{}) {}

var spawner = map[string]factory{
	"outer.UpdateGoldNtf":                     func() interface{} { return &UpdateGoldNtf{} },
	"outer.Unknown":                           func() interface{} { return &Unknown{} },
	"outer.TotalContributeRsp":                func() interface{} { return &TotalContributeRsp{} },
	"outer.TotalContributeReq":                func() interface{} { return &TotalContributeReq{} },
	"outer.SetScoreForDownRsp":                func() interface{} { return &SetScoreForDownRsp{} },
	"outer.SetScoreForDownReq":                func() interface{} { return &SetScoreForDownReq{} },
	"outer.SetRoleInfoRsp":                    func() interface{} { return &SetRoleInfoRsp{} },
	"outer.SetRoleInfoReq":                    func() interface{} { return &SetRoleInfoReq{} },
	"outer.SetMemberPositionRsp":              func() interface{} { return &SetMemberPositionRsp{} },
	"outer.SetMemberPositionReq":              func() interface{} { return &SetMemberPositionReq{} },
	"outer.SetAgentDownRebateRsp":             func() interface{} { return &SetAgentDownRebateRsp{} },
	"outer.SetAgentDownRebateReq":             func() interface{} { return &SetAgentDownRebateReq{} },
	"outer.SearchPlayerInfoRsp":               func() interface{} { return &SearchPlayerInfoRsp{} },
	"outer.SearchPlayerInfoReq":               func() interface{} { return &SearchPlayerInfoReq{} },
	"outer.RoomPlayerOnlineNtf":               func() interface{} { return &RoomPlayerOnlineNtf{} },
	"outer.RoomPlayerLeaveNtf":                func() interface{} { return &RoomPlayerLeaveNtf{} },
	"outer.RoomPlayerInfo":                    func() interface{} { return &RoomPlayerInfo{} },
	"outer.RoomPlayerEnterNtf":                func() interface{} { return &RoomPlayerEnterNtf{} },
	"outer.RoomListRsp":                       func() interface{} { return &RoomListRsp{} },
	"outer.RoomListReq":                       func() interface{} { return &RoomListReq{} },
	"outer.RoomInfo":                          func() interface{} { return &RoomInfo{} },
	"outer.RoleInfo":                          func() interface{} { return &RoleInfo{} },
	"outer.RecordingMessage":                  func() interface{} { return &RecordingMessage{} },
	"outer.Recording":                         func() interface{} { return &Recording{} },
	"outer.ReceiveMailItemRsp":                func() interface{} { return &ReceiveMailItemRsp{} },
	"outer.ReceiveMailItemReq":                func() interface{} { return &ReceiveMailItemReq{} },
	"outer.RebateScoreRsp":                    func() interface{} { return &RebateScoreRsp{} },
	"outer.RebateScoreReq":                    func() interface{} { return &RebateScoreReq{} },
	"outer.RebateParams":                      func() interface{} { return &RebateParams{} },
	"outer.RebateDetailInfo":                  func() interface{} { return &RebateDetailInfo{} },
	"outer.ReadMailRsp":                       func() interface{} { return &ReadMailRsp{} },
	"outer.ReadMailReq":                       func() interface{} { return &ReadMailReq{} },
	"outer.RangeParams":                       func() interface{} { return &RangeParams{} },
	"outer.PlayerInfo":                        func() interface{} { return &PlayerInfo{} },
	"outer.PlayerDailyStat":                   func() interface{} { return &PlayerDailyStat{} },
	"outer.PlayCardsRecord":                   func() interface{} { return &PlayCardsRecord{} },
	"outer.PlayCardTips":                      func() interface{} { return &PlayCardTips{} },
	"outer.NiuNiuToBettingRsp":                func() interface{} { return &NiuNiuToBettingRsp{} },
	"outer.NiuNiuToBettingReq":                func() interface{} { return &NiuNiuToBettingReq{} },
	"outer.NiuNiuToBeMasterRsp":               func() interface{} { return &NiuNiuToBeMasterRsp{} },
	"outer.NiuNiuToBeMasterReq":               func() interface{} { return &NiuNiuToBeMasterReq{} },
	"outer.NiuNiuStopCountDownNtf":            func() interface{} { return &NiuNiuStopCountDownNtf{} },
	"outer.NiuNiuStartCountDownNtf":           func() interface{} { return &NiuNiuStartCountDownNtf{} },
	"outer.NiuNiuShowCardsRsp":                func() interface{} { return &NiuNiuShowCardsRsp{} },
	"outer.NiuNiuShowCardsReq":                func() interface{} { return &NiuNiuShowCardsReq{} },
	"outer.NiuNiuShowCardsNtf":                func() interface{} { return &NiuNiuShowCardsNtf{} },
	"outer.NiuNiuSettlementNtf":               func() interface{} { return &NiuNiuSettlementNtf{} },
	"outer.NiuNiuSelectMasterNtf":             func() interface{} { return &NiuNiuSelectMasterNtf{} },
	"outer.NiuNiuSelectBettingNtf":            func() interface{} { return &NiuNiuSelectBettingNtf{} },
	"outer.NiuNiuReadyNtf":                    func() interface{} { return &NiuNiuReadyNtf{} },
	"outer.NiuNiuPlayerInfo":                  func() interface{} { return &NiuNiuPlayerInfo{} },
	"outer.NiuNiuParams":                      func() interface{} { return &NiuNiuParams{} },
	"outer.NiuNiuMasterNtf":                   func() interface{} { return &NiuNiuMasterNtf{} },
	"outer.NiuNiuGameInfo":                    func() interface{} { return &NiuNiuGameInfo{} },
	"outer.NiuNiuFinishShowCardsNtf":          func() interface{} { return &NiuNiuFinishShowCardsNtf{} },
	"outer.NiuNiuDealNtf":                     func() interface{} { return &NiuNiuDealNtf{} },
	"outer.NiuNiuCardsGroup":                  func() interface{} { return &NiuNiuCardsGroup{} },
	"outer.NiuNiuBettingNtf":                  func() interface{} { return &NiuNiuBettingNtf{} },
	"outer.ModifyPasswordRsp":                 func() interface{} { return &ModifyPasswordRsp{} },
	"outer.ModifyPasswordReq":                 func() interface{} { return &ModifyPasswordReq{} },
	"outer.MailListRsp":                       func() interface{} { return &MailListRsp{} },
	"outer.MailListReq":                       func() interface{} { return &MailListReq{} },
	"outer.MailInfo":                          func() interface{} { return &MailInfo{} },
	"outer.Mail":                              func() interface{} { return &Mail{} },
	"outer.MahjongPlayerInfo":                 func() interface{} { return &MahjongPlayerInfo{} },
	"outer.MahjongParams":                     func() interface{} { return &MahjongParams{} },
	"outer.MahjongBTETurnNtf":                 func() interface{} { return &MahjongBTETurnNtf{} },
	"outer.MahjongBTESettlementPlayerData":    func() interface{} { return &MahjongBTESettlementPlayerData{} },
	"outer.MahjongBTESettlementNtf":           func() interface{} { return &MahjongBTESettlementNtf{} },
	"outer.MahjongBTEReadyRsp":                func() interface{} { return &MahjongBTEReadyRsp{} },
	"outer.MahjongBTEReadyReq":                func() interface{} { return &MahjongBTEReadyReq{} },
	"outer.MahjongBTEReadyNtf":                func() interface{} { return &MahjongBTEReadyNtf{} },
	"outer.MahjongBTEPlayingNtf":              func() interface{} { return &MahjongBTEPlayingNtf{} },
	"outer.MahjongBTEPlayerReadyNtf":          func() interface{} { return &MahjongBTEPlayerReadyNtf{} },
	"outer.MahjongBTEPlayCardRsp":             func() interface{} { return &MahjongBTEPlayCardRsp{} },
	"outer.MahjongBTEPlayCardReq":             func() interface{} { return &MahjongBTEPlayCardReq{} },
	"outer.MahjongBTEOperateRsp":              func() interface{} { return &MahjongBTEOperateRsp{} },
	"outer.MahjongBTEOperateReq":              func() interface{} { return &MahjongBTEOperateReq{} },
	"outer.MahjongBTEOperaNtf":                func() interface{} { return &MahjongBTEOperaNtf{} },
	"outer.MahjongBTEHuResultNtf":             func() interface{} { return &MahjongBTEHuResultNtf{} },
	"outer.MahjongBTEHuInfo":                  func() interface{} { return &MahjongBTEHuInfo{} },
	"outer.MahjongBTEGangResultNtf":           func() interface{} { return &MahjongBTEGangResultNtf{} },
	"outer.MahjongBTEGameInfo":                func() interface{} { return &MahjongBTEGameInfo{} },
	"outer.MahjongBTEFinialSettlement":        func() interface{} { return &MahjongBTEFinialSettlement{} },
	"outer.MahjongBTEFinialPlayerInfo":        func() interface{} { return &MahjongBTEFinialPlayerInfo{} },
	"outer.MahjongBTEExchange3Rsp":            func() interface{} { return &MahjongBTEExchange3Rsp{} },
	"outer.MahjongBTEExchange3Req":            func() interface{} { return &MahjongBTEExchange3Req{} },
	"outer.MahjongBTEExchange3PlayerReadyNtf": func() interface{} { return &MahjongBTEExchange3PlayerReadyNtf{} },
	"outer.MahjongBTEExchange3Ntf":            func() interface{} { return &MahjongBTEExchange3Ntf{} },
	"outer.MahjongBTEExchange3EndNtf":         func() interface{} { return &MahjongBTEExchange3EndNtf{} },
	"outer.MahjongBTEDecideMasterNtf":         func() interface{} { return &MahjongBTEDecideMasterNtf{} },
	"outer.MahjongBTEDecideIgnoreRsp":         func() interface{} { return &MahjongBTEDecideIgnoreRsp{} },
	"outer.MahjongBTEDecideIgnoreReq":         func() interface{} { return &MahjongBTEDecideIgnoreReq{} },
	"outer.MahjongBTEDecideIgnoreReadyNtf":    func() interface{} { return &MahjongBTEDecideIgnoreReadyNtf{} },
	"outer.MahjongBTEDecideIgnoreNtf":         func() interface{} { return &MahjongBTEDecideIgnoreNtf{} },
	"outer.MahjongBTEDecideIgnoreEndNtf":      func() interface{} { return &MahjongBTEDecideIgnoreEndNtf{} },
	"outer.MahjongBTEDealNtf":                 func() interface{} { return &MahjongBTEDealNtf{} },
	"outer.LoginRsp":                          func() interface{} { return &LoginRsp{} },
	"outer.LoginReq":                          func() interface{} { return &LoginReq{} },
	"outer.LeaveRoomRsp":                      func() interface{} { return &LeaveRoomRsp{} },
	"outer.LeaveRoomReq":                      func() interface{} { return &LeaveRoomReq{} },
	"outer.KickOutMemberRsp":                  func() interface{} { return &KickOutMemberRsp{} },
	"outer.KickOutMemberReq":                  func() interface{} { return &KickOutMemberReq{} },
	"outer.JoinRoomRsp":                       func() interface{} { return &JoinRoomRsp{} },
	"outer.JoinRoomReq":                       func() interface{} { return &JoinRoomReq{} },
	"outer.InviteAllianceRsp":                 func() interface{} { return &InviteAllianceRsp{} },
	"outer.InviteAllianceReq":                 func() interface{} { return &InviteAllianceReq{} },
	"outer.HeartRsp":                          func() interface{} { return &HeartRsp{} },
	"outer.HeartReq":                          func() interface{} { return &HeartReq{} },
	"outer.GoldRecordsRsp":                    func() interface{} { return &GoldRecordsRsp{} },
	"outer.GoldRecordsReq":                    func() interface{} { return &GoldRecordsReq{} },
	"outer.GoldRecords":                       func() interface{} { return &GoldRecords{} },
	"outer.GameParams":                        func() interface{} { return &GameParams{} },
	"outer.FasterRunTurnNtf":                  func() interface{} { return &FasterRunTurnNtf{} },
	"outer.FasterRunSettlementPlayerData":     func() interface{} { return &FasterRunSettlementPlayerData{} },
	"outer.FasterRunSettlementNtf":            func() interface{} { return &FasterRunSettlementNtf{} },
	"outer.FasterRunReadyRsp":                 func() interface{} { return &FasterRunReadyRsp{} },
	"outer.FasterRunReadyReq":                 func() interface{} { return &FasterRunReadyReq{} },
	"outer.FasterRunReadyNtf":                 func() interface{} { return &FasterRunReadyNtf{} },
	"outer.FasterRunPlayerReadyNtf":           func() interface{} { return &FasterRunPlayerReadyNtf{} },
	"outer.FasterRunPlayerInfo":               func() interface{} { return &FasterRunPlayerInfo{} },
	"outer.FasterRunPlayCardRsp":              func() interface{} { return &FasterRunPlayCardRsp{} },
	"outer.FasterRunPlayCardReq":              func() interface{} { return &FasterRunPlayCardReq{} },
	"outer.FasterRunPassRsp":                  func() interface{} { return &FasterRunPassRsp{} },
	"outer.FasterRunPassReq":                  func() interface{} { return &FasterRunPassReq{} },
	"outer.FasterRunParams":                   func() interface{} { return &FasterRunParams{} },
	"outer.FasterRunGameStartNtf":             func() interface{} { return &FasterRunGameStartNtf{} },
	"outer.FasterRunGameInfo":                 func() interface{} { return &FasterRunGameInfo{} },
	"outer.FasterRunFinialSettlement":         func() interface{} { return &FasterRunFinialSettlement{} },
	"outer.FasterRunFinialPlayerInfo":         func() interface{} { return &FasterRunFinialPlayerInfo{} },
	"outer.FasterRunDealNtf":                  func() interface{} { return &FasterRunDealNtf{} },
	"outer.FasterRunCardsGroup":               func() interface{} { return &FasterRunCardsGroup{} },
	"outer.FailRsp":                           func() interface{} { return &FailRsp{} },
	"outer.Exchange3Info":                     func() interface{} { return &Exchange3Info{} },
	"outer.EnterGameRsp":                      func() interface{} { return &EnterGameRsp{} },
	"outer.EnterGameReq":                      func() interface{} { return &EnterGameReq{} },
	"outer.DisbandRoomRsp":                    func() interface{} { return &DisbandRoomRsp{} },
	"outer.DisbandRoomReq":                    func() interface{} { return &DisbandRoomReq{} },
	"outer.DisbandAllianceRsp":                func() interface{} { return &DisbandAllianceRsp{} },
	"outer.DisbandAllianceReq":                func() interface{} { return &DisbandAllianceReq{} },
	"outer.DeleteMailReq":                     func() interface{} { return &DeleteMailReq{} },
	"outer.CreateRoomRsp":                     func() interface{} { return &CreateRoomRsp{} },
	"outer.CreateRoomReq":                     func() interface{} { return &CreateRoomReq{} },
	"outer.ClaimRebateScoreRsp":               func() interface{} { return &ClaimRebateScoreRsp{} },
	"outer.ClaimRebateScoreReq":               func() interface{} { return &ClaimRebateScoreReq{} },
	"outer.CardsOfBTE":                        func() interface{} { return &CardsOfBTE{} },
	"outer.BombsWinScore":                     func() interface{} { return &BombsWinScore{} },
	"outer.BindPhoneRsp":                      func() interface{} { return &BindPhoneRsp{} },
	"outer.BindPhoneReq":                      func() interface{} { return &BindPhoneReq{} },
	"outer.Base":                              func() interface{} { return &Base{} },
	"outer.AllianceInfoNtf":                   func() interface{} { return &AllianceInfoNtf{} },
	"outer.AgentRebateInfoRsp":                func() interface{} { return &AgentRebateInfoRsp{} },
	"outer.AgentRebateInfoReq":                func() interface{} { return &AgentRebateInfoReq{} },
	"outer.AgentRebateDetailInfoRsp":          func() interface{} { return &AgentRebateDetailInfoRsp{} },
	"outer.AgentRebateDetailInfoReq":          func() interface{} { return &AgentRebateDetailInfoReq{} },
	"outer.AgentMembersRsp":                   func() interface{} { return &AgentMembersRsp{} },
	"outer.AgentMembersReq":                   func() interface{} { return &AgentMembersReq{} },
	"outer.AgentDownRebateInfo":               func() interface{} { return &AgentDownRebateInfo{} },
	"outer.AgentDownDailyStatRsp":             func() interface{} { return &AgentDownDailyStatRsp{} },
	"outer.AgentDownDailyStatReq":             func() interface{} { return &AgentDownDailyStatReq{} },
	"outer.AddMailNotify":                     func() interface{} { return &AddMailNotify{} },
}
