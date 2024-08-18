package gate

import (
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/util"
	"github.com/15mga/kiwi/util/rds"
	"github.com/gomodule/redigo/redis"
	"io"
	"net/http"
	"strconv"
)

type HttpHeadCache struct {
}

func (c *HttpHeadCache) GetId(r *http.Request) (string, bool) {
	token := r.Header.Get("token")
	claims, e := common.ParseToken(token)
	if e != nil {
		return "", false
	}
	return claims.UserId, true
}

func (c *HttpHeadCache) GetHead(id string, head util.M) (newAgent bool) {
	conn := rds.SpawnConn()
	key := common.GetRdsKey("http", "head", id)
	bytes, err := redis.Bytes(conn.Do(rds.GET, key))
	if err != nil {
		_ = conn.Close()
		return true
	}
	_ = head.FromBytes(bytes)
	_ = conn.Close()
	return false
}

func (c *HttpHeadCache) SetHead(id string, head util.M) {
	conn := rds.SpawnConn()
	key := common.GetRdsKey("http", "head", id)
	bytes, _ := head.ToBytes()
	_ = conn.Send(rds.SET, key, bytes)
	_ = conn.Send(rds.EXPIRE, key, 60*30)
	_ = conn.Flush()
	_ = conn.Close()
}

func (c *HttpHeadCache) DelHead(id string) {
	conn := rds.SpawnConn()
	_, _ = conn.Do(rds.DEL, id)
	_ = conn.Close()
}

func (c *HttpHeadCache) InitHead(head util.M) {
	head[common.HdMask] = util.GenMask(common.RGuest)
}

func HttpDisconnected(head util.M) {
	id, ok := util.MGet[string](head, common.HdUserId)
	if !ok {
		return
	}
	_svc.AsyncReq(0, head, &pb.UserDisconnectReq{
		UserId: id,
	}, nil, nil)
}

// RecordHttpErr 记录错误
func RecordHttpErr(head util.M, r *http.Request, err *util.Err) {
	err.AddParam("head", head)
	kiwi.Error(err)
}

func httpResErr(w http.ResponseWriter, code uint16) {
	resBytes, _ := util.JsonMarshal(Response{
		Code: code,
	})
	_, _ = w.Write(resBytes)
}

func HttpReceiver(head util.M, w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// 响应预检请求，允许所有来源的跨域请求
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Token")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Token")

	svcStr := r.PathValue("svc")
	svcInt, err := strconv.Atoi(svcStr)
	if err != nil {
		httpResErr(w, util.EcWrongSvc)
		return
	}
	codeStr := r.PathValue("code")
	codeInt, err := strconv.Atoi(codeStr)
	if err != nil {
		httpResErr(w, util.EcWrongMethod)
		return
	}
	svc := kiwi.TSvc(svcInt)
	code := kiwi.TMethod(codeInt)

	roleMask, _ := util.MGet[int64](head, common.HdMask)
	ok := kiwi.Gate().Authenticate(roleMask, svc, code)
	if !ok {
		httpResErr(w, util.EcNoAuth)
		return
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		_ = r.Body.Close()
		httpResErr(w, util.EcBadPacket)
		return
	}
	_ = r.Body.Close()

	failCh := make(chan uint16, 1)
	okCh := make(chan []byte, 1)
	_svc.AsyncReqBytes(0, svc, code, head, true, payload, func(code uint16) {
		failCh <- code
	}, func(bytes []byte) {
		okCh <- bytes
	})
	res := Response{}
	select {
	case c := <-failCh:
		res.Code = c
	case b := <-okCh:
		resCode, err := kiwi.Codec().ReqToResMethod(svc, code)
		if err != nil {
			httpResErr(w, util.EcServiceErr)
			return
		}
		msg, err := kiwi.Codec().JsonUnmarshal2(svc, resCode, b)
		if err != nil {
			httpResErr(w, util.EcUnmarshallErr)
			return
		}
		res.Data = msg
	}
	resBytes, _ := util.JsonMarshal(res)
	_, _ = w.Write(resBytes)
}

type Response struct {
	Code uint16 `json:"code"`
	Data any    `json:"data"`
}
