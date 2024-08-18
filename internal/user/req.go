package user

import (
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/sid"
	"github.com/15mga/kiwi/util"
	"github.com/15mga/kiwi/util/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (s *Svc) Start() {
	s.svc.Start()
	initModels()
}

func (s *Svc) BeforeShutdown() {
	StoreAllUsers()
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

	head := util.M{
		common.HdSignIn:    common.SignInMobile,
		common.HdAccountId: req.Mobile,
		common.HdMask:      util.GenMask(common.RPlayer),
	}
	if account.UserId != "" {
		res.ExistUser = true
		head[common.HdUserId] = account.UserId
	}

	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateUpdateReq{
		Id:   pkt.HeadId(),
		Head: headBytes,
	}, func(code uint16) {
		pkt.Fail(code)
	}, func(msg util.IMsg) {
		pkt.Ok(res)
	})
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
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)

	head := util.M{
		common.HdSignIn:    common.SignInEmail,
		common.HdAccountId: req.Email,
		common.HdMask:      util.GenMask(common.RPlayer),
	}
	if account.UserId != "" {
		res.ExistUser = true
		head[common.HdUserId] = account.UserId
	}

	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateAddrUpdateReq{
		Addr: addr,
		Head: headBytes,
	}, func(code uint16) {
		pkt.Fail(code)
	}, func(msg util.IMsg) {
		pkt.Ok(res)
	})
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

	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
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
	}

	user.SetAvatar(userInfo.HeadImgUrl)

	head := util.M{
		common.HdSignIn: common.SignInWechat,
		common.HdUserId: userId,
	}
	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateAddrUpdateReq{
		Addr: addr,
		Head: headBytes,
	}, func(code uint16) {
		pkt.Fail(code)
	}, func(msg util.IMsg) {
		res.ExistUser = true
		pkt.Ok(res)
	})
}

func (s *Svc) OnUserNew(pkt kiwi.IRcvRequest, req *pb.UserNewReq, res *pb.UserNewRes) {
	if req.User.Nick == "" {
		pkt.Fail(common.EcNickCanNotEmpty)
		return
	}
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	accountId, _ := util.MGet[string](pkt.Head(), common.HdAccountId)
	signIn, _ := util.MGet[string](pkt.Head(), common.HdSignIn)
	userId := sid.GetStrId()
	_, err := InsertUser(&pb.User{
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
		Token:           "",
		Head:            nil,
		OnlineDur:       0,
	})
	if err != nil {
		pkt.Fail(util.EcDbErr)
		return
	}

	switch signIn {
	case common.SignInMobile:
		account, err := GetMobileAccountWithId(accountId)
		if err != nil {
			pkt.Fail(util.EcServiceErr)
			return
		}
		account.SetUserId(userId)
	case common.SignInEmail:
		account, err := GetEmailAccountWithId(accountId)
		if err != nil {
			pkt.Fail(util.EcServiceErr)
			return
		}
		account.SetUserId(userId)
	case common.SignInWechat:
		account, err := GetWechatAccountWithId(accountId)
		if err != nil {
			pkt.Fail(util.EcServiceErr)
			return
		}
		account.SetUserId(userId)
	}
	head := util.M{
		common.HdUserId: userId,
	}
	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateAddrUpdateReq{
		Addr: addr,
		Head: headBytes,
	}, func(code uint16) {
		pkt.Fail(code)
	}, func(msg util.IMsg) {
		pkt.Ok(res)
	})
}

func (s *Svc) OnUserSignIn(pkt kiwi.IRcvRequest, req *pb.UserSignInReq, res *pb.UserSignInRes) {
	userId, _ := util.MGet[string](pkt.Head(), common.HdUserId)
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	now := time.Now().Unix()
	token, err := common.GenToken(addr, userId, now)
	if err != nil {
		pkt.Fail(util.EcServiceErr)
		return
	}
	user, err := GetUserWithId(userId)
	if err != nil {
		pkt.Fail(common.EcUserNotExist)
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
	_ = user.Store()

	head := util.M{
		common.HdMask: user.RoleMask,
	}
	headBytes, _ := head.ToBytes()
	pkt.AsyncReq(&pb.GateAddrUpdateReq{
		Addr: addr,
		Head: headBytes,
	}, func(code uint16) {
		pkt.Fail(code)
	}, func(msg util.IMsg) {
		res.User = user.User
		pkt.Ok(res)
	})
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
	tkn, _ := common.GenToken(addr, user.Id, user.LastSignInTime)
	user.SetLastSignInAddr(addr)
	user.SetToken(tkn)

	pkt.AsyncReq(&pb.GateAddrUpdateReq{
		Addr: addr,
		Head: user.Head,
	}, func(code uint16) {
		pkt.Fail(code)
	}, func(msg util.IMsg) {
		pkt.Ok(res)
	})
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
	_, err := mgo.UpdateOne(SchemaUser, bson.M{
		Id: req.Id,
	}, bson.M{
		"$set": bson.M{
			Head: req.Head,
		},
	})
	if err != nil {
		pkt.Fail(util.EcServiceErr)
		return
	}
	pkt.Ok(res)
}
