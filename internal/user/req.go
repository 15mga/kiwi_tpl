package user

import (
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/sid"
	"github.com/15mga/kiwi/util"
	"time"
)

func (s *Svc) Start() {
	s.svc.Start()
	initModels()
}

func (s *Svc) OnUserSignUpWithMobile(pkt kiwi.IRcvRequest, req *pb.UserSignUpWithMobileReq, res *pb.UserSignUpWithMobileRes) {
	ok := CheckMobileCode(req.Mobile, req.Code)
	if !ok {
		pkt.Fail(common.EcSmsWrong)
		return
	}

	_, err := InsertMobileAccount(&pb.MobileAccount{
		Id:       req.Mobile,
		Password: common.SaltPw(req.Password),
	})
	if err != nil {
		pkt.Fail(common.EcMobileExist)
		return
	}
	pkt.Ok(res)
}

func (s *Svc) OnUserSignInWithMobile(pkt kiwi.IRcvRequest, req *pb.UserSignInWithMobileReq, res *pb.UserSignInWithMobileRes) {
	account, err := GetMobileAccountWithId(req.Mobile)
	if err != nil || account.Password != common.SaltPw(req.Password) {
		pkt.Fail(common.EcWrongMobileOrPassword)
		return
	}

	userId := account.UserId
	head := util.M{
		common.HdSignInCh:  common.SignInMobile,
		common.HdAccountId: account.Id,
		common.HdMask:      util.GenMask(common.RPlayer),
	}
	if userId == "" {
		userId = account.Id
	} else {
		res.ExistUser = true
	}
	head[common.HdUserId] = userId

	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateUpdateReq{
		Id:   pkt.HeadId(),
		Head: headBytes,
	}, nil, nil)

	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	res.Token, _ = common.GenToken(common.SignInMobile, addr, userId, time.Now().Unix())
	pkt.Ok(res)
}

func (s *Svc) OnUserResetPasswordWithMobile(pkt kiwi.IRcvRequest, req *pb.UserResetPasswordWithMobileReq, res *pb.UserResetPasswordWithMobileRes) {
	ok := CheckMobileCode(req.Mobile, req.Code)
	if !ok {
		pkt.Fail(common.EcSmsWrong)
		return
	}
	account, err := GetMobileAccountWithId(req.Mobile)
	if err != nil {
		pkt.Fail(common.EcMobileNotExist)
		return
	}
	if account.Password != common.SaltPw(req.OldPassword) {
		pkt.Fail(common.EcPasswordWrong)
		return
	}
	account.SetPassword(common.SaltPw(req.NewPassword))
	pkt.Ok(res)
}

func (s *Svc) OnUserCodeWithMobile(pkt kiwi.IRcvRequest, req *pb.UserCodeWithMobileReq, res *pb.UserCodeWithMobileRes) {
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	ok := CheckSendCodeAddr(addr)
	if !ok {
		pkt.Fail(util.EcIllegalOp)
		return
	}
	code, err := SendMobileCode(req.Mobile)
	if err != nil {
		pkt.Fail(err.Code())
		return
	}
	res.Code = code
	pkt.Ok(res)
}

func (s *Svc) OnUserSignUpWithEmail(pkt kiwi.IRcvRequest, req *pb.UserSignUpWithEmailReq, res *pb.UserSignUpWithEmailRes) {
	ok := CheckEmailCode(req.Email, req.Code)
	if !ok {
		pkt.Fail(common.EcSmsWrong)
		return
	}
	_, err := InsertEmailAccount(&pb.EmailAccount{
		Id:       req.Email,
		Password: common.SaltPw(req.Password),
	})
	if err != nil {
		pkt.Fail(common.EcEmailExist)
		return
	}
	pkt.Ok(res)
}

