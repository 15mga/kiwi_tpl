// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package common

import (
	"github.com/15mga/kiwi"
)

func RoleToStr(role int64) string {
	switch role {
	case RGuest:
		return "guest"
	case RPlayer:
		return "player"
	case ROps:
		return "ops"
	default:
		return ""
	}
}

func StrToRole(role string) int64 {
	switch role {
	case "guest":
		return RGuest
	case "player":
		return RPlayer
	case "ops":
		return ROps
	default:
		return 0
	}
}

var MsgRole = map[kiwi.TSvcCode][]int64{
	1000: {RGuest, RPlayer, ROps},
	1004: {RPlayer, ROps},
	1006: {RGuest},
	2000: {RGuest},
	2002: {RGuest},
	2004: {RGuest},
	2006: {RGuest},
	2008: {RPlayer},
}
