// Code generated by protoc-gen-go-kiwi. DO NOT EDIT.

package common

import (
	"github.com/15mga/kiwi"
)

const (
	Gate kiwi.TSvc = 1
	User kiwi.TSvc = 2
)

const (
	SGate = "gate"
	SUser = "user"
)

var SvcToName = map[kiwi.TSvc]string{
	Gate: SGate,
	User: SUser,
}

var NameToSvc = map[string]kiwi.TSvc{
	SGate: Gate,
	SUser: User,
}
