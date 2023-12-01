package login

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cast"

	"server/common/log"
	"server/proto/innermsg/inner"
)

const (
	accessTokenURL    = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code"
	refreshTokenURL   = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%v&grant_type=refresh_token&refresh_token=%v"
	weChatUserInfoURL = "https://api.weixin.qq.com/sns/userinfo?access_token=%v&openid=%v"

	appid  = "wx3483e9b50015f7e1"
	secret = "14a391e492ee7c532c4ac2b0d6fbc176"
)

// https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html

type WeChatAccessInfo struct {
	AccessToken  string `json:"access_token"`  // 接口调用凭证
	ExpiresIn    int64  `json:"expires_in"`    // AccessToken过期时间 单位(秒)
	RefreshToken string `json:"refresh_token"` // 刷新AccessToken用
	OpenId       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // snsapi_userinfo
	UnionId      string `json:"unionid"`       // o6_bmasdasdsad6_2sgVt7hMZOPfL
}

// WeChatAccessToken 获取微信用户接口调用凭证
func WeChatAccessToken(code string) *WeChatAccessInfo {
	addr := fmt.Sprintf(accessTokenURL, appid, secret, code)
	rsp, err := http.Get(addr)
	if err != nil {
		log.Errorw("WeiXin access token get failed", "err", err, "code", code)
		return nil
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Errorw("WeiXin access token body read failed", "err", err, "code", code)
		return nil
	}

	var info WeChatAccessInfo
	err = json.Unmarshal(bodyBytes, &info)
	if err != nil {
		log.Errorw("WeiXin access token body json unmarshal failed", "err", err, "code", code, "bodyBytes", string(bodyBytes))
		return nil
	}
	return &info
}

// RefreshAccessTokenExpiration 刷新access_token时限, 返回refresh_token是否失效，失效需要重新登录
func (w *WeChatAccessInfo) RefreshAccessTokenExpiration() (refreshExpire bool) {
	addr := fmt.Sprintf(refreshTokenURL, appid, w.RefreshToken)
	rsp, err := http.Get(addr)
	if err != nil {
		log.Errorw("WeiXin refresh token get failed", "err", err, "info", w)
		return false
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Errorw("WeiXin refresh token body read failed", "err", err, "info", w)
		return false
	}

	var rspInfo = struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		OpenId       string `json:"openid"`
		Scope        string `json:"scope"`
	}{}

	err = json.Unmarshal(bodyBytes, &rspInfo)
	if err != nil {
		log.Errorw("WeiXin access token body json unmarshal failed", "err", err, "bodyBytes", string(bodyBytes))
		return true
	}

	w.AccessToken = rspInfo.AccessToken
	w.ExpiresIn = rspInfo.ExpiresIn
	w.OpenId = rspInfo.OpenId
	return false
}

func (w *WeChatAccessInfo) UserInfo() (userInfo *inner.WeChatUserInfo) {
	addr := fmt.Sprintf(weChatUserInfoURL, w.AccessToken, w.OpenId)
	rsp, err := http.Get(addr)
	if err != nil {
		log.Errorw("WeiXin UserInfo get failed", "err", err)
		return
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Errorw("WeiXin UserInfo body read failed", "err", err)
		return
	}

	var rspInfo = struct {
		OpenId     string   `json:"openid"`
		NickName   string   `json:"nickname"`
		Sex        int64    `json:"sex"`
		Province   string   `json:"province"`   // 普通用户个人资料填写的省份
		City       string   `json:"city"`       // 普通用户个人资料填写的城市
		Country    string   `json:"country"`    // 国家，如中国为 CN
		HeadImgURL string   `json:"headimgurl"` // 用户头像，最后一个数值代表正方形头像大小（有 0、46、64、96、132 数值可选，0 代表 640*640 正方形头像），用户没有头像时该项为空
		Privilege  []string `json:"privilege"`  // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
		UnionId    string   `json:"unionid"`    // 用户统一标识。针对一个微信开放平台账号下的应用，同一用户的 unionid 是唯一的。
	}{}

	err = json.Unmarshal(bodyBytes, &rspInfo)
	if err != nil {
		log.Errorw("WeiXin UserInfo body json unmarshal failed", "err", err, "bodyBytes", string(bodyBytes))
		return
	}

	return &inner.WeChatUserInfo{
		Icon:   rspInfo.HeadImgURL,
		Gender: cast.ToInt32(rspInfo.Sex) - 1,
		Name:   rspInfo.NickName,
	}
}
