// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package gate

import (
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/util"
)

func (s *svc) OnGateHeartbeat(pkt kiwi.IRcvRequest, req *pb.GateHeartbeatReq, res *pb.GateHeartbeatRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateSendToId(pkt kiwi.IRcvRequest, req *pb.GateSendToIdReq, res *pb.GateSendToIdRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateSendToAddr(pkt kiwi.IRcvRequest, req *pb.GateSendToAddrReq, res *pb.GateSendToAddrRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateSendToMultiId(pkt kiwi.IRcvRequest, req *pb.GateSendToMultiIdReq, res *pb.GateSendToMultiIdRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateSendToMultiAddr(pkt kiwi.IRcvRequest, req *pb.GateSendToMultiAddrReq, res *pb.GateSendToMultiAddrRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateSendToAll(pkt kiwi.IRcvRequest, req *pb.GateSendToAllReq, res *pb.GateSendToAllRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateCloseId(pkt kiwi.IRcvRequest, req *pb.GateCloseIdReq, res *pb.GateCloseIdRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateCloseAddr(pkt kiwi.IRcvRequest, req *pb.GateCloseAddrReq, res *pb.GateCloseAddrRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateUpdate(pkt kiwi.IRcvRequest, req *pb.GateUpdateReq, res *pb.GateUpdateRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateRemove(pkt kiwi.IRcvRequest, req *pb.GateRemoveReq, res *pb.GateRemoveRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateGet(pkt kiwi.IRcvRequest, req *pb.GateGetReq, res *pb.GateGetRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}

func (s *svc) OnGateUpdateRoles(pkt kiwi.IRcvRequest, req *pb.GateUpdateRolesReq, res *pb.GateUpdateRolesRes) {
	pkt.Err2(util.EcNotImplement, util.M{"req": req})
}
