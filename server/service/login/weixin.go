package login

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"server/common/log"
	"server/proto/innermsg/inner"
)

const (
	accessTokenURL    = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code"
	refreshTokenURL   = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%v&grant_type=refresh_token&refresh_token=%v"
	weChatUserInfoURL = "https://api.weixin.qq.com/sns/userinfo?access_token=%v&openid=%v"

	appid  = "fewfjewiofewjifw"
	secret = "jfeohewouiofdgherio"
)

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
		log.Errorw("WeiXin access token get failed", "err", err)
		return nil
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Errorw("WeiXin access token body read failed", "err", err)
		return nil
	}

	var info WeChatAccessInfo
	err = json.Unmarshal(bodyBytes, &info)
	if err != nil {
		log.Errorw("WeiXin access token body json unmarshal failed", "err", err, "bodyBytes", string(bodyBytes))
		return nil
	}
	return &info
}

// RefreshAccessTokenExpiration 刷新access_token时限, 返回refresh_token是否失效，失效需要重新登录
func (w *WeChatAccessInfo) RefreshAccessTokenExpiration() (refreshExpire bool) {
	addr := fmt.Sprintf(refreshTokenURL, appid, w.RefreshToken)
	rsp, err := http.Get(addr)
	if err != nil {
		log.Errorw("WeiXin refresh token get failed", "err", err)
		return false
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Errorw("WeiXin refresh token body read failed", "err", err)
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
		log.Errorw("WeiXin refresh token get failed", "err", err)
		return
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Errorw("WeiXin refresh token body read failed", "err", err)
		return
	}

	var rspInfo = struct {
		AccessToken  string   `json:"nickname"`
		ExpiresIn    int64    `json:"sex"`
		RefreshToken string   `json:"province"`
		OpenId       string   `json:"city"`
		Scope        string   `json:"country"`
		Scope        string   `json:"headimgurl"`
		Privilege    []string `json:"privilege"`
		Scope        string   `json:"unionid"`
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
