package common

import (
	"github.com/15mga/kiwi"
	"github.com/15mga/kiwi/core"
	"github.com/15mga/kiwi/util"
	"github.com/15mga/kiwi/util/rds"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func StartDiscovery() {
	renewal()
	registerSvc()
	connectSvc()
}

func GetSnowflakeNodeId() int64 {
	conn := rds.SpawnConn()
	defer conn.Close()

	sid := int64(0)
	_ = rds.Lock(GetRdsKey("node_lock"), func() {
		var sidKeys []string
		err := rds.Scan(conn, GetRdsKey("node")+"*", 16, func(keys []string) {
			sidKeys = append(sidKeys, keys...)
		})
		if err != nil {
			kiwi.Error3(util.EcRedisErr, err)
			return
		}
		idMap := make(map[int64]struct{})
		for _, key := range sidKeys {
			strSlc := SplitRdsKey(key)
			id, err := strconv.ParseInt(strSlc[len(strSlc)-1], 10, 64)
			if err != nil {
				continue
			}
			idMap[id] = struct{}{}
		}
		for i := int64(0); i < 1024; i++ {
			if _, ok := idMap[i]; ok {
				continue
			}
			sid = i
			// 锁内抢占id
			sidStr := strconv.FormatInt(sid, 10)
			conn.Send(rds.HSET, GetRdsKey("node", sidStr), "startTime", kiwi.GetNodeMeta().StartTime)
			conn.Send(rds.EXPIRE, 5)
			conn.Flush()
			return
		}
	})
	return sid
}

func registerSvc() {
	conn := rds.SpawnConn()
	defer conn.Close()

	services := core.AllService()
	for _, service := range services {
		svcMeta := service.Meta()
		key := GetRdsKey("svc", SvcToName[service.Svc()])
		addRenewalKey(key)
		conn.Send(rds.HSET, key,
			"id", svcMeta.Id,
			"sid", svcMeta.Sid,
			"nid", svcMeta.Nid,
			"ip", svcMeta.Ip,
			"port", svcMeta.Port,
			"startTime", svcMeta.StartTime,
			"svc", service.Svc(),
			"ver", service.Ver(),
		)
	}
	nodeMeta := kiwi.GetNodeMeta()
	sidStr := strconv.FormatInt(nodeMeta.Sid, 10)
	key := GetRdsKey("node", sidStr)
	addRenewalKey(key)
	conn.Send(rds.HSET, key,
		"nid", nodeMeta.Nid,
		"ip", nodeMeta.Ip,
		"port", nodeMeta.Port,
		"startTime", nodeMeta.StartTime,
	)
	_ = conn.Flush()
}

var (
	_renewalKeyCh = make(chan string, 1)
)

func addRenewalKey(key string) {
	_renewalKeyCh <- key
}

// 续约 key
func renewal() {
	go func() {
		keys := make([]string, 0, 16)
		for {
			select {
			case key := <-_renewalKeyCh:
				keys = append(keys, key)
			case <-time.After(time.Second * 5):
				conn := rds.SpawnConn()
				for _, key := range keys {
					conn.Send(rds.EXPIRE, key, 10)
				}
				conn.Flush()
				conn.Close()
				connectSvc()
			case <-util.Ctx().Done():
				conn := rds.SpawnConn()
				for _, key := range keys {
					conn.Send(rds.DEL, key)
				}
				conn.Flush()
				conn.Close()
				return
			}
		}
	}()
}

var (
	_SvcIdMap = make(map[int64]struct{})
)

func connectSvc() {
	conn := rds.SpawnConn()
	defer conn.Close()

	var svcKeys []string
	err := rds.Scan(conn, GetRdsKey("svc")+"*", 16, func(keys []string) {
		svcKeys = append(svcKeys, keys...)
	})
	if err != nil {
		kiwi.Error3(util.EcRedisErr, err)
		return
	}
	newIdMap := make(map[int64]struct{}, len(svcKeys))
	selfNodeId := kiwi.GetNodeMeta().Nid
	for _, key := range svcKeys {
		m, err := redis.StringMap(conn.Do(rds.HGETALL, key))
		if err != nil {
			continue
		}
		nid, _ := strconv.ParseInt(m["nid"], 10, 64)
		if nid == selfNodeId {
			continue
		}
		id, _ := strconv.ParseInt(m["id"], 10, 64)
		sid, _ := strconv.ParseInt(m["sid"], 10, 64)
		port, _ := strconv.Atoi(m["port"])
		startTime, _ := strconv.ParseInt(m["startTime"], 10, 64)
		svc, _ := strconv.ParseInt(m["svc"], 10, 64)
		meta := kiwi.SvcMeta{
			Id:        id,
			Sid:       sid,
			Nid:       nid,
			Ip:        m["ip"],
			Port:      port,
			StartTime: startTime,
			Svc:       kiwi.TSvc(svc),
			Ver:       m["ver"],
		}
		newIdMap[id] = struct{}{}
		kiwi.Node().Connect(meta)
	}
	for id := range _SvcIdMap {
		if _, ok := newIdMap[id]; !ok {
			kiwi.Node().Disconnect(id)
		}
	}
	_SvcIdMap = newIdMap
}
