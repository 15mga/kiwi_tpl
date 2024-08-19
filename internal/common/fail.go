// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package common

import (
	"github.com/15mga/kiwi/util"
)

// gate 1
const (
	//id不存在
	EcGateNotExistId = 1000
	//地址不存在
	EcGateNotExistAddr = 1001
)

// user 2
const (
	//手机号错误
	EcMobileWrong = 2000
	//密码错误
	EcPasswordWrong = 2001
	//手机号已存在
	EcMobileExist = 2002
	//手机号已存在
	EcNickExist = 2003
	//手机验证码错误
	EcSmsWrong = 2004
	//账号或密码错误
	EcWrongMobileOrPassword = 2005
	//邮件验证码错误
	EcEmailCodeWrong = 2006
	//邮箱已存在
	EcEmailExist = 2007
	//账号或密码错误
	EcWrongEmailOrPassword = 2008
	//获取微信token失败
	EcGetWechatTokenFailed = 2009
	//获取微信用户信息失败
	EcGetWechatUserInfoFailed = 2010
	//昵称不能为空
	EcNickCanNotEmpty = 2011
	//没有登录
	EcNotSignIn = 2012
	//用户不存在
	EcUserNotExist = 2013
	//手机号不存在
	EcMobileNotExist = 2014
	//邮箱不存在
	EcEmailNotExist = 2015
	//无效的token
	EcInvalidToken = 2016
	//用户已存在
	EcUserExist = 2017
	//用户被禁用
	EcUserBanned = 2018
)

func init() {
	util.SetErrCodesToStrMap(map[util.TErrCode]string{
		1000: "gate not exist id",
		1001: "gate not exist addr",
		2000: "mobile wrong",
		2001: "password wrong",
		2002: "mobile exist",
		2003: "nick exist",
		2004: "sms wrong",
		2005: "wrong mobile or password",
		2006: "email code wrong",
		2007: "email exist",
		2008: "wrong email or password",
		2009: "get wechat token failed",
		2010: "get wechat user info failed",
		2011: "nick can not empty",
		2012: "not sign in",
		2013: "user not exist",
		2014: "mobile not exist",
		2015: "email not exist",
		2016: "invalid token",
		2017: "user exist",
		2018: "user banned",
	})
}
