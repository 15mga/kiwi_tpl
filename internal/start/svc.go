// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package start

import (
	"game/internal/common"
	"game/internal/gate"
	"game/internal/group"
	"game/internal/user"
	"github.com/15mga/kiwi"
)

var SvcToNew = map[kiwi.TSvc]func(string) kiwi.IService{
	common.Gate:  gate.New,
	common.User:  user.New,
	common.Group: group.New,
}
