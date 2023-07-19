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
var _ = reg("c", func(arg ...string) {
	req := &outer.CreateRoomReq{
		GameType: outer.GameType_Mahjong,
		GameParams: &outer.GameParams{
			Mahjong: &outer.MahjongParams{
				ZiMoJia:           1,
				DianGangHua:       0,
				HuanSanZhang:      1,
				YaoJiuDui:         true,
				MenQingZhongZhang: true,
				TianDiHu:          true,
				DianPaoPingHu:     true,
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
		log.Infof("joinRoomRsp gambling info:%v", info)
	}
})

// 准备
var _ = reg("r", func(arg ...string) {
	ready := true
	if len(arg) > 0 {
		ready = cast.ToBool(arg[0])
	}
	req := &outer.ReadyReq{Ready: ready}

	client.Req(outer.Msg_IdReadyReq, req)
})

// 退出房间
var _ = reg("q", func(arg ...string) {
	req := &outer.LeaveRoomReq{}

	client.Req(outer.Msg_IdLeaveRoomReq, req)
})