func (s *Svc) OnUserSignInWithEmail(pkt kiwi.IRcvRequest, req *pb.UserSignInWithEmailReq, res *pb.UserSignInWithEmailRes) {
	account, err := GetEmailAccountWithId(req.Email)
	if err != nil || account.Password != common.SaltPw(req.Password) {
		pkt.Fail(common.EcWrongMobileOrPassword)
		return
	}

	userId := account.UserId
	head := util.M{
		common.HdSignInCh:  common.SignInEmail,
		common.HdAccountId: account.Id,
		common.HdMask:      util.GenMask(common.RPlayer),
	}
	if account.UserId != "" {
		res.ExistUser = true
		userId = account.Id
	}
	head[common.HdUserId] = userId

	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateUpdateReq{
		Id:   pkt.HeadId(),
		Head: headBytes,
	}, nil, nil)

	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	res.Token, _ = common.GenToken(common.SignInEmail, addr, userId, time.Now().Unix())
	pkt.Ok(res)
}

func (s *Svc) OnUserResetPasswordWithEmail(pkt kiwi.IRcvRequest, req *pb.UserResetPasswordWithEmailReq, res *pb.UserResetPasswordWithEmailRes) {
	ok := CheckMobileCode(req.Email, req.Code)
	if !ok {
		pkt.Fail(common.EcEmailCodeWrong)
		return
	}
	account, err := GetEmailAccountWithId(req.Email)
	if err != nil {
		pkt.Fail(common.EcEmailNotExist)
		return
	}
	if account.Password != common.SaltPw(req.OldPassword) {
		pkt.Fail(common.EcPasswordWrong)
		return
	}
	account.SetPassword(common.SaltPw(req.NewPassword))
	pkt.Ok(res)
}

func (s *Svc) OnUserCodeWithEmail(pkt kiwi.IRcvRequest, req *pb.UserCodeWithEmailReq, res *pb.UserCodeWithEmailRes) {
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	ok := CheckSendCodeAddr(addr)
	if !ok {
		pkt.Fail(util.EcIllegalOp)
		return
	}
	code, err := SendEmailCode(req.Email)
	if err != nil {
		pkt.Fail(err.Code())
		return
	}
	res.Code = code
	pkt.Ok(res)
}

func (s *Svc) OnUserSignInWithWechat(pkt kiwi.IRcvRequest, req *pb.UserSignInWithWechatReq, res *pb.UserSignInWithWechatRes) {
	// 使用code获取access token
	var token WechatAccessToken
	e := GetWechatAccessToken(req.Code, &token)
	if e != nil {
		pkt.Fail(common.EcGetWechatTokenFailed)
		return
	}

	// 使用access token获取用户信息
	var userInfo WechatUserInfo
	e = GetWechatUserInfo(&token, &userInfo)
	if e != nil {
		pkt.Fail(common.EcGetWechatUserInfoFailed)
		return
	}

	var (
		userId string
		user   *User
	)

	account, err := GetWechatAccountWithId(userInfo.UnionId)
	if err != nil { //没有账号
		_, err = InsertWechatAccount(&pb.WechatAccount{
			Id:     userInfo.UnionId,
			UserId: userId,
		})
		if err != nil {
			pkt.Fail(util.EcDbErr)
			return
		}
	}
	if account.UserId != "" {
		userId = account.UserId
		user, _ = GetUserWithId(userId)
	}
	if user == nil {
		userId = sid.GetStrId()
		account.SetUserId(userId)
		user, err = InsertUser(&pb.User{
			Id:              userId,
			RoleMask:        util.GenMask(common.RPlayer),
			Ban:             false,
			Nick:            userInfo.NickName,
			IdCard:          "",
			RealName:        "",
			CreateTime:      time.Now().Unix(),
			LastSignInTime:  0,
			LastSignInAddr:  "",
			LastOfflineTime: 0,
			LastOs:          "",
			State:           pb.OnlineState_Disconnected,
			Avatar:          userInfo.HeadImgUrl,
			Head:            nil,
			OnlineDur:       0,
		})
		if err != nil {
			pkt.Fail(util.EcDbErr)
			return
		}
	}

	user.SetAvatar(userInfo.HeadImgUrl)

	head := util.M{
		common.HdSignInCh:  common.SignInWechat,
		common.HdAccountId: account.Id,
		common.HdUserId:    userId,
	}
	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateUpdateReq{
		Id:   pkt.HeadId(),
		Head: headBytes,
	}, nil, nil)

	res.ExistUser = true
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	res.Token, _ = common.GenToken(common.SignInWechat, addr, userId, time.Now().Unix())
	pkt.Ok(res)
}

