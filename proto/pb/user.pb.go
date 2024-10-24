// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.1
// source: model/user.proto

package pb

import (
	_ "github.com/15mga/kiwi_tool"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OnlineState int32

const (
	OnlineState_Disconnected OnlineState = 0 //离线
	OnlineState_Connected    OnlineState = 1 //在线
)

// Enum value maps for OnlineState.
var (
	OnlineState_name = map[int32]string{
		0: "Disconnected",
		1: "Connected",
	}
	OnlineState_value = map[string]int32{
		"Disconnected": 0,
		"Connected":    1,
	}
)

func (x OnlineState) Enum() *OnlineState {
	p := new(OnlineState)
	*p = x
	return p
}

func (x OnlineState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OnlineState) Descriptor() protoreflect.EnumDescriptor {
	return file_model_user_proto_enumTypes[0].Descriptor()
}

func (OnlineState) Type() protoreflect.EnumType {
	return &file_model_user_proto_enumTypes[0]
}

func (x OnlineState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OnlineState.Descriptor instead.
func (OnlineState) EnumDescriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	RoleMask        int64       `protobuf:"varint,2,opt,name=roleMask,proto3" json:"roleMask,omitempty" bson:"role_mask"`                        //角色组遮罩
	Ban             bool        `protobuf:"varint,3,opt,name=ban,proto3" json:"ban,omitempty" bson:"ban"`                                        //禁用
	Nick            string      `protobuf:"bytes,4,opt,name=nick,proto3" json:"nick,omitempty" bson:"nick"`                                      //昵称
	IdCard          string      `protobuf:"bytes,5,opt,name=idCard,proto3" json:"idCard,omitempty" bson:"id_card"`                               //身份证
	RealName        string      `protobuf:"bytes,6,opt,name=realName,proto3" json:"realName,omitempty" bson:"real_name"`                         //实名
	CreateTime      int64       `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime,omitempty" bson:"create_time"`                  //创建时间
	LastSignInTime  int64       `protobuf:"varint,8,opt,name=lastSignInTime,proto3" json:"lastSignInTime,omitempty" bson:"last_sign_in_time"`    //最后登录时间
	LastSignInAddr  string      `protobuf:"bytes,9,opt,name=lastSignInAddr,proto3" json:"lastSignInAddr,omitempty" bson:"last_sign_in_addr"`     //最后登录地址
	LastOfflineTime int64       `protobuf:"varint,10,opt,name=lastOfflineTime,proto3" json:"lastOfflineTime,omitempty" bson:"last_offline_time"` //最后下线时间
	LastOs          string      `protobuf:"bytes,11,opt,name=lastOs,proto3" json:"lastOs,omitempty" bson:"last_os"`                              //最后登录使用的系统
	State           OnlineState `protobuf:"varint,12,opt,name=state,proto3,enum=pb.OnlineState" json:"state,omitempty" bson:"state"`             //状态
	Avatar          string      `protobuf:"bytes,13,opt,name=avatar,proto3" json:"avatar,omitempty" bson:"avatar"`                               //头像地址
	Token           string      `protobuf:"bytes,14,opt,name=token,proto3" json:"token,omitempty" bson:"token"`                                  //token
	Head            []byte      `protobuf:"bytes,15,opt,name=head,proto3" json:"head,omitempty" bson:"head"`                                     //会话
	OnlineDur       int64       `protobuf:"varint,16,opt,name=onlineDur,proto3" json:"onlineDur,omitempty" bson:"online_dur"`                    //在线时长
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetRoleMask() int64 {
	if x != nil {
		return x.RoleMask
	}
	return 0
}

func (x *User) GetBan() bool {
	if x != nil {
		return x.Ban
	}
	return false
}

func (x *User) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *User) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *User) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *User) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *User) GetLastSignInTime() int64 {
	if x != nil {
		return x.LastSignInTime
	}
	return 0
}

func (x *User) GetLastSignInAddr() string {
	if x != nil {
		return x.LastSignInAddr
	}
	return ""
}

func (x *User) GetLastOfflineTime() int64 {
	if x != nil {
		return x.LastOfflineTime
	}
	return 0
}

func (x *User) GetLastOs() string {
	if x != nil {
		return x.LastOs
	}
	return ""
}

func (x *User) GetState() OnlineState {
	if x != nil {
		return x.State
	}
	return OnlineState_Disconnected
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *User) GetHead() []byte {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *User) GetOnlineDur() int64 {
	if x != nil {
		return x.OnlineDur
	}
	return 0
}

type MobileAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"` // mobile
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" bson:"password"`
	UserId   string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty" bson:"user_id"`
}

func (x *MobileAccount) Reset() {
	*x = MobileAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MobileAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MobileAccount) ProtoMessage() {}

