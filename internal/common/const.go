package common

// header key
const (
	HdGateAddr  = "addr" //用户连接网关的远程地址 string
	HdMask      = "mask" //role mask int64
	HdSvc       = "svc"  //当前请求服务 uint16
	HdMtd       = "cod"  //当前请求方法 uint16
	HdSignInCh  = "ch"   //登录方式
	HdAccountId = "aid"  //账号id
	HdUserId    = "uid"  //user id string
	HdId        = "id"   //id string 可能是HdGateAddr,HdAccountId,HdUserId
)

const (
	SignInMobile = "mobile"
	SignInEmail  = "email"
	SignInWechat = "wechat"
)
