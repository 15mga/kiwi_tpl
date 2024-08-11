package user

import (
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/sid"
	"github.com/15mga/kiwi/util"
	"github.com/15mga/kiwi/util/mgo"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

func (s *Svc) BeforeShutdown() {
	StoreAllUsers()
}

func (s *Svc) OnUserSignUp(pkt kiwi.IRcvRequest, req *pb.UserSignUpReq, res *pb.UserSignUpRes) {
	addr, _ := util.MGet[string](pkt.Head(), common.HdGateAddr)
	_, err := mgo.InsertOne(SchemaUser, &pb.User{
		Id:         sid.GetStrId(),
		Password:   common.SaltPw(req.Password),
		RoleMask:   util.GenMask(common.RPlayer),
		Ban:        false,
		Nick:       req.Nick,
		Addr:       addr,
		Mobile:     req.Mobile,
		SignUpTime: time.Now().Unix(),
		Avatar:     req.Avatar,
	})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			es := err.Error()
			switch {
			case strings.Contains(es, "index: mobile"):
				pkt.Fail(common.EcMobileExist)
			case strings.Contains(es, "index: nick"):
				pkt.Fail(common.EcNickExist)
			default:
				pkt.Fail(util.EcServiceErr)
			}
		}
		return
	}
	pkt.Ok(res)
}
