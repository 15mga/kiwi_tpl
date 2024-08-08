package common

import (
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/util"
)

func PackUserReq(svc kiwi.TSvc, code kiwi.TCode, msg util.IMsg) ([]byte, *util.Err) {
	bytes, err := util.JsonMarshal(msg)
	if err != nil {
		return nil, err
	}
	var buffer util.ByteBuffer
	buffer.InitCap(2 + len(bytes))
	sc := kiwi.MergeSvcCode(svc, code)
	buffer.WUint32(sc)
	_, e := buffer.Write(bytes)
	if e != nil {
		return nil, util.NewErr(util.EcWriteFail, util.M{
			"error": e,
		})
	}
	return buffer.All(), nil
}

func UnpackUserReq(bytes []byte) (svc kiwi.TSvc, code kiwi.TCode, payload []byte, err *util.Err) {
	var buffer util.ByteBuffer
	buffer.InitBytes(bytes)
	sc, err := buffer.RUint32()
	if err != nil {
		return
	}
	svc, code = kiwi.SplitSvcCode(sc)
	payload = buffer.RAvailable()
	return
}

func PackUserOk(resCode kiwi.TSvcCode, payload []byte) ([]byte, *util.Err) {
	var buffer util.ByteBuffer
	buffer.InitCap(len(payload) + 2)
	buffer.WUint32(resCode)
	_, e := buffer.Write(payload)
	if e != nil {
		return nil, util.NewErr(util.EcWriteFail, util.M{
			"error": e,
		})
	}
	return buffer.All(), nil
}

func UnpackUserOk(bytes []byte) (svc kiwi.TSvc, code kiwi.TCode, msg util.IMsg, err *util.Err) {
	var buffer util.ByteBuffer
	buffer.InitBytes(bytes)
	var resCode uint32
	resCode, err = buffer.RUint32()
	if err != nil {
		return
	}
	svc, code = kiwi.SplitSvcCode(resCode)
	payload := buffer.RAvailable()
	msg, err = kiwi.Codec().JsonUnmarshal2(svc, code, payload)
	return
}

func PackUserFail(failMsgCode kiwi.TSvcCode, resSvcCode kiwi.TSvcCode, errCode uint16) ([]byte, *util.Err) {
	bytes, err := util.JsonMarshal(&pb.GateErrPus{
		MsgCode: int32(resSvcCode),
		ErrCode: int32(errCode),
	})
	if err != nil {
		return nil, err
	}
	return PackUserOk(failMsgCode, bytes)
}

func PackUserPus(svc kiwi.TSvc, code kiwi.TCode, ntc []byte) ([]byte, *util.Err) {
	return PackUserOk(kiwi.MergeSvcCode(svc, code), ntc)
}
