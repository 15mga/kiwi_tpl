// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package codec

import (
	"github.com/15mga/kiwi"
)

const (
	//注册
	UserSignUpWithMobileReq kiwi.TMethod = 0
	//注册
	UserSignUpWithMobileRes kiwi.TMethod = 1
	//登录
	UserSignInWithMobileReq kiwi.TMethod = 2
	//登录
	UserSignInWithMobileRes kiwi.TMethod = 3
	//重置密码
	UserResetPasswordWithMobileReq kiwi.TMethod = 4
	//重置密码
	UserResetPasswordWithMobileRes kiwi.TMethod = 5
	//短信验证码
	UserCodeWithMobileReq kiwi.TMethod = 6
	UserCodeWithMobileRes kiwi.TMethod = 7
	//注册
	UserSignUpWithEmailReq kiwi.TMethod = 8
	//注册
	UserSignUpWithEmailRes kiwi.TMethod = 9
	//登录
	UserSignInWithEmailReq kiwi.TMethod = 10
	//登录
	UserSignInWithEmailRes kiwi.TMethod = 11
	//重置密码
	UserResetPasswordWithEmailReq kiwi.TMethod = 12
	//重置密码
	UserResetPasswordWithEmailRes kiwi.TMethod = 13
	//短信验证码
	UserCodeWithEmailReq kiwi.TMethod = 14
	UserCodeWithEmailRes kiwi.TMethod = 15
	//登录
	UserSignInWithWechatReq kiwi.TMethod = 16
	//登录
	UserSignInWithWechatRes kiwi.TMethod = 17
	UserNewReq              kiwi.TMethod = 18
	UserNewRes              kiwi.TMethod = 19
	UserSignInReq           kiwi.TMethod = 20
	UserSignInRes           kiwi.TMethod = 21
	//登出
	UserSignOutReq      kiwi.TMethod = 22
	UserSignOutRes      kiwi.TMethod = 23
	UserReconnectReq    kiwi.TMethod = 24
	UserReconnectRes    kiwi.TMethod = 25
	UserRepeatSignInPus kiwi.TMethod = 26
	UserDisconnectReq   kiwi.TMethod = 500
	UserDisconnectRes   kiwi.TMethod = 501
	UserDisconnectedNtc kiwi.TMethod = 502
	UserConnectedNtc    kiwi.TMethod = 503
	//更新用户head cache
	UserUpdateHeadReq kiwi.TMethod = 504
	UserUpdateHeadRes kiwi.TMethod = 505
)
