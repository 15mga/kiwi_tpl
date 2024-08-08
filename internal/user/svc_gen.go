// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package user

import (
	"game/internal/common"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/core"
)

var (
	_svc *Svc
)

func New(ver string) kiwi.IService {
	_svc = &Svc{
		svc: svc{
			Service: core.NewService(common.User, ver),
		}}
	return _svc
}

type Svc struct {
	svc
}

type svc struct {
	core.Service
}

func (s *svc) Start() {
	initColl()
	registerReq()
}

func (s *svc) AfterStart() {
}
