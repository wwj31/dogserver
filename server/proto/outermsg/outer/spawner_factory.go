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
	"outer.ReceiveMailItemRsp":                func() interface{} { return &ReceiveMailItemRsp{} },
	"outer.ReceiveMailItemReq":                func() interface{} { return &ReceiveMailItemReq{} },
	"outer.RebateScoreRsp":                    func() interface{} { return &RebateScoreRsp{} },
	"outer.RebateScoreReq":                    func() interface{} { return &RebateScoreReq{} },
	"outer.RebateParams":                      func() interface{} { return &RebateParams{} },
	"outer.ReadMailRsp":                       func() interface{} { return &ReadMailRsp{} },
	"outer.ReadMailReq":                       func() interface{} { return &ReadMailReq{} },
	"outer.RangeParams":                       func() interface{} { return &RangeParams{} },
	"outer.PlayerInfo":                        func() interface{} { return &PlayerInfo{} },
	"outer.PlayCardTips":                      func() interface{} { return &PlayCardTips{} },
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
	"outer.LatestOperaCard":                   func() interface{} { return &LatestOperaCard{} },
	"outer.KickOutMemberRsp":                  func() interface{} { return &KickOutMemberRsp{} },
	"outer.KickOutMemberReq":                  func() interface{} { return &KickOutMemberReq{} },
	"outer.JoinRoomRsp":                       func() interface{} { return &JoinRoomRsp{} },
	"outer.JoinRoomReq":                       func() interface{} { return &JoinRoomReq{} },
	"outer.InviteAllianceRsp":                 func() interface{} { return &InviteAllianceRsp{} },
	"outer.InviteAllianceReq":                 func() interface{} { return &InviteAllianceReq{} },
	"outer.HeartRsp":                          func() interface{} { return &HeartRsp{} },
	"outer.HeartReq":                          func() interface{} { return &HeartReq{} },
	"outer.GameParams":                        func() interface{} { return &GameParams{} },
	"outer.FasterRunParams":                   func() interface{} { return &FasterRunParams{} },
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
	"outer.BindPhoneRsp":                      func() interface{} { return &BindPhoneRsp{} },
	"outer.BindPhoneReq":                      func() interface{} { return &BindPhoneReq{} },
	"outer.Base":                              func() interface{} { return &Base{} },
	"outer.AllianceInfoNtf":                   func() interface{} { return &AllianceInfoNtf{} },
	"outer.AgentRebateInfoRsp":                func() interface{} { return &AgentRebateInfoRsp{} },
	"outer.AgentRebateInfoReq":                func() interface{} { return &AgentRebateInfoReq{} },
	"outer.AgentMembersRsp":                   func() interface{} { return &AgentMembersRsp{} },
	"outer.AgentMembersReq":                   func() interface{} { return &AgentMembersReq{} },
	"outer.AddMailNotify":                     func() interface{} { return &AddMailNotify{} },
}
