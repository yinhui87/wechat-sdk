package wechat_mp_sdk

import (
	"github.com/yinhui87/wechat-sdk"
)

type WechatOauthToken struct {
	*wechatsdk.WechatError
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint   `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

type WechatOauthUserInfo struct {
	*wechatsdk.WechatError
	OpenId     string   `json:"openid"`
	NickName   string   `json:"nickname"`
	Sex        uint     `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionId    string   `json:"unionid"`
}
