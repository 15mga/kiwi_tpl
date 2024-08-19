package common

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var (
	Issuer          = "15m.games"
	TokenSecret, _  = base64.URLEncoding.DecodeString("95eh.com")
	_ErrWrongIssuer = errors.New("wrong issuer")
)

type Claims struct {
	Issuer    string
	SignInCh  string
	Addr      string
	UserId    string
	Timestamp int64
}

func (c *Claims) Valid() error {
	if c.Issuer != Issuer {
		return _ErrWrongIssuer
	}
	return nil
}

func GenToken(signInCh, addr, uid string, ts int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Issuer:    Issuer,
		SignInCh:  signInCh,
		Addr:      addr,
		UserId:    uid,
		Timestamp: ts,
	})
	tkn, e := token.SignedString(TokenSecret)
	if e != nil {
		return "", e
	}
	return tkn, nil
}

func ParseToken(tkn string) (*Claims, error) {
	token, e := jwt.ParseWithClaims(tkn, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	if e != nil {
		return nil, e
	}
	c, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return c, nil
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
