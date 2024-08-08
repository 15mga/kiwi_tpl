package main

import (
	"flag"
	"fmt"
	"game/internal/common"
	"game/internal/start"
	"github.com/15mga/kiwi"
	"strings"
)

var (
	version = "0.0.1"
)

func main() {
	s := *flag.String("svc", "all", "service list")
	flag.Parse()
	fmt.Println(s)
	svcSlc := make([]kiwi.TSvc, 0, 1)
	if s != "all" {
		nameSlc := strings.Split(s, ",")
		for _, name := range nameSlc {
			svc, ok := common.NameToSvc[name]
			if !ok {
				continue
			}
			svcSlc = append(svcSlc, svc)
		}
	}
	if len(svcSlc) == 0 {
		for _, svc := range common.NameToSvc {
			svcSlc = append(svcSlc, svc)
		}
	}
	start.Start(version, svcSlc...)
}
