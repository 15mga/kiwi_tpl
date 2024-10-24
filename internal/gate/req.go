package gate

import (
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/util"
	"time"
)

func (s *Svc) OnGateHeartbeat(pkt kiwi.IRcvRequest, req *pb.GateHeartbeatReq, res *pb.GateHeartbeatRes) {
	res.ReqTs = req.ReqTs
	res.ResTs = time.Now().UnixMilli()
	pkt.Ok(res)
}

func (s *Svc) OnGateSendToId(pkt kiwi.IRcvRequest, req *pb.GateSendToIdReq, res *pb.GateSendToIdRes) {
	kiwi.Gate().Send(pkt.Tid(), req.Id, req.Payload, func(ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistId)
			return
		}
		pkt.Ok(res)
	})
}

func (s *Svc) OnGateSendToAddr(pkt kiwi.IRcvRequest, req *pb.GateSendToAddrReq, res *pb.GateSendToAddrRes) {
	kiwi.Gate().AddrSend(pkt.Tid(), req.Addr, req.Payload, func(ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistAddr)
			return
		}
		pkt.Ok(res)
	})
}

func (s *Svc) OnGateSendToMultiId(pkt kiwi.IRcvRequest, req *pb.GateSendToMultiIdReq, res *pb.GateSendToMultiIdRes) {
	payload := req.Payload
	var buffer util.ByteBuffer
	buffer.InitBytes(payload)
	count, err := buffer.RUint16()
	if err != nil {
		kiwi.TE(pkt.Tid(), err)
		return
	}
	idToMsg := make(map[string][]byte, count)
	for i := uint16(0); i < count; i++ {
		id, err := buffer.RString()
		if err != nil {
			kiwi.TE(pkt.Tid(), err)
			return
		}
		p, err := buffer.RBytes()
		if err != nil {
			kiwi.TE(pkt.Tid(), err)
			return
		}
		idToMsg[id] = p
	}
	kiwi.Gate().MultiSend(pkt.Tid(), idToMsg, func(m map[string]bool) {
		res.Result = m
	})
	pkt.Ok(res)
}

func (s *Svc) OnGateSendToMultiAddr(pkt kiwi.IRcvRequest, req *pb.GateSendToMultiAddrReq, res *pb.GateSendToMultiAddrRes) {
	payload := req.Payload
	var buffer util.ByteBuffer
	buffer.InitBytes(payload)
	count, err := buffer.RUint16()
	if err != nil {
		kiwi.TE(pkt.Tid(), err)
		return
	}
	addrToMsg := make(map[string][]byte, count)
	for i := uint16(0); i < count; i++ {
		id, err := buffer.RString()
		if err != nil {
			kiwi.TE(pkt.Tid(), err)
			return
		}
		p, err := buffer.RBytes()
		if err != nil {
			kiwi.TE(pkt.Tid(), err)
			return
		}
		addrToMsg[id] = p
	}
	kiwi.Gate().MultiAddrSend(pkt.Tid(), addrToMsg, func(m map[string]bool) {
		res.Result = m
	})
	pkt.Ok(res)
}

func (s *Svc) OnGateSendToAll(pkt kiwi.IRcvRequest, req *pb.GateSendToAllReq, res *pb.GateSendToAllRes) {
	kiwi.Gate().AllSend(pkt.Tid(), req.Payload)
	pkt.Ok(res)
}

func (s *Svc) OnGateCloseId(pkt kiwi.IRcvRequest, req *pb.GateCloseIdReq, res *pb.GateCloseIdRes) {
	kiwi.Gate().CloseWithId(pkt.Tid(), req.Id)
	pkt.Ok(res)
}

func (s *Svc) OnGateCloseAddr(pkt kiwi.IRcvRequest, req *pb.GateCloseAddrReq, res *pb.GateCloseAddrRes) {
	kiwi.Gate().CloseWithAddr(pkt.Tid(), req.Addr)
	pkt.Ok(res)
}

func (s *Svc) OnGateUpdate(pkt kiwi.IRcvRequest, req *pb.GateUpdateReq, res *pb.GateUpdateRes) {
	var head util.M
	if req.Head != nil {
		head = make(util.M)
		err := kiwi.Packer().UnpackM(req.Head, head)
		if err != nil {
			kiwi.TE(pkt.Tid(), err)
			return
		}
	}
	kiwi.Gate().UpdateHead(pkt.Tid(), req.Id, head, func(head util.M, ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistId)
			return
		}
		pkt.Ok(res)
		s.updateHead(pkt, head)
		kiwi.TD(pkt.Tid(), "update head", util.M{
			"id":   req.Id,
			"head": head,
		})
	})
}

func (s *Svc) updateHead(pkt kiwi.IRcvRequest, head util.M) {
	_, ok := util.MGet[string](pkt.Head(), common.HdUserId)
	if !ok {
		return
	}
	delete(head, common.HdGateAddr)
	delete(head, common.HdSvc)
	delete(head, common.HdMtd)
	headBytes, _ := head.ToBytes()
	s.AsyncReq(pkt.Tid(), pkt.Head().Copy(), &pb.UserUpdateHeadReq{
		Head: headBytes,
	}, nil, nil)
}

func (s *Svc) OnGateRemove(pkt kiwi.IRcvRequest, req *pb.GateRemoveReq, res *pb.GateRemoveRes) {
	kiwi.Gate().RemoveHead(pkt.Tid(), req.Id, req.Head, func(head util.M, ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistId)
			return
		}
		pkt.Ok(res)
		s.updateHead(pkt, head)
	})
}

func (s *Svc) OnGateGet(pkt kiwi.IRcvRequest, req *pb.GateGetReq, res *pb.GateGetRes) {
	kiwi.Gate().GetHead(pkt.Tid(), req.Id, func(head util.M, ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistId)
			return
		}
		res.Head, _ = kiwi.Packer().PackM(head)
		pkt.Ok(res)
	})
	if req.Close {
		kiwi.Gate().CloseWithId(pkt.Tid(), req.Id)
	}
}

func (s *Svc) OnGateUpdateRoles(pkt kiwi.IRcvRequest, req *pb.GateUpdateRolesReq, res *pb.GateUpdateRolesRes) {
	kiwi.Gate().UpdateRoles(req.Roles)
	pkt.Ok(res)
}
