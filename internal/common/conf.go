package common

import (
	"fmt"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/loader"
)

var (
	_Conf Config
)

func Conf() *Config {
	return &_Conf
}

type Config struct {
	Mode  string
	Log   LogConfig
	Redis RedisConf
	Mongo MgoConf
	Gate  GateConf
}

type LogConfig struct {
	Std LogStd
	Mgo LogMgo
}

type LogStd struct {
	Enable bool
	Color  bool
	Log    []string
	Trace  []string
}

type LogMgo struct {
	Enable bool
	Log    []string
	Trace  []string
	Uri    string
	Db     string
	Ttl    int32
}

type RedisConf struct {
	Addr     string
	User     string
	Password string
	Db       int
}

type MgoConf struct {
	Uri string
	Db  string
}

type GateConf struct {
	Ip           string
	Tcp          int
	Web          int
	Udp          int
	Http         int
	HttpSuper    int
	ConnCap      int32
	DeadlineSecs int
	PacketLimit  uint16
	ErrLimit     uint16
	KeyFile      string
	PemFile      string
}

func LoadConf(conf any, confFolder string, svc ...kiwi.TSvc) {
	loader.SetConfRoot(confFolder)
	convPath := loader.ConvertConfLocalPath
	slc := make([]string, 0, len(svc)+1)
	slc = append(slc, "common.yml") //通用
	for _, s := range svc {
		p, ok := SvcToName[s]
		if !ok {
			continue
		}
		slc = append(slc, fmt.Sprintf("%s.yml", p))
	}
	loader.LoadConf(conf, convPath(slc...)...)
}
