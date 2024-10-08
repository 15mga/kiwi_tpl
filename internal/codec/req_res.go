// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package codec

import (
	"game/internal/common"
	"github.com/15mga/kiwi"
)

func BindReqToRes() {
	kiwi.Codec().BindReqToRes(common.Gate, GateHeartbeatReq, GateHeartbeatRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateSendToIdReq, GateSendToIdRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateSendToAddrReq, GateSendToAddrRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateSendToMultiIdReq, GateSendToMultiIdRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateSendToMultiAddrReq, GateSendToMultiAddrRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateSendToAllReq, GateSendToAllRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateCloseIdReq, GateCloseIdRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateCloseAddrReq, GateCloseAddrRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateUpdateReq, GateUpdateRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateRemoveReq, GateRemoveRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateGetReq, GateGetRes)
	kiwi.Codec().BindReqToRes(common.Gate, GateUpdateRolesReq, GateUpdateRolesRes)
	kiwi.Codec().BindReqToRes(common.User, UserSignUpWithMobileReq, UserSignUpWithMobileRes)
	kiwi.Codec().BindReqToRes(common.User, UserSignInWithMobileReq, UserSignInWithMobileRes)
	kiwi.Codec().BindReqToRes(common.User, UserResetPasswordWithMobileReq, UserResetPasswordWithMobileRes)
	kiwi.Codec().BindReqToRes(common.User, UserCodeWithMobileReq, UserCodeWithMobileRes)
	kiwi.Codec().BindReqToRes(common.User, UserSignUpWithEmailReq, UserSignUpWithEmailRes)
	kiwi.Codec().BindReqToRes(common.User, UserSignInWithEmailReq, UserSignInWithEmailRes)
	kiwi.Codec().BindReqToRes(common.User, UserResetPasswordWithEmailReq, UserResetPasswordWithEmailRes)
	kiwi.Codec().BindReqToRes(common.User, UserCodeWithEmailReq, UserCodeWithEmailRes)
	kiwi.Codec().BindReqToRes(common.User, UserSignInWithWechatReq, UserSignInWithWechatRes)
	kiwi.Codec().BindReqToRes(common.User, UserNewReq, UserNewRes)
	kiwi.Codec().BindReqToRes(common.User, UserSignInReq, UserSignInRes)
	kiwi.Codec().BindReqToRes(common.User, UserSignOutReq, UserSignOutRes)
	kiwi.Codec().BindReqToRes(common.User, UserReconnectReq, UserReconnectRes)
	kiwi.Codec().BindReqToRes(common.User, UserDisconnectReq, UserDisconnectRes)
	kiwi.Codec().BindReqToRes(common.User, UserUpdateHeadReq, UserUpdateHeadRes)
}