func (s *Svc) OnUserNew(pkt kiwi.IRcvRequest, req *pb.UserNewReq, res *pb.UserNewRes) {
	userId, _ := util.MGet[string](pkt.Head(), common.HdUserId)
	accountId, _ := util.MGet[string](pkt.Head(), common.HdAccountId)
	if userId != accountId {
		pkt.Fail(common.EcUserExist)
		return
	}
	if req.User == nil {
		pkt.Fail(util.EcBadPacket)
		return
	}
	if req.User.Nick == "" {
		pkt.Fail(common.EcNickCanNotEmpty)
		return
	}
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	signIn, _ := util.MGet[string](pkt.Head(), common.HdSignInCh)
	now := time.Now().Unix()
	userId = sid.GetStrId()
	token, err := common.GenToken(signIn, addr, userId, now)
	if err != nil {
		pkt.Fail(util.EcServiceErr)
		return
	}

	var accountIntfc IAccount
	switch signIn {
	case common.SignInMobile:
		account, err := GetMobileAccountWithId(accountId)
		if err != nil {
			pkt.Fail(util.EcServiceErr)
			return
		}
		if account.UserId != "" {
			pkt.Fail(common.EcUserExist)
			return
		}
		accountIntfc = account
	case common.SignInEmail:
		account, err := GetEmailAccountWithId(accountId)
		if err != nil {
			pkt.Fail(util.EcServiceErr)
			return
		}
		if account.UserId != "" {
			pkt.Fail(common.EcUserExist)
			return
		}
		accountIntfc = account
	case common.SignInWechat:
		account, err := GetWechatAccountWithId(accountId)
		if err != nil {
			pkt.Fail(util.EcServiceErr)
			return
		}
		if account.UserId != "" {
			pkt.Fail(common.EcUserExist)
			return
		}
		accountIntfc = account
	default:
		pkt.Fail(util.EcServiceErr)
		return
	}
	_, err = InsertUser(&pb.User{
		Id:              userId,
		RoleMask:        util.GenMask(common.RPlayer),
		Ban:             false,
		Nick:            req.User.Nick,
		IdCard:          "",
		RealName:        "",
		CreateTime:      time.Now().Unix(),
		LastSignInTime:  0,
		LastSignInAddr:  "",
		LastOfflineTime: 0,
		LastOs:          "",
		State:           pb.OnlineState_Disconnected,
		Avatar:          req.User.Avatar,
		Token:           token,
		Head:            nil,
		OnlineDur:       0,
	})
	if err != nil {
		pkt.Fail(common.EcNickExist)
		return
	}
	accountIntfc.SetUserId(userId)

	head := util.M{
		common.HdUserId: userId,
	}
	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateUpdateReq{
		Id:   pkt.HeadId(),
		Head: headBytes,
	}, nil, nil)

	res.Token = token
	pkt.Ok(res)
}