func (x *MobileAccount) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MobileAccount.ProtoReflect.Descriptor instead.
func (*MobileAccount) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{1}
}

func (x *MobileAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MobileAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MobileAccount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type EmailAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"` // email
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" bson:"password"`
	UserId   string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty" bson:"user_id"`
}

func (x *EmailAccount) Reset() {
	*x = EmailAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailAccount) ProtoMessage() {}

func (x *EmailAccount) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailAccount.ProtoReflect.Descriptor instead.
func (*EmailAccount) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{2}
}

func (x *EmailAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EmailAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *EmailAccount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type WechatAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"` // wechatUnionId 微信联合 id
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty" bson:"user_id"`
}

func (x *WechatAccount) Reset() {
	*x = WechatAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WechatAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WechatAccount) ProtoMessage() {}

func (x *WechatAccount) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WechatAccount.ProtoReflect.Descriptor instead.
func (*WechatAccount) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{3}
}

func (x *WechatAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WechatAccount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_model_user_proto protoreflect.FileDescriptor

var file_model_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x31, 0x35, 0x6d, 0x67, 0x61, 0x2f, 0x6b, 0x69, 0x77, 0x69, 0x5f, 0x74, 0x6f,
	0x6f, 0x6c, 0x2f, 0x6b, 0x69, 0x77, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x04,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0e, 0x80, 0xea, 0x49, 0x01, 0x8a, 0xea, 0x49, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4d, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0x8a, 0xea, 0x49, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x10,
	0x0a, 0x03, 0x62, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x62, 0x61, 0x6e,
	0x12, 0x22, 0x0a, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e,
	0x80, 0xea, 0x49, 0x01, 0x8a, 0xea, 0x49, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x04,
	0x6e, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x8a, 0xea, 0x49, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x8a,
	0xea, 0x49, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x1a, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0x80, 0xea, 0x49, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x75, 0x72, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x75, 0x72, 0x3a,
	0x60, 0xf0, 0xe3, 0x49, 0x01, 0xfa, 0xe3, 0x49, 0x07, 0x0a, 0x05, 0x0a, 0x03, 0x62, 0x61, 0x6e,
	0xfa, 0xe3, 0x49, 0x10, 0x0a, 0x0e, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x75, 0x70, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0xfa, 0xe3, 0x49, 0x15, 0x0a, 0x13, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0xfa, 0xe3, 0x49,
	0x13, 0x0a, 0x11, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69,
	0x6e, 0x5f, 0x69, 0x70, 0xfa, 0xe3, 0x49, 0x09, 0x0a, 0x07, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x68, 0x0a, 0x0d, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x3a, 0x13, 0xf0, 0xe3, 0x49, 0x01, 0xfa, 0xe3, 0x49, 0x0b,
	0x0a, 0x09, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x0c, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x3a,
	0x13, 0xf0, 0xe3, 0x49, 0x01, 0xfa, 0xe3, 0x49, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x0d, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x3a, 0x13, 0xf0,
	0xe3, 0x49, 0x01, 0xfa, 0xe3, 0x49, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x2a, 0x2e, 0x0a, 0x0b, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x10, 0x01, 0x42, 0x14, 0xca, 0xdd, 0x49, 0x06, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x5a, 0x03,
	0x2f, 0x70, 0x62, 0xaa, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_user_proto_rawDescOnce sync.Once
	file_model_user_proto_rawDescData = file_model_user_proto_rawDesc
)

func file_model_user_proto_rawDescGZIP() []byte {
	file_model_user_proto_rawDescOnce.Do(func() {
		file_model_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_user_proto_rawDescData)
	})
	return file_model_user_proto_rawDescData
}

var file_model_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_user_proto_goTypes = []interface{}{
	(OnlineState)(0),      // 0: pb.OnlineState
	(*User)(nil),          // 1: pb.User
	(*MobileAccount)(nil), // 2: pb.MobileAccount
	(*EmailAccount)(nil),  // 3: pb.EmailAccount
	(*WechatAccount)(nil), // 4: pb.WechatAccount
}
var file_model_user_proto_depIdxs = []int32{
	0, // 0: pb.User.state:type_name -> pb.OnlineState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_user_proto_init() }
func file_model_user_proto_init() {
	if File_model_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MobileAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WechatAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_user_proto_goTypes,
		DependencyIndexes: file_model_user_proto_depIdxs,
		EnumInfos:         file_model_user_proto_enumTypes,
		MessageInfos:      file_model_user_proto_msgTypes,
	}.Build()
	File_model_user_proto = out.File
	file_model_user_proto_rawDesc = nil
	file_model_user_proto_goTypes = nil
	file_model_user_proto_depIdxs = nil
}
