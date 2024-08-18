package user

import (
	"fmt"
	"game/internal/common"
	"github.com/15mga/kiwi/util"
	"io"
	"net/http"
)

type WechatAccessToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	OpenId       string `json:"openid,omitempty"`
	Scope        string `json:"scope,omitempty"`
	UnionId      string `json:"unionid,omitempty"`
}

type WechatUserInfo struct {
	OpenId     string `json:"openid,omitempty"`
	Language   string `json:"language,omitempty"`
	Province   string `json:"province,omitempty"`
	Country    string `json:"country,omitempty"`
	Privileges []any  `json:"privileges,omitempty"`
	UnionId    string `json:"unionid,omitempty"`
	NickName   string `json:"nickname,omitempty"`
	Sex        int    `json:"sex,omitempty"`
	City       string `json:"city,omitempty"`
	HeadImgUrl string `json:"headimgurl,omitempty"`
}

// GetWechatAccessToken 使用code获取access token
func GetWechatAccessToken(code string, token *WechatAccessToken) *util.Err {
	conf := common.Conf().User.Wechat
	reqURL := fmt.Sprintf(conf.Url, conf.AppId, conf.AppSecret, code)
	resp, e := http.Get(reqURL)
	if e != nil {
		return util.WrapErr(util.EcServiceErr, e)
	}
	defer resp.Body.Close()
	body, e := io.ReadAll(resp.Body)
	if e != nil {
		return util.WrapErr(util.EcServiceErr, e)
	}

	return util.JsonUnmarshal(body, &token)
}

// GetWechatUserInfo 使用access token获取用户信息
func GetWechatUserInfo(token *WechatAccessToken, userInfo *WechatUserInfo) *util.Err {
	reqURL := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", token.AccessToken, token.OpenId)
	resp, e := http.Get(reqURL)
	if e != nil {
		return util.WrapErr(util.EcServiceErr, e)
	}
	defer resp.Body.Close()
	body, e := io.ReadAll(resp.Body)
	if e != nil {
		return util.WrapErr(util.EcServiceErr, e)
	}

	return util.JsonUnmarshal(body, &userInfo)
}