func (s *Svc) OnUserSignIn(pkt kiwi.IRcvRequest, req *pb.UserSignInReq, res *pb.UserSignInRes) {
	userId, _ := util.MGet[string](pkt.Head(), common.HdUserId)
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	signIn, _ := util.MGet[string](pkt.Head(), common.HdSignInCh)
	now := time.Now().Unix()
	token, err := common.GenToken(signIn, addr, userId, now)
	if err != nil {
		pkt.Fail(util.EcServiceErr)
		return
	}
	user, err := GetUserWithId(userId)
	if err != nil {
		pkt.Fail(common.EcUserNotExist)
		return
	}
	if user.Ban {
		pkt.Fail(common.EcUserBanned)
		return
	}

	common.PusUser(s, pkt.Tid(), userId, &pb.UserRepeatSignInPus{})
	pkt.AsyncReq(&pb.GateCloseAddrReq{
		Addr: user.LastSignInAddr,
	}, nil, nil)

	user.SetLastSignInAddr(addr)
	user.SetLastOs(req.Os)
	user.SetLastSignInTime(now)
	user.SetState(pb.OnlineState_Connected)
	user.SetToken(token)

	head := util.M{
		common.HdMask:   user.RoleMask,
		common.HdSignIn: true,
	}
	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateUpdateReq{
		Id:   pkt.HeadId(),
		Head: headBytes,
	}, nil, nil)

	res.User = &pb.User{}
	user.CopyPlayerTag(res.User)
	res.Token = token
	pkt.Ok(res)
}

func (s *Svc) userOffline(userId string) uint16 {
	user, err := GetUserWithId(userId)
	if err != nil {
		return common.EcUserNotExist
	}
	now := time.Now().Unix()
	user.SetState(pb.OnlineState_Disconnected)
	user.SetLastOfflineTime(now)
	user.SetOnlineDur(user.OnlineDur + now - user.LastSignInTime)
	return 0
}

func (s *Svc) OnUserSignOut(pkt kiwi.IRcvRequest, req *pb.UserSignOutReq, res *pb.UserSignOutRes) {
	userId, ok := util.MGet[string](pkt.Head(), common.HdUserId)
	if !ok {
		pkt.Fail(common.EcNotSignIn)
		return
	}
	code := s.userOffline(userId)
	if code > 0 {
		pkt.Fail(code)
		return
	}
	pkt.Ok(res)
}

func (s *Svc) OnUserReconnect(pkt kiwi.IRcvRequest, req *pb.UserReconnectReq, res *pb.UserReconnectRes) {
	claims, err := common.ParseToken(req.Token)
	if err != nil {
		pkt.Fail(common.EcInvalidToken)
		return
	}
	user, err := GetUserWithToken(req.Token)
	if err != nil {
		pkt.Fail(common.EcInvalidToken)
		return
	}
	if claims.UserId != user.Id ||
		claims.Addr != user.LastSignInAddr ||
		claims.Timestamp != user.LastSignInTime {
		pkt.Fail(common.EcInvalidToken)
		return
	}

	pkt.AsyncReq(&pb.GateCloseAddrReq{Addr: claims.Addr}, nil, nil)

	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	tkn, _ := common.GenToken(claims.SignInCh, addr, user.Id, user.LastSignInTime)
	user.SetLastSignInAddr(addr)
	user.SetToken(tkn)

	pkt.AsyncReq(&pb.GateUpdateReq{
		Id:   pkt.HeadId(),
		Head: user.Head,
	}, nil, nil)

	pkt.Ok(res)
}

func (s *Svc) OnUserDisconnect(pkt kiwi.IRcvRequest, req *pb.UserDisconnectReq, res *pb.UserDisconnectRes) {
	code := s.userOffline(req.UserId)
	if code > 0 {
		pkt.Fail(code)
		return
	}
	pkt.Ok(res)
}

func (s *Svc) OnUserUpdateHead(pkt kiwi.IRcvRequest, req *pb.UserUpdateHeadReq, res *pb.UserUpdateHeadRes) {
	head := util.M{}
	head.FromBytes(req.Head)
	user, er := GetUserWithId(req.GetId())
	if er != nil {
		pkt.Fail(common.EcUserNotExist)
		return
	}
	user.SetHead(req.Head)
	pkt.Ok(res)
	kiwi.TI(pkt.Tid(), "user update", util.M{
		"head": head,
	})
}

type IAccount interface {
	SetUserId(id string)
}
