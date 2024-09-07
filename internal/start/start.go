package start

import (
	"fmt"
	"game/internal/codec"
	"game/internal/common"
	"game/internal/gate"
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/core"
	"github.com/15mga/kiwi/log"
	"github.com/15mga/kiwi/network"
	"github.com/15mga/kiwi/util"
	"github.com/15mga/kiwi/util/mgo"
	"github.com/dgraph-io/ristretto"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
)

func Start(ver string, svc ...kiwi.TSvc) {
	util.SetJsonConf(jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	})
	wd, _ := os.Getwd()
	//加载配置文件
	common.LoadConf(common.Conf(), fmt.Sprintf("%s/config", wd), svc...)
	conf := common.Conf()
	nodeInfo := kiwi.GetNodeMeta()
	nodeInfo.Mode = conf.Mode

	//文件日志
	var loggers []kiwi.ILogger
	logStd := conf.Log.Std
	if logStd.Enable {
		loggers = append(loggers, log.NewStd(
			log.StdColor(logStd.Color),
			log.StdLogStrLvl(logStd.Log...),
			log.StdTraceStrLvl(logStd.Trace...),
			log.StdColor(true),
		))
	} else {
		loggers = append(loggers, log.NewStd(
			log.StdColor(logStd.Color),
			log.StdLogStrLvl(logStd.Log...),
			log.StdTraceStrLvl(logStd.Trace...),
			log.StdFile(fmt.Sprintf("%s/log/%s.log", wd, util.ExeName())),
		))
	}

	//mongo日志
	logMgo := conf.Log.Mgo
	if logMgo.Enable {
		loggers = append(loggers, log.NewMgo(
			log.MgoLogLvl(logMgo.Log...),
			log.MgoTraceLvl(logMgo.Trace...),
			log.MgoClientOptions(options.Client().ApplyURI(logMgo.Uri)),
			log.MgoDb(logMgo.Db),
			log.MgoTtl(logMgo.Ttl),
		))
	}
	// 排除打印的链路日志
	core.ExcludeTrace(common.Gate) //codec.GateUpdateReq, codec.GateUpdateRes,

	// 排除打印的消息，用于当
	core.ExcludeMsg(common.Gate) //codec.GateUpdateReq, codec.GateUpdateRes,

	service := make([]kiwi.IService, 0, len(svc))
	hasGate := false
	for _, s := range svc {
		if s == common.Gate {
			hasGate = true
		}
		fn, ok := SvcToNew[s]
		if ok {
			service = append(service, fn(ver))
		} else {
			kiwi.Warn2(util.EcNotImplement, util.M{
				"svc": s,
			})
		}
	}

	common.InitCron()
	opts := []core.Option{
		core.SetSnowflakeNodeId(common.GetSnowflakeNodeId),
		core.SetLoggers(loggers...),
		core.SetMongoDB(conf.Mongo.Uri, conf.Mongo.Db, nil),
		core.SetRedis(conf.Redis.Addr, conf.Redis.User, conf.Redis.Password, conf.Redis.Db),
		//设置服务
		core.SetServices(service...),
		core.SetNode(),
		core.SetBefore(func() {
			codec.BindReqToRes()
			codec.BindPool()
		}),
		core.SetAfter(func() {
			common.StartDiscovery()
			mgo.InitCache(&ristretto.Config{
				NumCounters: 1e7,     // 10 million counters
				MaxCost:     1 << 30, //1GB
				BufferItems: 64,
				Metrics:     true,
			})
		}),
	}

	keyFile, pemFile := "", ""
	if conf.Gate.KeyFile != "" {
		keyFile = fmt.Sprintf("%s/%s", wd, conf.Gate.KeyFile)
	}
	if conf.Gate.PemFile != "" {
		pemFile = fmt.Sprintf("%s/%s", wd, conf.Gate.PemFile)
	}
	if hasGate {
		gateOptions := []core.GateOption{
			core.GateIp(conf.Gate.Ip),
			core.GateHttpPort(conf.Gate.Http),
			core.GateRoles(gate.MsgRole),
			core.GateInitHead(gate.InitHead),
			core.GateHttpStatic(gate.StaticHandler),
			core.GateSocketReceiver(gate.SocketReceiver),
			core.GateWebsocketPort(conf.Gate.Web),
			core.GateConnCap(conf.Gate.ConnCap),
			core.GateDeadlineSecs(conf.Gate.DeadlineSecs),
			core.GateCheckIp(gate.CheckIp),
			core.GateDisconnected(gate.Disconnected),
			core.GateHeadLen(4),
			core.GateWebsocketOptions(
				network.WebUpgrader(func(upgrader *websocket.Upgrader) {
					upgrader.CheckOrigin = func(r *http.Request) bool {
						return true
					}
				}),
				network.WebTls(pemFile, keyFile),
			),
			core.GateHttpReceiver(gate.HttpReceiver),
			core.GateHttpHeadCache(&gate.HttpHeadCache{}),
			core.GateHttpDisconnected(gate.HttpDisconnected),
			core.GateHttpExpireMins(1),
		}
		opts = append(opts, core.SetGate(gateOptions...))
	}

	core.Default(opts...)

	kiwi.WaitExit()
}
