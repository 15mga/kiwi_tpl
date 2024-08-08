package common

// header key
const (
	HdGateAddr    = "addr"  //user token string
	HdGateId      = "g_id"  //gate node id
	HdMask        = "mask"  //role mask int64
	HdAdmin       = "admin" //admin bool 游戏管理超管
	HdSuper       = "super" //super bool 平台管理超管
	HdSvc         = "svc"   //服务号 uint16
	HdCode        = "cod"   //请求消息 code uint16
	HdUserId      = "id"    //user id string
	HdGroupId     = "gid"   //group id 亲友圈 string
	HdRoomId      = "rid"   //room id string
	HdRoomTypeI32 = "rt"    //房间类型 int32
	HdRoomSeatI32 = "seat"  //房间座位号 int32
	HdGame        = "game"  //加入的游戏类型 string
)
