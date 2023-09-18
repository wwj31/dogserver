package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cast"

	"server/common/log"
	"server/proto/outermsg/outer"
)

var cmds = map[string]func(arg ...string){}
var client *Client

func Run(c *Client) {
	client = c
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入:")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		input = strings.Replace(input, "\r", "", -1)
		str := strings.Split(input, " ")
		if len(str) == 0 {
			fmt.Println("无效命令", input)
			continue
		}
		cmd := str[0]
		fn, ok := cmds[cmd]
		if !ok {
			fmt.Println("找不到命令", input)
			continue
		}

		arg := str[1:]
		fn(arg...)
	}
}

func reg(name string, fn func(arg ...string)) bool {
	cmds[name] = fn
	return true
}

// 创建房间
var _ = reg("help", func(arg ...string) {
	for name, _ := range cmds {
		fmt.Println(name)
	}
})

// 创建房间
var _ = reg("c", func(arg ...string) {
	req := &outer.CreateRoomReq{
		GameType: outer.GameType_FasterRun,
		GameParams: &outer.GameParams{
			FasterRun: &outer.FasterRunParams{
				PlayCountLimit:           4,
				BaseScore:                2,
				PlayerNumber:             0,
				CardsNumber:              0,
				DecideMasterType:         0,
				FirstSpades3:             false,
				ShowSpareNumber:          false,
				DoubleHeartsTen:          false,
				PlayTolerance:            false,
				FollowPlayTolerance:      false,
				SpareOnlyOneWithoutLose:  false,
				AgainstSpring:            false,
				SpecialThreeCards:        false,
				SpecialThreeCardsWithOne: false,
				AAAIsBombs:               false,
				AllowScoreSmallZero:      false,
				BigWinner:                false,
				ReBate:                   nil,
			},
		},
	}

	client.Req(outer.Msg_IdCreateRoomReq, req)
})

// 房间列表
var _ = reg("l", func(arg ...string) {
	req := &outer.RoomListReq{}

	v := client.Req(outer.Msg_IdRoomListReq, req)
	if rsp, ok := v.(*outer.RoomListRsp); ok {
		for _, info := range rsp.RoomList {
			log.Infof("roominfo %v", info.String())
		}
	}
})

// 房间列表
var _ = reg("agent", func(arg ...string) {
	req := &outer.AgentMembersReq{}
	v := client.Req(outer.Msg_IdAgentMembersReq, req)
	if rsp, ok := v.(*outer.AgentMembersRsp); ok {
		for _, member := range rsp.DownMembers {
			log.Infof("roominfo %v", *member)
		}
	}
})

// 加入房间
var _ = reg("j", func(arg ...string) {
	if len(arg) != 1 {
		return
	}

	roomId := cast.ToInt64(arg[0])
	req := &outer.JoinRoomReq{
		RoomId: roomId,
	}

	rsp := client.Req(outer.Msg_IdJoinRoomReq, req)
	joinRsp, ok := rsp.(*outer.JoinRoomRsp)
	if ok {
		info := outer.MahjongBTEGameInfo{}
		_ = proto.Unmarshal(joinRsp.GamblingData, &info)
		log.Infof("joinRoomRsp gambling info:%v", &info)
	}
})

// 准备
var _ = reg("r", func(arg ...string) {
	ready := true
	if len(arg) > 0 {
		ready = cast.ToBool(arg[0])
	}
	req := &outer.MahjongBTEReadyReq{Ready: ready}

	client.Req(outer.Msg_IdFasterRunReadyReq, req)
})

// 退出房间
var _ = reg("q", func(arg ...string) {
	req := &outer.LeaveRoomReq{}

	client.Req(outer.Msg_IdLeaveRoomReq, req)
})

// 出牌
var _ = reg("out", func(arg ...string) {
	if len(arg) != 1 {
		log.Warnw("out arg len err", len(arg))
		return
	}
	req := &outer.MahjongBTEPlayCardReq{Index: cast.ToInt32(arg[0])}
	client.Req(outer.Msg_IdMahjongBTEPlayCardReq, req)
})

// 碰
var _ = reg("guo", func(arg ...string) {
	req := &outer.MahjongBTEOperateReq{ActionType: outer.ActionType_ActionPass}
	client.Req(outer.Msg_IdMahjongBTEOperateReq, req)
})

// 碰
var _ = reg("peng", func(arg ...string) {
	req := &outer.MahjongBTEOperateReq{ActionType: outer.ActionType_ActionPong}
	client.Req(outer.Msg_IdMahjongBTEOperateReq, req)
})

// 杠
var _ = reg("gang", func(arg ...string) {
	if len(arg) != 1 {
		return
	}

	req := &outer.MahjongBTEOperateReq{
		ActionType: outer.ActionType_ActionGang,
		Gang:       cast.ToInt32(arg[0]),
	}
	client.Req(outer.Msg_IdMahjongBTEOperateReq, req)
})

// hu
var _ = reg("hu", func(arg ...string) {
	if len(arg) != 1 {
		return
	}

	req := &outer.MahjongBTEOperateReq{
		ActionType: outer.ActionType_ActionHu,
		Hu:         outer.HuType(cast.ToInt32(arg[0])),
	}
	client.Req(outer.Msg_IdMahjongBTEOperateReq, req)
})
