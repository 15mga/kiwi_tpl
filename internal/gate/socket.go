package gate

import (
	"game/internal/codec"
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/util"
	"strconv"
	"sync"
)

// CheckIp 检查是否允许连接
func CheckIp(ip string) bool {
	return true
}

// RecordSocketErr 记录错误
func RecordSocketErr(agent kiwi.IAgent, err *util.Err) {
	head := util.M{}
	agent.CopyHead(head)
	err.AddParam("head", head)
	kiwi.Error(err)
}

func Disconnected(agent kiwi.IAgent, err *util.Err) {
	if err != nil {
		err.AddParam("addr", agent.Addr())
		kiwi.Error(err)
	}
	head := util.M{}
	agent.CopyHead(head)
	id, _ := util.MGet[string](head, common.HdUserId)
	_svc.AsyncReq(0, head, &pb.UserDisconnectReq{
		UserId: id,
	}, nil, nil)
}

func InitHead(m util.M) {
	m[common.HdMask] = util.GenMask(common.RGuest)
	m[strconv.Itoa(int(common.Gate))] = _svc.Meta().Id
}

func SocketReceiver(agent kiwi.IAgent, bytes []byte) {
	svc, code, payload, err := common.UnpackUserReq(bytes)
	if err != nil {
		RecordSocketErr(agent, err)
		return
	}
	roleMask, _ := agent.GetHead(common.HdMask)
	mask := roleMask.(int64)
	ok := kiwi.Gate().Authenticate(mask, svc, code)
	if !ok {
		RecordSocketErr(agent, util.NewErr(util.EcNoAuth, util.M{
			"service": svc,
			"code":    code,
		}))
		return
	}

	head := util.M{}
	agent.CopyHead(head)
	req := _requestPool.Get().(*request)
	//req := &request{}
	req.Request(svc, code, head, payload)
}

var (
	_requestPool = sync.Pool{
		New: func() interface{} {
			return &request{}
		},
	}
)

type request struct {
	svc  kiwi.TSvc
	code kiwi.TMethod
	json bool
	addr string
}

func (r *request) Request(svc kiwi.TSvc, code kiwi.TMethod, head util.M, payload []byte) {
	r.svc = svc
	r.code = code
	r.addr, _ = util.MGet[string](head, common.HdGateAddr)
	_svc.AsyncReqBytes(0, svc, code, head, true, payload, r.OnErr, r.OnOk)
}

func (r *request) RequestNode(nodeId int64, svc kiwi.TSvc, code kiwi.TMethod, head util.M, payload []byte) {
	r.svc = svc
	r.code = code
	r.addr, _ = util.MGet[string](head, common.HdGateAddr)
	_svc.AsyncReqNodeBytes(0, nodeId, svc, code, head, true, payload, r.OnErr, r.OnOk)
}

func (r *request) OnOk(payload []byte) {
	defer _requestPool.Put(r)
	resCode, _ := kiwi.Codec().ReqToResMethod(r.svc, r.code)
	pkt, e := common.PackUserOk(kiwi.MergeSvcMethod(r.svc, resCode), payload)
	if e != nil {
		kiwi.Error(e)
		return
	}

	kiwi.Gate().AddrSend(0, r.addr, pkt, nil)
}

func (r *request) OnErr(errCode uint16) {
	defer _requestPool.Put(r)
	resCode, _ := kiwi.Codec().ReqToResMethod(r.svc, r.code)
	pkt, e := common.PackUserFail(kiwi.MergeSvcMethod(common.Gate, codec.GateErrPus), kiwi.MergeSvcMethod(r.svc, resCode), errCode)
	if e != nil {
		kiwi.Error(e)
		return
	}
	kiwi.Gate().AddrSend(0, r.addr, pkt, nil)
}
