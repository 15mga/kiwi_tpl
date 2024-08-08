package common

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/15mga/kiwi/util"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	Issuer          = "15m.games"
	TokenSecret, _  = base64.URLEncoding.DecodeString("95eh.com")
	_ErrWrongIssuer = errors.New("wrong issuer")
)

type Claims struct {
	Issuer    string
	Addr      string
	Uid       string
	Timestamp int64
}

func (c *Claims) Valid() error {
	if c.Issuer != Issuer {
		return _ErrWrongIssuer
	}
	return nil
}

func GenToken(addr, uid string) (string, *util.Err) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Issuer:    Issuer,
		Addr:      addr,
		Uid:       uid,
		Timestamp: time.Now().Unix(),
	})
	tkn, e := token.SignedString(TokenSecret)
	if e != nil {
		return "", util.WrapErr(util.EcServiceErr, e)
	}
	return tkn, nil
}

func ParseToken(tkn string) (claims *Claims, err *util.Err) {
	token, e := jwt.ParseWithClaims(tkn, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	if e != nil {
		err = util.NewErr(util.EcIllegalOp, util.M{
			"token": tkn,
			"error": e.Error(),
		})
		return
	}
	c, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		err = util.NewErr(util.EcIllegalOp, util.M{
			"token": tkn,
		})
		return
	}
	claims = c
	return
}

const (
	Salt = "15m.games"
)

func addSalt(pw string) string {
	return pw + Salt
}

func SaltPw(pw string) string {
	bytes := md5.Sum([]byte(addSalt(pw)))
	return hex.EncodeToString(bytes[:])
}
