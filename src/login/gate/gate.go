package gate

import (
	"github.com/snowyyj001/loumiao"
	lgate "github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/network"
)

var GateInfos map[string]int

func StartGate() {
	watchdog := &lgate.GateServer{ServerType: network.CLIENT_CONNECT}
	loumiao.Prepare(watchdog, "GateServer", false)

	lgate.RegisterNet(watchdog, gorpc.M{Name: "G_L_GateInfo", Data: "GateServer"})
	watchdog.RegisterGate("G_L_GateInfo", handlerGateInfo)
}

func GetPerfactGate() string {
	var minNum = 0x7fffffff
	var as string = ""
	for addr, num := range GateInfos {
		if minNum > num {
			minNum = num
			as = addr
		}
	}
	return as
}

func init() {
	GateInfos = make(map[string]int)
}
