// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package group

import (
	"game/internal/codec"
	"game/internal/common"
	"game/proto/pb"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/core"
	"github.com/15mga/kiwi/util"
)

func registerReq() {
	kiwi.Router().BindReq(common.Group, codec.GroupNewReq, func(req kiwi.IRcvRequest) {
		if _svc.IsShutdown() {
			return
		}
		req.SetReceiver(_svc)
		key, _ := util.MGet[string](req.Head(), "group_id")
		core.ActivePrcReq[*pb.GroupNewReq](req, key, _svc.OnGroupNew)
	})
}
