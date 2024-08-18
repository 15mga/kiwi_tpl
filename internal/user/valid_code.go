package user

import (
	"crypto/tls"
	"fmt"
	"game/internal/common"
	"github.com/15mga/kiwi/util"
	"github.com/15mga/kiwi/util/rds"
	"github.com/gomodule/redigo/redis"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
)

func CheckSendCodeAddr(addr string) bool {
	//todo 检查发送code的ip地址是否正常
	return true
}

func CheckMobileCode(mobile, code string) bool {
	testCode := common.Conf().User.Sms.Code
	if testCode != "" && code == testCode {
		return true
	}
	conn := rds.SpawnConn()
	defer conn.Close()

	key := common.GetRdsKey("mobile:sms", mobile)
	rdsCode, err := redis.String(conn.Do(rds.GET, key))
	if err != nil {
		return false
	}
	if code != rdsCode {
		return false
	}
	_, _ = conn.Do(rds.DEL, key)
	return true
}

func SendMobileCode(mobile string) (string, *util.Err) {
	conn := rds.SpawnConn()
	defer conn.Close()

	key := common.GetRdsKey("mobile:sms", mobile)
	_, err := redis.Int(conn.Do(rds.GET, key))
	if err == nil {
		return "", util.NewErr(util.EcTooMuch, util.M{
			"mobile": mobile,
		})
	}

	conf := common.Conf().User.Sms
	code := rand.Intn(9000) + 1000

	_ = conn.Send(rds.SET, key, code)
	_ = conn.Send(rds.EXPIRE, key, 60*5)
	err = conn.Flush()
	if err != nil {
		return "", util.WrapErr(util.EcRedisErr, err)
	}

	if conf.Debug {
		return strconv.Itoa(code), nil
	} else {
		//todo send sms
		return "", nil
	}
}

func CheckEmailCode(email, code string) bool {
	testCode := common.Conf().User.Smtp.Test
	if testCode != "" && code == testCode {
		return true
	}
	conn := rds.SpawnConn()
	defer conn.Close()

	key := common.GetRdsKey("email:code", email)
	rdsCode, err := redis.String(conn.Do(rds.GET, key))
	if err != nil {
		return false
	}
	if code != rdsCode {
		return false
	}
	_, _ = conn.Do(rds.DEL, key)
	return true
}

func SendEmailCode(toEmail string) (string, *util.Err) {
	conn := rds.SpawnConn()
	defer conn.Close()

	key := common.GetRdsKey("email:code", toEmail)
	_, err := redis.Int(conn.Do(rds.GET, key))
	if err == nil {
		return "", util.NewErr(util.EcTooMuch, util.M{
			"email": toEmail,
		})
	}
	conf := common.Conf().User.Smtp
	code := rand.Intn(9000) + 1000

	_ = conn.Send(rds.SET, key, code)
	_ = conn.Send(rds.EXPIRE, key, 60*5)
	err = conn.Flush()
	if err != nil {
		return "", util.WrapErr(util.EcRedisErr, err)
	}

	if conf.Debug {
		return strconv.Itoa(code), nil
	} else {
		e := email.NewEmail()
		e.From = conf.Email
		e.To = []string{toEmail}
		e.Subject = conf.Subject
		e.Text = []byte(fmt.Sprintf(conf.Body, code))
		auth := smtp.PlainAuth("", conf.Email, conf.Password, conf.Host)
		addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
		err = e.SendWithTLS(addr, auth, &tls.Config{ServerName: conf.Host})
		if err != nil {
			return "", util.WrapErr(util.EcServiceErr, err)
		}
		return "", nil
	}
}
