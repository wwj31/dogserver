package mahjong

import (
	"time"

	"github.com/wwj31/dogactor/tools"

	"server/common/log"
	"server/proto/outermsg/outer"
)

// 游戏状态
const (
	pongGangHuGuoExpire = 10 * time.Second // 碰、杠、胡、过持续时间
	playCardExpire      = 15 * time.Second // 摸牌后的行为持续时间(出牌，杠，胡)
)

type checkCardType int32

const (
	drawCardType checkCardType = 1 // 摸牌
	playCardType checkCardType = 2 // 打牌
	GangType1    checkCardType = 3 // 明杠,自己摸牌，杠碰的牌(可抢杠胡)
	GangType2    checkCardType = 4 // 明杠,补杠，别人打牌出我碰的牌
	GangType3    checkCardType = 4 // 明杠,直杠，别人打牌出我手牌里有三张
	GangType4    checkCardType = 5 // 暗杠,自己摸牌，自己手牌有三张
)

type (
	peerCard struct {
		typ  checkCardType
		card Card
		seat int
	}
)

type StatePlaying struct {
	*Mahjong
	actionTimerId string
	peerCards     []peerCard // 每次操作追加操作记录
}

func (s *StatePlaying) State() int {
	return Playing
}

func (s *StatePlaying) Enter() {
	s.peerCards = make([]peerCard, 0)
	s.actionMap = make(map[int]*action)
	s.actionTimerId = ""
	log.Infow("[Mahjong] enter state playing", "room", s.room.RoomId)
	s.drawCard(s.masterIndex)
}

func (s *StatePlaying) Leave() {
	s.cancelActionTimer()
	log.Infow("[Mahjong] leave state playing", "room", s.room.RoomId)
}

func (s *StatePlaying) Handle(shortId int64, v any) (result any) {
	player, seatIndex, err := s.getPlayerAndSeatE(shortId)
	if err != outer.ERROR_OK {
		return err
	}

	switch msg := v.(type) {
	case *outer.MahjongBTEPlayCardReq: // 打牌
		if msg.Index < 0 || int(msg.Index) >= player.handCards.Len() {
			return outer.ERROR_MSG_REQ_PARAM_INVALID
		}

		// 进入打牌逻辑
		if ok, errCode := s.playCard(int(msg.Index), seatIndex); !ok {
			return errCode
		}
		return &outer.MahjongBTEPlayCardRsp{AllCards: player.allCardsToPB()}

	case *outer.MahjongBTEOperateReq: // 碰、杠、胡、过
		if ok, errCode := s.operate(player, seatIndex, msg.ActionType, Card(msg.Gang)); !ok {
			return errCode
		}
		return &outer.MahjongBTEOperateRsp{AllCards: player.allCardsToPB()}

	}
	return nil
}

func (s *StatePlaying) getPlayerAndSeatE(shortId int64) (*mahjongPlayer, int, outer.ERROR) {
	player, seatIndex := s.findMahjongPlayer(shortId)
	if player == nil {
		return nil, -1, outer.ERROR_PLAYER_NOT_IN_ROOM
	}

	if _, ok := s.actionMap[seatIndex]; !ok {
		return nil, -1, outer.ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH
	}
	return player, seatIndex, outer.ERROR_OK
}

// 取消行动倒计时
func (s *StatePlaying) cancelActionTimer() {
	s.room.CancelTimer(s.actionTimerId)
}

// 行动倒计时
func (s *StatePlaying) actionTimer(expireAt time.Time) {
	s.cancelActionTimer()
	s.currentActionEndAt = expireAt
	s.actionTimerId = s.room.AddTimer(tools.XUID(), s.currentActionEndAt, func(dt time.Duration) {
		if len(s.actionMap) == 0 {
			// 所有行动计时器都会被正常释放，无行动者超时属于异常
			log.Warnw("action timer timeout len == 0")
			return
		}

		for seatIndex, act := range s.actionMap {
			player := s.mahjongPlayers[seatIndex]
			// 出牌人，只可能有一个行动者
			if act.isValidAction(outer.ActionType_ActionPlayCard) {
				var defaultPlayCard Card

				// 优先打定缺花色,没有定缺花色的牌，就选手牌
				ignoreCards := player.handCards.colorCards(player.ignoreColor)
				if ignoreCards.Len() != 0 {
					defaultPlayCard = ignoreCards.Random()
				} else {
					defaultPlayCard = player.handCards.Random()
				}

				// 把超时后随机选的牌打出去
				playIndex := player.handCards.CardIndex(defaultPlayCard)
				if playIndex == -1 {
					log.Errorw("action timeout, playIndex == -1",
						"roomId", s.room.RoomId, "player", player.ShortId, "act", act,
						"hand", player.handCards, "play", defaultPlayCard)
					return
				}
				s.playCard(playIndex, seatIndex)
				break
			} else {
				var (
					defaultOperaType outer.ActionType
					card             Card
				)

				// (碰杠胡过)行动者，优先打胡->杠->碰
				if act.isValidAction(outer.ActionType_ActionHu) {
					defaultOperaType = outer.ActionType_ActionHu
				} else if act.isValidAction(outer.ActionType_ActionGang) {
					defaultOperaType = outer.ActionType_ActionGang
					card = Card(act.currentGang[0])
				} else if act.isValidAction(outer.ActionType_ActionPong) {
					defaultOperaType = outer.ActionType_ActionPong
					card = s.cards[s.cards.Len()-1]
				} else {
					log.Warnw("action exception",
						"roomId", s.room.RoomId, "player", seatIndex, "act", act)
					continue
				}

				s.operate(player, seatIndex, defaultOperaType, card)
			}
		}
	})
}

func (s *StatePlaying) AppendPeerCard(typ checkCardType, card Card, seat int) {
	s.peerCards = append(s.peerCards, peerCard{
		typ:  typ,
		card: card,
		seat: seat,
	})
}
