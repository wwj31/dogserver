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

// 帮助
var _ = reg("help", func(arg ...string) {
	for name, _ := range cmds {
		fmt.Println(name)
	}
})

// 获得清单列表
var _ = reg("manifest", func(arg ...string) {
	rsp := client.Req(&outer.RoomManifestListReq{
		GameType: outer.GameType_FasterRun,
	})
	_ = rsp
})

// 创建房间
var _ = reg("c", func(arg ...string) {
	if len(arg) != 1 {
		return
	}

	req := &outer.SetRoomManifestReq{
		GameType: outer.GameType_FasterRun,
		GameParams: &outer.GameParams{
			MaintainEmptyRoom: cast.ToInt32(arg[0]),
			FasterRun: &outer.FasterRunParams{
				PlayCountLimit:           4,
				BaseScore:                2,
				PlayerNumber:             0,
				CardsNumber:              0,
				DecideMasterType:         0,
				FirstSpades3:             true,
				ShowSpareNumber:          true,
				DoubleHeartsTen:          true,
				PlayTolerance:            true,
				FollowPlayTolerance:      true,
				SpareOnlyOneWithoutLose:  true,
				AgainstSpring:            true,
				SpecialThreeCards:        true,
				SpecialThreeCardsWithOne: true,
				AAAIsBombs:               true,
				AllowScoreSmallZero:      true,
				BigWinner:                true,
				ReBate: &outer.RebateParams{
					RangeL1: &outer.RangeParams{
						Valid:            false,
						Min:              0,
						Max:              100,
						RebateRatio:      15,
						MinimumGuarantee: 5,
						MinimumRebate:    3,
					},
				},
			},
		},
	}

	client.Req(req)
})

// 房间列表
var _ = reg("l", func(arg ...string) {
	req := &outer.RoomListReq{}

	v := client.Req(req)
	if rsp, ok := v.(*outer.RoomListRsp); ok {
		for _, info := range rsp.RoomList {
			log.Infof("roominfo %v", info.String())
		}
	}
})

var _ = reg("agent", func(arg ...string) {
	req := &outer.AgentMembersReq{}
	v := client.Req(req)
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

	rsp := client.Req(req)
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

	client.Req(req)
})

// 退出房间
var _ = reg("q", func(arg ...string) {
	req := &outer.LeaveRoomReq{}

	client.Req(req)
})

// 出牌
var _ = reg("out", func(arg ...string) {
	if len(arg) != 1 {
		log.Warnw("out arg len err", len(arg))
		return
	}
	req := &outer.MahjongBTEPlayCardReq{Index: cast.ToInt32(arg[0])}
	client.Req(req)
})

// 碰
var _ = reg("guo", func(arg ...string) {
	req := &outer.MahjongBTEOperateReq{ActionType: outer.ActionType_ActionPass}
	client.Req(req)
})

// 碰
var _ = reg("peng", func(arg ...string) {
	req := &outer.MahjongBTEOperateReq{ActionType: outer.ActionType_ActionPong}
	client.Req(req)
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
	client.Req(req)
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
	client.Req(req)
})

// 牛牛 抢庄
var _ = reg("master", func(arg ...string) {
	if len(arg) != 1 {
		return
	}

	req := &outer.NiuNiuToBeMasterReq{
		Times: cast.ToInt32(arg[0]),
	}
	client.Req(req)
})

// 牛牛 押注
var _ = reg("bet", func(arg ...string) {
	if len(arg) != 1 {
		return
	}

	req := &outer.NiuNiuToBettingReq{
		Gold: cast.ToFloat32(arg[0]),
	}
	client.Req(req)
})
