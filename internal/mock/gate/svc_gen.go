// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package gate

import (
	"game/internal/codec"
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/graph"
	"github.com/15mga/kiwi/mock"
	"github.com/15mga/kiwi/util"
	"strconv"
)

type Svc struct {
	svc
}

type svc struct {
	client *mock.Client
}

func InitClient(client *mock.Client) {
	s := &Svc{svc{client: client}}
	s.client.BindPointMsg("gate", "GateHeartbeat", s.inGateHeartbeatReq)
	s.client.BindNetMsg(&pb.GateHeartbeatRes{}, s.onGateHeartbeatRes)
	s.client.BindNetMsg(&pb.GateErrPus{}, s.onGateErrPus)
	s.client.BindPointMsg("gate", "GateSendToId", s.inGateSendToIdReq)
	s.client.BindNetMsg(&pb.GateSendToIdRes{}, s.onGateSendToIdRes)
	s.client.BindPointMsg("gate", "GateSendToAddr", s.inGateSendToAddrReq)
	s.client.BindNetMsg(&pb.GateSendToAddrRes{}, s.onGateSendToAddrRes)
	s.client.BindPointMsg("gate", "GateSendToMultiId", s.inGateSendToMultiIdReq)
	s.client.BindNetMsg(&pb.GateSendToMultiIdRes{}, s.onGateSendToMultiIdRes)
	s.client.BindPointMsg("gate", "GateSendToMultiAddr", s.inGateSendToMultiAddrReq)
	s.client.BindNetMsg(&pb.GateSendToMultiAddrRes{}, s.onGateSendToMultiAddrRes)
	s.client.BindPointMsg("gate", "GateSendToAll", s.inGateSendToAllReq)
	s.client.BindNetMsg(&pb.GateSendToAllRes{}, s.onGateSendToAllRes)
	s.client.BindPointMsg("gate", "GateCloseId", s.inGateCloseIdReq)
	s.client.BindNetMsg(&pb.GateCloseIdRes{}, s.onGateCloseIdRes)
	s.client.BindPointMsg("gate", "GateCloseAddr", s.inGateCloseAddrReq)
	s.client.BindNetMsg(&pb.GateCloseAddrRes{}, s.onGateCloseAddrRes)
	s.client.BindPointMsg("gate", "GateUpdate", s.inGateUpdateReq)
	s.client.BindNetMsg(&pb.GateUpdateRes{}, s.onGateUpdateRes)
	s.client.BindPointMsg("gate", "GateRemove", s.inGateRemoveReq)
	s.client.BindNetMsg(&pb.GateRemoveRes{}, s.onGateRemoveRes)
	s.client.BindPointMsg("gate", "GateGet", s.inGateGetReq)
	s.client.BindNetMsg(&pb.GateGetRes{}, s.onGateGetRes)
	s.client.BindPointMsg("gate", "GateUpdateRoles", s.inGateUpdateRolesReq)
	s.client.BindNetMsg(&pb.GateUpdateRolesRes{}, s.onGateUpdateRolesRes)
}

func (s *svc) Dispose() {
}

func (s *svc) AsyncReq(req util.IMsg) *util.Err {
	kiwi.Debug("request", util.M{string(req.ProtoReflect().Descriptor().Name()): req})
	svc, code := kiwi.Codec().MsgToSvcMethod(req)
	bytes, err := common.PackUserReq(svc, code, req)
	if err != nil {
		return err
	}
	return s.client.Dialer().Agent().Send(bytes)
}

func (s *svc) inGateHeartbeatReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateHeartbeatReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateHeartbeatRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateHeartbeat", nil
}

func (s *svc) onGateErrPus(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "", nil
}

func (s *svc) inGateSendToIdReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateSendToIdReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateSendToIdRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateSendToId", nil
}

func (s *svc) inGateSendToAddrReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateSendToAddrReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateSendToAddrRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateSendToAddr", nil
}

func (s *svc) inGateSendToMultiIdReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateSendToMultiIdReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateSendToMultiIdRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateSendToMultiId", nil
}

func (s *svc) inGateSendToMultiAddrReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateSendToMultiAddrReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateSendToMultiAddrRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateSendToMultiAddr", nil
}

func (s *svc) inGateSendToAllReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateSendToAllReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateSendToAllRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateSendToAll", nil
}

func (s *svc) inGateCloseIdReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateCloseIdReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateCloseIdRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateCloseId", nil
}

func (s *svc) inGateCloseAddrReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateCloseAddrReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateCloseAddrRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateCloseAddr", nil
}

func (s *svc) inGateUpdateReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateUpdateReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateUpdateRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateUpdate", nil
}

func (s *svc) inGateRemoveReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateRemoveReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateRemoveRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateRemove", nil
}

func (s *svc) inGateGetReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateGetReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateGetRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateGet", nil
}

func (s *svc) inGateUpdateRolesReq(msg graph.IMsg) *util.Err {
	req := s.client.GetRequest(common.Gate, codec.GateUpdateRolesReq)
	return s.AsyncReq(req)
}

func (s *svc) onGateUpdateRolesRes(msg util.IMsg) (point string, data any) {
	sc := kiwi.MergeSvcCode(kiwi.Codec().MsgToSvcMethod(msg))
	s.client.Graph().Data().Set(strconv.Itoa(int(sc)), msg)
	return "GateUpdateRoles", nil
}
