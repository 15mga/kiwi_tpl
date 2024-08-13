package gate

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"game/internal/codec"
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/util"
	"os"
	"time"
)

func (s *Svc) OnGateHeartbeat(pkt kiwi.IRcvRequest, req *pb.GateHeartbeatReq, res *pb.GateHeartbeatRes) {
	res.ReqTs = req.ReqTs
	res.ResTs = time.Now().UnixMilli()
	pkt.Ok(res)
}

func (s *Svc) OnGateUploadFile(pkt kiwi.IRcvRequest, req *pb.GateUploadFileReq, res *pb.GateUploadFileRes) {
	wd, _ := os.Getwd()
	relativeDir := fmt.Sprintf("static/%s", req.Type)
	dir := fmt.Sprintf("%s/%s", wd, relativeDir)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		pkt.Fail(util.EcServiceErr)
		return
	}
	relativePath := fmt.Sprintf("%s/%s_%d%s", relativeDir, pkt.HeadId(), time.Now().UnixMilli(), req.FileExt)
	path := fmt.Sprintf("%s/%s", wd, relativePath)
	bytes, err := hex.DecodeString(req.Data)
	if err != nil {
		pkt.Fail(util.EcUnmarshallErr)
		return
	}
	err = os.WriteFile(path, bytes, os.ModePerm)
	if err != nil {
		pkt.Fail(util.EcServiceErr)
		return
	}
	res.Url = relativePath
	pkt.Ok(res)
}

func (s *Svc) OnGateUploadWithToken(pkt kiwi.IRcvRequest, req *pb.GateUploadWithTokenReq, res *pb.GateUploadWithTokenRes) {
	claims, e := common.ParseToken(req.Token)
	if e != nil {
		pkt.Fail(util.EcNoAuth)
		return
	}
	wd, _ := os.Getwd()
	relativeDir := fmt.Sprintf("static/%s", req.Type)
	dir := fmt.Sprintf("%s/%s", wd, relativeDir)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		pkt.Fail(util.EcServiceErr)
		return
	}
	bytes, err := hex.DecodeString(req.Data)
	if err != nil {
		pkt.Fail(util.EcUnmarshallErr)
		return
	}

	md5Bytes := md5.Sum(bytes)
	name := hex.EncodeToString(md5Bytes[:])
	relativePath := fmt.Sprintf("%s/%s%s", relativeDir, name, req.FileExt)
	path := fmt.Sprintf("%s/%s", wd, relativePath)

	err = os.WriteFile(path, bytes, os.ModePerm)
	if err != nil {
		pkt.Fail(util.EcServiceErr)
		return
	}
	res.Url = relativePath
	payload, _ := util.JsonMarshal(&pb.GateUploadWithTokenPus{
		Url: relativePath,
	})
	bytes, e = common.PackUserPus(common.Gate, codec.GateUploadWithTokenPus, payload)
	if e != nil {
		kiwi.TE(pkt.Tid(), e)
		return
	}
	kiwi.Gate().Send(pkt.Tid(), claims.Uid, bytes, func(ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistAddr)
			return
		}
		pkt.Ok(res)
	})
}

func (s *Svc) OnGateSendToId(pkt kiwi.IRcvRequest, req *pb.GateSendToIdReq, res *pb.GateSendToIdRes) {
	svc, code := kiwi.SplitSvcCode(req.SvcCode)
	bytes, err := common.PackUserPus(svc, code, req.Payload)
	if err != nil {
		kiwi.TE(pkt.Tid(), err)
		pkt.Err(err)
		return
	}
	kiwi.Gate().Send(pkt.Tid(), req.Id, bytes, func(ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistId)
			return
		}
		pkt.Ok(res)
	})
}

func (s *Svc) OnGateSendToAddr(pkt kiwi.IRcvRequest, req *pb.GateSendToAddrReq, res *pb.GateSendToAddrRes) {
	svc, code := kiwi.SplitSvcCode(req.SvcCode)
	bytes, err := common.PackUserPus(svc, code, req.Payload)
	if err != nil {
		kiwi.TE(pkt.Tid(), err)
		return
	}
	kiwi.Gate().AddrSend(pkt.Tid(), req.Addr, bytes, func(ok bool) {
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
		pkt.Ok(res)
	})
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
		pkt.Ok(res)
	})
}

func (s *Svc) OnGateSendToAll(pkt kiwi.IRcvRequest, req *pb.GateSendToAllReq, res *pb.GateSendToAllRes) {
	svc, code := kiwi.SplitSvcCode(req.SvcCode)
	bytes, err := common.PackUserPus(svc, code, req.Payload)
	if err != nil {
		kiwi.TE(pkt.Tid(), err)
		return
	}
	kiwi.Gate().AllSend(pkt.Tid(), bytes)
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
	})
}

func (s *Svc) OnGateAddrUpdate(pkt kiwi.IRcvRequest, req *pb.GateAddrUpdateReq, res *pb.GateAddrUpdateRes) {
	var head util.M
	if req.Head != nil {
		head = make(util.M)
		err := kiwi.Packer().UnpackM(req.Head, head)
		if err != nil {
			kiwi.TE(pkt.Tid(), err)
			pkt.Fail(err.Code())
			return
		}
	}
	kiwi.Gate().UpdateAddrHead(pkt.Tid(), req.Addr, head, func(head util.M, ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistAddr)
			return
		}
		pkt.Ok(res)
		s.updateHead(pkt, head)
	})
}

func (s *Svc) updateHead(pkt kiwi.IRcvRequest, head util.M) {
	uid, _ := util.MGet[string](head, common.HdUserId)
	delete(head, common.HdGateAddr)
	delete(head, common.HdGateId)
	delete(head, common.HdSvc)
	delete(head, common.HdCode)
	headBytes, _ := head.ToBytes()
	s.Req(pkt.Tid(), util.M{
		common.HdUserId: uid,
	}, &pb.UserUpdateHeadReq{
		Id:   uid,
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

func (s *Svc) OnGateAddrRemove(pkt kiwi.IRcvRequest, req *pb.GateAddrRemoveReq, res *pb.GateAddrRemoveRes) {
	kiwi.Gate().RemoveAddrHead(pkt.Tid(), req.Addr, req.Head, func(head util.M, ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistAddr)
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

func (s *Svc) OnGateAddrGet(pkt kiwi.IRcvRequest, req *pb.GateAddrGetReq, res *pb.GateAddrGetRes) {
	kiwi.Gate().GetAddrHead(pkt.Tid(), req.Addr, func(head util.M, ok bool) {
		if !ok {
			pkt.Fail(common.EcGateNotExistAddr)
			return
		}
		res.Head, _ = kiwi.Packer().PackM(head)
		pkt.Ok(res)
	})
	if req.Close {
		kiwi.Gate().CloseWithAddr(pkt.Tid(), req.Addr)
	}
}
