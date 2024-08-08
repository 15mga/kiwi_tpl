// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package user

import (
	"game/proto/pb"
	"github.com/15mga/kiwi/util"
	"github.com/15mga/kiwi/util/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserTest() *UserTest {
	m := &UserTest{
		UserTest: &pb.UserTest{},
	}
	m.Model = util.NewModel(SchemaUserTest, 1, m.GetVal)
	return m
}

type UserTest struct {
	*pb.UserTest
	*util.Model
}

func (this *UserTest) LoadWithId(id string) error {
	return mgo.FindOne(SchemaUserTest, bson.M{"_id": id}, &this.UserTest)
}

func (this *UserTest) Load(filter any) error {
	return mgo.FindOne(SchemaUserTest, filter, &this.UserTest)
}

func (this *UserTest) UpdateDb() (*mongo.UpdateResult, error) {
	update := bson.M{}
	this.GenUpdate(update)
	return mgo.UpdateOne(SchemaUserTest, bson.M{"_id": this.Id}, bson.M{"$set": update})
}

func (this *UserTest) GetVal(key string) any {
	switch key {
	default:
		return nil
	}
}

func NewUser() *User {
	m := &User{
		User: &pb.User{},
	}
	m.Model = util.NewModel(SchemaUser, 23, m.GetVal)
	return m
}

type User struct {
	*pb.User
	*util.Model
}

func (this *User) PasswordSet(val string) {
	this.Password = val
	this.SetDirty(Password)
}

func (this *User) RoleMaskSet(val int64) {
	this.RoleMask = val
	this.SetDirty(RoleMask)
}

func (this *User) BanSet(val bool) {
	this.Ban = val
	this.SetDirty(Ban)
}

func (this *User) NickSet(val string) {
	this.Nick = val
	this.SetDirty(Nick)
}

func (this *User) AddrSet(val string) {
	this.Addr = val
	this.SetDirty(Addr)
}

func (this *User) IdCardSet(val string) {
	this.IdCard = val
	this.SetDirty(IdCard)
}

func (this *User) RealNameSet(val string) {
	this.RealName = val
	this.SetDirty(RealName)
}

func (this *User) MobileSet(val string) {
	this.Mobile = val
	this.SetDirty(Mobile)
}

func (this *User) SignUpTimeSet(val int64) {
	this.SignUpTime = val
	this.SetDirty(SignUpTime)
}

func (this *User) LastSignInTimeSet(val int64) {
	this.LastSignInTime = val
	this.SetDirty(LastSignInTime)
}

func (this *User) LastSignInAddrSet(val string) {
	this.LastSignInAddr = val
	this.SetDirty(LastSignInAddr)
}

func (this *User) LastOsSet(val string) {
	this.LastOs = val
	this.SetDirty(LastOs)
}

func (this *User) StatusSet(val pb.OnlineState) {
	this.Status = val
	this.SetDirty(Status)
}

func (this *User) AvatarSet(val string) {
	this.Avatar = val
	this.SetDirty(Avatar)
}

func (this *User) WechatUnionIdSet(val string) {
	this.WechatUnionId = val
	this.SetDirty(WechatUnionId)
}

func (this *User) WechatCodeSet(val string) {
	this.WechatCode = val
	this.SetDirty(WechatCode)
}

func (this *User) TokenSet(val string) {
	this.Token = val
	this.SetDirty(Token)
}

func (this *User) HeadSet(val []byte) {
	this.Head = val
	this.SetDirty(Head)
}

func (this *User) LastOfflineTimeSet(val int64) {
	this.LastOfflineTime = val
	this.SetDirty(LastOfflineTime)
}

func (this *User) OnlineDurSet(val int64) {
	this.OnlineDur = val
	this.SetDirty(OnlineDur)
}

func (this *User) TestSet(val bool) {
	this.Test = val
	this.SetDirty(Test)
}

func (this *User) SetCharacterIds(val []string) {
	this.CharacterIds = val
	this.SetDirty(CharacterIds)
}

func (this *User) PushCharacterIds(items ...string) {
	this.CharacterIds = append(this.CharacterIds, items...)
	this.SetDirty(CharacterIds)
}

func (this *User) AddToSetCharacterIds(items ...string) {
	for _, item := range items {
		for _, v := range this.CharacterIds {
			if v == item {
				return
			}
		}
		this.CharacterIds = append(this.CharacterIds, item)
	}
	this.SetDirty(CharacterIds)
}

func (this *User) PullCharacterIds(items ...string) {
	if this.CharacterIds == nil || len(this.CharacterIds) == 0 {
		return
	}
	dirty := false
	for _, item := range items {
		for i, v := range this.CharacterIds {
			if v == item {
				this.CharacterIds = append(this.CharacterIds[:i], this.CharacterIds[i+1:]...)
				dirty = true
				break
			}
		}
	}
	if dirty {
		this.SetDirty(CharacterIds)
	}
}

func (this *User) LoadWithId(id string) error {
	return mgo.FindOne(SchemaUser, bson.M{"_id": id}, &this.User)
}

func (this *User) Load(filter any) error {
	return mgo.FindOne(SchemaUser, filter, &this.User)
}

func (this *User) UpdateDb() (*mongo.UpdateResult, error) {
	update := bson.M{}
	this.GenUpdate(update)
	return mgo.UpdateOne(SchemaUser, bson.M{"_id": this.Id}, bson.M{"$set": update})
}

func (this *User) GetVal(key string) any {
	switch key {
	case Password:
		return this.Password
	case RoleMask:
		return this.RoleMask
	case Ban:
		return this.Ban
	case Nick:
		return this.Nick
	case Addr:
		return this.Addr
	case IdCard:
		return this.IdCard
	case RealName:
		return this.RealName
	case Mobile:
		return this.Mobile
	case SignUpTime:
		return this.SignUpTime
	case LastSignInTime:
		return this.LastSignInTime
	case LastSignInAddr:
		return this.LastSignInAddr
	case LastOs:
		return this.LastOs
	case Status:
		return this.Status
	case Avatar:
		return this.Avatar
	case WechatUnionId:
		return this.WechatUnionId
	case WechatCode:
		return this.WechatCode
	case Token:
		return this.Token
	case Head:
		return this.Head
	case LastOfflineTime:
		return this.LastOfflineTime
	case OnlineDur:
		return this.OnlineDur
	case Test:
		return this.Test
	case CharacterIds:
		return this.CharacterIds
	default:
		return nil
	}
}
